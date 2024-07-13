package handler

import (
	"net/http"

	"ZeZeIM/apps/social/api/internal/logic"
	"ZeZeIM/apps/social/api/internal/svc"
	"ZeZeIM/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 申请进群
func groupPutInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupPutInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGroupPutInLogic(r.Context(), svcCtx)
		resp, err := l.GroupPutIn(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
