package user

import (
	"net/http"

	"ZeZeIM/apps/user/api/internal/logic/user"
	"ZeZeIM/apps/user/api/internal/svc"
	"ZeZeIM/apps/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取所有用户
func AlluserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AllUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewAlluserLogic(r.Context(), svcCtx)
		resp, err := l.Alluser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
