package user

import (
	"context"

	"ZeZeIM/apps/user/api/internal/svc"
	"ZeZeIM/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FinduserbynameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用name查找用户
func NewFinduserbynameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FinduserbynameLogic {
	return &FinduserbynameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FinduserbynameLogic) Finduserbyname(req *types.FindUserByNameReq) (resp *types.FindUserByNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
