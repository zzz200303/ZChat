package logic

import (
	"context"
	"fmt"

	"ZeZeIM/apps/user/rpc/internal/svc"
	"ZeZeIM/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *user.Request) (*user.Response, error) {
	// todo: add your logic here and delete this line
	fmt.Println("User Ping通了")
	return &user.Response{
		Pong: "pong",
	}, nil
}
