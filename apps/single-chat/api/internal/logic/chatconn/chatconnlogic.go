package chatconn

import (
	"ZChat/apps/single-chat/api/internal/types"
	"context"
	"github.com/gorilla/websocket"

	"ZChat/apps/single-chat/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatConnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatConnLogic {
	return &ChatConnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UsersMap 存储在线用户的映射，键是用户ID，值是用户节点
var usersMap map[int64]*Node

// Node 结构体定义了用户节点
type Node struct {
	Uid                int64                  // 用户ID
	WsConn             *websocket.Conn        // 用户的WebSocket连接
	CacheOnlineMessage chan types.MessageInfo // 缓存在线消息的通道
	SvcCtx             *svc.ServiceContext    // 服务上下文
}

func (l *ChatConnLogic) ChatConn() error {
	// todo: add your logic here and delete this line

	return nil
}
