package logic

import (
	"context"

	"ZeZeIM/user/rpc/internal/svc"
	"ZeZeIM/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.UserRequest) (*pb.UserResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UserResponse{}, nil
}
