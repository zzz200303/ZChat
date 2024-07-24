package chatconn

import (
	"context"

	"ZChat/apps/group-chat/api/internal/svc"
	"ZChat/apps/group-chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserByGroupIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUserByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserByGroupIdLogic {
	return &GetAllUserByGroupIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUserByGroupIdLogic) GetAllUserByGroupId(req *types.GetAllUserByGroupIdRequest) (resp *types.GetAllUserByGroupIdResponse, err error) {
	gm, err := l.svcCtx.GmemberModel.FindAllUserByGroupId(l.ctx, req.GroupId)
	if err != nil {
		return nil, err
	}
	res := new(types.GetAllUserByGroupIdResponse)
	for _, v := range gm {
		res.Users = append(res.Users, v.Uid)
	}
	return res, nil
}
