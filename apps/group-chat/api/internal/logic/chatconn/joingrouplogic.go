package chatconn

import (
	"ZChat/apps/group-chat/api/internal/svc"
	"ZChat/apps/group-chat/api/internal/types"
	"ZChat/apps/group-chat/model"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinGroupLogic) JoinGroup(req *types.JoinGroupRequest) (resp *types.JoinGroupResponse, err error) {
	uidJson := l.ctx.Value("uid").(json.Number) // 从jwt里面提取uid
	uid, err := uidJson.Int64()
	if err != nil {
		fmt.Println("json.Number换出了问题")
		return
	}
	_, err = l.svcCtx.GmemberModel.Insert(l.ctx, &model.Gmember{
		Gid: req.GroupId,
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	return &types.JoinGroupResponse{Response: "成功加入群聊"}, nil
}
