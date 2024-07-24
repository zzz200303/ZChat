package record

import (
	"ZChat/apps/group-chat/api/internal/logic/record"
	"ZChat/apps/group-chat/api/internal/svc"
	"ZChat/apps/group-chat/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecordListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecordListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := record.NewRecordListLogic(r.Context(), svcCtx)
		resp, err := l.RecordList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
