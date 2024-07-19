package user

import (
	"net/http"

	"ZeZeIM/apps/user/api/internal/logic/user"
	"ZeZeIM/apps/user/api/internal/svc"
	"ZeZeIM/apps/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用name查找用户
func FinduserbynameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindUserByNameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewFinduserbynameLogic(r.Context(), svcCtx)
		resp, err := l.Finduserbyname(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
