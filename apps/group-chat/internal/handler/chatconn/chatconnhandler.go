package chatconn

import (
	"net/http"

	"ZChat/apps/group-chat/internal/logic/chatconn"
	"ZChat/apps/group-chat/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatConnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chatconn.NewChatConnLogic(r.Context(), svcCtx)
		err := l.ChatConn()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
