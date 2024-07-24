package chatconn

import (
	"ZChat/apps/group-chat/api/internal/svc"
	"ZChat/apps/group-chat/api/internal/types"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllGroupByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllGroupByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllGroupByUserIdLogic {
	return &GetAllGroupByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllGroupByUserIdLogic) GetAllGroupByUserId(req *types.GetAllGroupByUserIdRequest) (resp *types.GetAllGroupByUserIdResponse, err error) {
	uidJson := l.ctx.Value("uid").(json.Number) // 从jwt里面提取uid
	uid, err := uidJson.Int64()
	if err != nil {
		fmt.Println("json.Number换出了问题")
		return
	}
	gm, err := l.svcCtx.GmemberModel.FindAllGroupByUserId(l.ctx, uid)
	if err != nil {
		return nil, err
	}
	for _, v := range gm {
		resp.Groups = append(resp.Groups, v.Gid)
	}
	return resp, nil
}
