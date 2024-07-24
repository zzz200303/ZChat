package chatconn

import (
	"net/http"

	"ZChat/apps/group-chat/api/internal/logic/chatconn"
	"ZChat/apps/group-chat/api/internal/svc"
	"ZChat/apps/group-chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllUserByGroupIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllUserByGroupIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chatconn.NewGetAllUserByGroupIdLogic(r.Context(), svcCtx)
		resp, err := l.GetAllUserByGroupId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
