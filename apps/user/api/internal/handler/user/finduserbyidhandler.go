package user

import (
	"net/http"

	"ZeZeIM/apps/user/api/internal/logic/user"
	"ZeZeIM/apps/user/api/internal/svc"
	"ZeZeIM/apps/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用id查找用户
func FinduserbyidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindUserByIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewFinduserbyidLogic(r.Context(), svcCtx)
		resp, err := l.Finduserbyid(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
