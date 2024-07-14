package logic

import (
	"ZeZeIM/apps/social/rpc/internal/svc"
	"ZeZeIM/apps/social/rpc/pb/social"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInListLogic {
	return &FriendPutInListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInListLogic) FriendPutInList(in *social.FriendPutInListReq) (*social.FriendPutInListResp, error) {
	// todo: add your logic here and delete this line

	friendReqList, err := l.svcCtx.FriendRequestsModel.ListNoHandler(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "find list friend req err %v req %v", err, in.UserId)
	}

	var resp []*social.FriendRequests
	err = copier.Copy(&resp, &friendReqList)
	if err != nil {
		return nil, err
	}

	return &social.FriendPutInListResp{
		List: resp,
	}, nil

}
