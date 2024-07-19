package main

// Hub 维护活跃客户端集合，并向客户端广播消息。
type Hub struct {
	// 已注册的客户端。
	clients map[*Client]bool

	// 从客户端接收的入站消息。
	broadcast chan []byte

	// 客户端的注册请求。
	register chan *Client

	// 客户端的注销请求。
	unregister chan *Client
}

// newHub 创建一个新的 Hub 实例。
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),      // 初始化广播通道。
		register:   make(chan *Client),     // 初始化注册通道。
		unregister: make(chan *Client),     // 初始化注销通道。
		clients:    make(map[*Client]bool), // 初始化客户端集合。
	}
}

// run 方法监听各种通道以管理客户端。
func (h *Hub) run() {
	for {
		select {
		// 处理注册请求。
		case client := <-h.register:
			h.clients[client] = true
		// 处理注销请求。
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				// 如果客户端存在，删除并关闭其发送通道。
				delete(h.clients, client)
				close(client.send)
			}
		// 处理入站消息并广播给所有客户端。
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
					// 将消息发送给客户端。
				default:
					// 如果发送失败，关闭并移除客户端。
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
