package logic

import (
	"ZeZeIM/apps/user/models"
	"context"
	"github.com/jinzhu/copier"

	"ZeZeIM/apps/user/rpc/internal/svc"
	"ZeZeIM/apps/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserByNameLogic {
	return &FindUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserByNameLogic) FindUserByName(in *user.FindUserByNameReq) (*user.FindUserByNameResp, error) {
	var (
		userEntitys []*models.Users
		err         error
	)

	userEntitys, err = l.svcCtx.UsersModel.ListByName(l.ctx, in.Name)

	if err != nil {
		return nil, err
	}
	var resp []*user.UserEntity
	err = copier.Copy(&resp, &userEntitys)
	if err != nil {
		return nil, err
	}
	return &user.FindUserByNameResp{User: resp}, nil
}
