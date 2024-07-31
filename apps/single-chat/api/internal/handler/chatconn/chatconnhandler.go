package chatconn

import (
	"ZChat/apps/single-chat/api/internal/svc"
	"ZChat/apps/single-chat/api/internal/types"
	"ZChat/apps/user/rpc/pb/user"
	"ZChat/pkg/constants"
	utils "ZChat/pkg/pool"
	"ZChat/pkg/response"
	"context"                                // 导入上下文包，用于处理超时和取消操作
	"encoding/json"                          // 导入JSON编码解码包
	"fmt"                                    // 导入格式化IO包
	"github.com/gorilla/websocket"           // 导入WebSocket包
	"github.com/zeromicro/go-zero/core/logx" // 导入日志包
	"net/http"                               // 导入HTTP包
	"time"                                   // 导入时间包
)

// UsersMap 存储在线用户的映射，键是用户ID，值是用户节点
var usersMap map[int64]*Node

// Node 结构体定义了用户节点
type Node struct {
	Uid                int64                  // 用户ID
	WsConn             *websocket.Conn        // 用户的WebSocket连接
	CacheOnlineMessage chan types.MessageInfo // 缓存在线消息的通道
	SvcCtx             *svc.ServiceContext    // 服务上下文
}

// ChatConnHandler 处理WebSocket连接请求
func ChatConnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uidJson := r.Context().Value("uid").(json.Number) // 从jwt里面提取uid
		uid, err := uidJson.Int64()
		if err != nil {
			fmt.Println("json.Number换出了问题")
			return
		}
		node, ok := usersMap[uid] // 检查用户是否在线
		if !ok {
			fmt.Println("用户不在线")
			return // 如果用户不在线，直接返回
		}
		conn, err := upgrade(w, r, svcCtx) // 升级HTTP连接到WebSocket连接
		if err != nil {
			response.Response(w, nil, response.NewRespError(response.WEBCOKET_ERROR, response.WEBCOKET_ERROR_MESSAGE)) // 处理升级错误
			return
		}
		// 绑定用户和WebSocket连接
		usersMap[uid].WsConn = conn
		// 每个用户开启goroutine监听消息收发
		utils.Pool.Submit(node.SendMQMessage) // 启动发送消息的goroutine
		utils.Pool.Submit(node.RecvMessage)   // 启动接收消息的goroutine
		//将 SendMQMessage 和 RecvMessage 方法提交到线程池中，以异步运行这些方法
		logx.Infof("用户%d连接成功......", uid) // 记录连接成功日志
	}
}

// upgrade 升级HTTP连接到WebSocket连接
func upgrade(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext) (*websocket.Conn, error) {
	ws := websocket.Upgrader{
		HandshakeTimeout: time.Duration(svcCtx.Config.Client.Upgrade.HandshakeTimeout) * time.Second, // 设置握手超时时间
		ReadBufferSize:   int(svcCtx.Config.Client.Upgrade.ReadBufferSize),                           // 设置读取缓冲区大小
		WriteBufferSize:  int(svcCtx.Config.Client.Upgrade.WriteBufferSize),                          // 设置写入缓冲区大小
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有来源的连接
		},
	}
	conn, err := ws.Upgrade(w, r, nil) // 执行协议升级
	if err != nil {
		return nil, err // 处理升级错误
	}
	return conn, nil // 返回升级后的连接
}

// InitAllUser 初始化所有用户信息
func InitAllUser(svcCtx *svc.ServiceContext) error {
	ctx := context.Background() // 创建上下文
	userList, err := svcCtx.UserRpcService.AllUser(ctx, &user.AllUserReq{})
	var userIdList []int64
	for _, t := range userList.User {
		userIdList = append(userIdList, t.Id)
	} // 获取所有用户ID列表

	if err != nil {
		logx.Error("初始化用户异常") // 记录错误日志
		return err
	}
	usersMap = make(map[int64]*Node, len(userIdList)) // 初始化用户映射

	for _, u := range userIdList { // 遍历用户ID列表
		usersMap[u] = &Node{ // 为每个用户创建节点
			Uid:                u,
			WsConn:             (*websocket.Conn)(nil),
			CacheOnlineMessage: make(chan types.MessageInfo, svcCtx.Config.Client.MessageBuf), // 初始化在线消息通道
			SvcCtx:             svcCtx,
		}
	}

	logx.Info("初始化用户成功, 用户Uid:") // 记录初始化成功日志
	fmt.Println(userIdList)      // 输出用户ID列表

	return nil // 返回nil表示成功
}

func InitUser(svcCtx *svc.ServiceContext, uid int64) error {
	usersMap[uid] = &Node{ // 为新用户创建节点
		Uid:                uid,
		WsConn:             (*websocket.Conn)(nil),
		CacheOnlineMessage: make(chan types.MessageInfo, svcCtx.Config.Client.MessageBuf), // 初始化在线消息通道
		SvcCtx:             svcCtx,
	}
	logx.Info("初始化用户成功") // 记录初始化成功日志
	return nil           // 返回nil表示成功
}

// SendMQMessage 从前端接收到消息,发送到Kafka
func (n *Node) SendMQMessage() {
	defer func() {
		n.WsConn.Close()                                // 关闭WebSocket连接
		usersMap[n.Uid].WsConn = (*websocket.Conn)(nil) // 清理用户节点的连接
		logx.Errorf("用户%d的websocket连接断开", n.Uid)        // 记录连接断开日志
	}()
	for {
		_, buf, err := n.WsConn.ReadMessage() // 读取WebSocket消息
		logx.Infof(string(buf))               // 记录收到的消息
		if err != nil {
			return // 处理读取错误
		}
		err = n.SvcCtx.KqPusher.Push(string(buf)) // 将消息推送到Kafka
		logx.Info("kafka存储消息....")                // 记录Kafka存储日志
		if err != nil {
			logx.Errorf("kafka接收消息异常", n.Uid) // 处理Kafka推送错误
			return
		}
	}
}

// RecvMessage 接受调度过来的消息，推送到前端
func (n *Node) RecvMessage() {
	for {
		select {
		case message := <-n.CacheOnlineMessage: // 从在线消息通道读取消息
			buf, err := json.Marshal(message)             // 将消息编码为JSON
			logx.Infof("用户%d收到消息：%s", n.Uid, string(buf)) // 记录收到的消息
			if err != nil {
				logx.Error(err) // 处理编码错误
				return
			}
			err = n.WsConn.WriteMessage(websocket.TextMessage, buf) // 发送消息到前端
			if err != nil {
				logx.Error(err) // 处理发送错误
				return
			}
		}
	}
}

// GetOfflineMessage 处理离线消息 从Redis中获取离线消息, 发送到用户通道中
func (n *Node) GetOfflineMessage() {
	offlineMessage, err := n.SvcCtx.Redis.Lrange(fmt.Sprintf("%s::%d", constants.SingleChatMsg, n.Uid), 0, -1)
	// 从Redis获取离线消息
	//生成一个 Redis 列表的键，格式为 utils.SingleChat::用户ID，例如 SingleChat::12345。
	//调用 Redis 客户端的 Lrange 方法，获取该键对应的列表中从第一个到最后一个元素（即获取所有的离线消息）。
	//将获取到的离线消息存储在 offlineMessage 切片中，如果发生错误，则存储在 err 变量中。
	if err != nil {
		logx.Error(err) // 处理获取错误
		return
	}
	// 没有缓存消息
	if len(offlineMessage) == 0 {
		return // 如果没有离线消息，直接返回
	}
	for _, val := range offlineMessage { // 遍历离线消息
		var msg types.MessageInfo
		err := json.Unmarshal([]byte(val), &msg) // 解码JSON消息
		if err != nil {
			logx.Error(err) // 处理解码错误
			return
		}
		n.CacheOnlineMessage <- msg // 将消息发送到在线消息通道
	}
	// 删除Redis中的离线消息键
	_, err = n.SvcCtx.Redis.Del(fmt.Sprintf("%s::%d", constants.SingleChatMsg, n.Uid))
	if err != nil {
		logx.Error(err) // 处理删除错误
		return
	}
}
