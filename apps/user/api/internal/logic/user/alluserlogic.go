package user

import (
	"ZChat/apps/user/api/internal/svc"
	"ZChat/apps/user/api/internal/types"
	"ZChat/apps/user/rpc/pb/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AlluserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取所有用户
func NewAlluserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlluserLogic {
	return &AlluserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlluserLogic) Alluser(req *types.AllUserReq) (resp *types.AllUserResp, err error) {

	users, err := l.svcCtx.User.AllUser(l.ctx, &user.AllUserReq{})
	if err != nil {
		return nil, err
	}
	var r []types.UserEntity
	for _, user := range users.User {
		r = append(r, types.UserEntity{
			Id:   user.Id,
			Name: user.Name,
		})
	}
	return &types.AllUserResp{Users: r}, nil
}
