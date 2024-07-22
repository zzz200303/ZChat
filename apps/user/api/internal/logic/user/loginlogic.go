package user

import (
	"ZChat/apps/user/api/internal/svc"
	"ZChat/apps/user/api/internal/types"
	"ZChat/apps/user/rpc/pb/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登入
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.User.Login(l.ctx, &user.LoginReq{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var res types.LoginResp
	res.Expire = loginResp.Expire
	res.Token = loginResp.Token
	return &res, nil
}
