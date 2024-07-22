package logic

import (
	"ZChat/apps/user/model"
	"context"
	"github.com/jinzhu/copier"

	"ZChat/apps/user/rpc/internal/svc"
	"ZChat/apps/user/rpc/pb/user"

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

	var (
		userEntitys []*model.Users
		err         error
	)
	userEntitys, err = l.svcCtx.UsersModel.AllUser(l.ctx)

	if err != nil {
		return nil, err
	}

	var resp []*user.UserEntity
	copier.Copy(&resp, &userEntitys)

	return &user.AllUserResp{User: resp}, nil
}
