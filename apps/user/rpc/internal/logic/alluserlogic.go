package logic

import (
	"context"

	"ZeZeIM/apps/user/rpc/internal/svc"
	"ZeZeIM/apps/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllUserLogic {
	return &AllUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllUserLogic) AllUser(in *user.AllUserReq) (*user.AllUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.AllUserResp{}, nil
}
