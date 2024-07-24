package chatconn

import (
	"context"
	"fmt"

	"ZChat/apps/group-chat/api/internal/svc"
	"ZChat/apps/group-chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuitGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuitGroupLogic {
	return &QuitGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuitGroupLogic) QuitGroup(req *types.QuitGroupRequest) (resp *types.QuitGroupResponse, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.GmemberModel.QuietGroup(l.ctx, req.Uid, req.Gid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &types.QuitGroupResponse{Response: "Success"}, nil
}
