package logic

import (
	"ZeZeIM/apps/user/models"
	"context"
	"github.com/jinzhu/copier"

	"ZeZeIM/apps/user/rpc/internal/svc"
	"ZeZeIM/apps/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserByIDLogic {
	return &FindUserByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserByIDLogic) FindUserByID(in *user.FindUserByIDReq) (*user.FindUserByIDResp, error) {
	var (
		userEntity *models.Users
		err        error
	)

	userEntity, err = l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	var resp []*user.UserEntity
	err = copier.Copy(&resp, &userEntity)
	if err != nil {
		return nil, err
	}

	return &user.FindUserByIDResp{User: resp}, nil
}
