package logic

import (
	"ZeZeIM/apps/user/models"
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"ZeZeIM/apps/user/rpc/internal/svc"
	"ZeZeIM/apps/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	userEntiy, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errors.New("没查到")
		}
		return nil, err
	}
	var resp user.UserEntity
	err = copier.Copy(&resp, userEntiy)
	if err != nil {
		return nil, err
	}

	return &user.GetUserInfoResp{
		User: &resp,
	}, nil
}
