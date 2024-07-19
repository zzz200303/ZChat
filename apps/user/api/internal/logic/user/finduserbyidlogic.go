package user

import (
	"context"

	"ZeZeIM/apps/user/api/internal/svc"
	"ZeZeIM/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FinduserbyidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用id查找用户
func NewFinduserbyidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FinduserbyidLogic {
	return &FinduserbyidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FinduserbyidLogic) Finduserbyid(req *types.FindUserByIDReq) (resp *types.FindUserByIDResp, err error) {
	// todo: add your logic here and delete this line

	return
}
