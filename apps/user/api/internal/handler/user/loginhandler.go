package user

import (
	"net/http"
	"fmt"

	"ZChat/apps/user/api/internal/logic/user"
	"ZChat/apps/user/api/internal/svc"
	"ZChat/apps/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登入
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	fmt.Println("1")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("2")
		var req types.LoginReq
		fmt.Println("3")
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("4")
			httpx.ErrorCtx(r.Context(), w, err)
			fmt.Println("5")
			return
		}
		fmt.Println("6")
		l := user.NewLoginLogic(r.Context(), svcCtx)
		fmt.Println("7")
		resp, err := l.Login(&req)
		fmt.Println("8")
		if err != nil {
			fmt.Println("9")
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
