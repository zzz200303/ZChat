package chatconn

//
//import (
//	"ZChat/apps/group-chat/internal/types"
//	"encoding/json"
//	"fmt"
//	"net/http"
//
//	"ZChat/apps/group-chat/internal/svc"
//)
//
//func ChatConnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var u = types.User{}
//		uidJson := r.Context().Value("uid").(json.Number) // 从jwt里面提取uid
//		uid, err := uidJson.Int64()
//		if err != nil {
//			fmt.Println("json.Number换出了问题")
//			return
//		}
//		name := r.Context().Value("name").(string) // 从jwt里面提取uid
//		if err != nil {
//			fmt.Println("json.Number换出了问题")
//			return
//		}
//		u.Id = uid
//		u.Name = name
//
//		r.Context()
//		serveWs(svcCtx.Hub, w, r, u)
//	}
//}
