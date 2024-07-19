// 导入所需的包。
package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// 定义写入消息到websocket的超时时间为 10 秒。
	writeWait = 10 * time.Second

	// 定义读取下一个 pong 消息的超时时间为 60 秒。
	pongWait = 60 * time.Second

	// 发送 ping 消息的频率，需小于 pongWait。
	pingPeriod = (pongWait * 9) / 10

	// 定义从客户端接收的最大消息大小为 512 字节。
	maxMessageSize = 512
)

// 定义换行符和空格字节数组。
var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// 配置 WebSocket 升级器，定义读写缓冲区大小。
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client 结构体充当 websocket 连接与 hub 之间的中介。
type Client struct {
	// 指向 hub 的引用。
	hub *Hub

	// 客户端的 websocket 连接。
	conn *websocket.Conn

	// 用于存储出站消息的缓冲通道。
	send chan []byte
}

// readPump 方法从 websocket 连接读取消息并发送到 hub。
func (c *Client) readPump() {
	defer func() {
		// 从 hub 注销客户端并关闭连接。
		c.hub.unregister <- c
		c.conn.Close()
	}()
	// 设置读取限制和 pong 消息的处理函数。
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		// 读取消息，如果发生错误则退出循环。
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			// 检查是否是意外关闭错误。
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// 清理消息中的换行符并广播到 hub。
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump 方法从 hub 接收消息并写入 websocket 连接。
func (c *Client) writePump() {
	// 创建一个定时器以定期发送 ping 消息。
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		// 停止定时器并关闭连接。
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		// 从发送通道读取消息并写入连接。
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// 如果通道关闭，发送关闭消息。
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// 获取写入器并写入消息。
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// 将排队的消息添加到当前 websocket 消息中。
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			// 关闭写入器。
			if err := w.Close(); err != nil {
				return
			}
		// 定时器触发时发送 ping 消息。
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs 处理客户端的 websocket 请求。
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 websocket 连接。
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 创建一个新的客户端实例并注册到 hub。
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// 启动读写 goroutine 来处理消息。
	go client.writePump()
	go client.readPump()
}
