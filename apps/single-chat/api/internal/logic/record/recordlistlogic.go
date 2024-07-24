package record

import (
	"ZChat/apps/single-chat/api/internal/svc"
	"ZChat/apps/single-chat/api/internal/types"
	"ZChat/pkg/constants"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type RecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordListLogic {
	return &RecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordListLogic) RecordList(req *types.RecordListRequest) (resp *types.RecordListResponse, err error) {
	pattern1 := fmt.Sprintf("%s:from:%d:to:%d", constants.SingleChatMsg, req.FriendUid, req.Uid) // 生成Redis键
	recodeJsonList1, err := l.svcCtx.Redis.LrangeCtx(l.ctx, pattern1, 0, -1)
	pattern2 := fmt.Sprintf("%s:from:%d:to:%d", constants.SingleChatMsg, req.Uid, req.FriendUid) // 生成Redis键
	recodeJsonList2, err := l.svcCtx.Redis.LrangeCtx(l.ctx, pattern2, 0, -1)
	var combinedList []string
	// 追加第一个列表的所有元素到新切片
	combinedList = append(combinedList, recodeJsonList1...)
	combinedList = append(combinedList, recodeJsonList2...)

	if err != nil {
		logx.Errorf("l.svcCtx.Redis.KeysCtx error: %v", err)
		return nil, err
	}

	var messages []types.MessageInfo

	for _, recodeJson := range combinedList {
		m := types.MessageInfo{}
		fmt.Println(recodeJson)
		err := json.Unmarshal([]byte(recodeJson), &m)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		messages = append(messages, m)
	}

	return &types.RecordListResponse{
		RecordList: messages,
	}, err
}
