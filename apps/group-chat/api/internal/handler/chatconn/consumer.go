package chatconn

import (
	"ZChat/apps/group-chat/api/internal/svc"
	"ZChat/apps/group-chat/api/internal/types"
	"ZChat/pkg/constants"
	"context"
	"encoding/json" // 导入JSON编码/解码包
	"fmt"           // 导入格式化包
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-queue/kq"       // 导入Kafka队列包
	"github.com/zeromicro/go-zero/core/logx" // 导入日志包
	"strconv"
)

// StartMq 启动消息队列消费
func StartMq(svcCtx *svc.ServiceContext) {
	// 创建Kafka队列
	q := kq.MustNewQueue(svcCtx.Config.GroupMsgKqConf, kq.WithHandle( //kafka接收到消息后的处理函数
		func(k, v string) error {
			logx.Info("kafka消费消息.......") // 记录日志，表示正在消费Kafka消息
			err := dispatch(v, svcCtx)    // 调用dispatch处理消息
			if err != nil {
				return err // 如果处理消息时发生错误，返回错误
			}
			return nil // 处理成功，返回nil
		}))
	defer q.Stop() // 在函数退出时停止队列
	q.Start()      // 启动队列
}

// dispatch 处理Kafka的消息，分类消费
func dispatch(value string, svcCtx *svc.ServiceContext) error {
	var messageJson types.MessageJson                  // 创建一个Message对象
	err := json.Unmarshal([]byte(value), &messageJson) // 将JSON字符串解析为Message对象
	if err != nil {
		logx.Errorf("解析kafka消息异常") // 如果解析发生错误，记录错误日志
		return err                 // 返回错误
	}
	message := types.MessageInfo{
		Content:  messageJson.Content,
		SendTime: messageJson.SendTime,
	}
	message.From, err = strconv.ParseInt(messageJson.From, 10, 64)
	if err != nil {
		return err // 返回错误
	}
	message.To, err = strconv.ParseInt(messageJson.To, 10, 64)
	if err != nil {
		return err // 返回错误
	}
	sendGroupMessage(message, svcCtx) // 处理用户消息
	return nil                        // 处理成功，返回nil
}

// sendUserMessage 按照群聊号发送给所有群晕
func sendGroupMessage(message types.MessageInfo, svcCtx *svc.ServiceContext) {
	userList, err := svcCtx.GmemberModel.FindAllUserByGroupId(context.Background(), message.To)
	if err != nil {
		logx.Errorf("查询群聊用户异常")
		return
	}

	for _, u := range userList {

		node, ok := usersMap[u.Uid] // 从用户映射中获取目标用户节点
		// 序列化 MessageInfo 结构体到 JSON 字符串
		jsonData, err := json.Marshal(message)
		if err != nil {
			// 处理序列化错误
			fmt.Printf("序列化错误: %v", err)
			return
		}
		if !ok {
			return // 如果用户不在线，直接返回
		}
		if node.WsConn == (*websocket.Conn)(nil) {
			// 用户不在线，将消息存储到Redis
			key := fmt.Sprintf("%s:from:%d:group:%d:to:%d", constants.GroupChatMsg, message.From, message.To, u.Uid) // 生成Redis键
			_, err := node.SvcCtx.Redis.Lpush(key, string(jsonData))                                                 // 将消息推送到Redis列表中
			if err != nil {
				fmt.Println("Lpush错误")
				logx.Error(err) // 如果存储发生错误，记录错误日志
			}
			return // 返回
		}
		// 如果用户在线，将消息发送到在线消息通道，不用redis
		node.CacheOnlineMessage <- message
	}
}

// StartNewUserMq 启动新用户消息队列消费
func StartNewUserMq(svcCtx *svc.ServiceContext) {
	// 创建Kafka队列
	q := kq.MustNewQueue(svcCtx.Config.NewUserKqConf, kq.WithHandle( //kafka接收到消息后的处理函数
		func(k, v string) error {
			logx.Info("kafka消费消息.......") // 记录日志，表示正在消费Kafka消息
			uid, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				fmt.Println(err)
				return err // 返回错误
			}
			err = InitUser(svcCtx, uid) // 调用dispatch处理消息
			if err != nil {
				return err // 如果处理消息时发生错误，返回错误
			}
			return nil // 处理成功，返回nil
		}))
	defer q.Stop() // 在函数退出时停止队列
	q.Start()      // 启动队列
}
