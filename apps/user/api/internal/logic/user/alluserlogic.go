package user

import (
	"context"

	"ZeZeIM/apps/user/api/internal/svc"
	"ZeZeIM/apps/user/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
