// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	chatconn "ZChat/apps/single-chat/api/internal/handler/chatconn"
	record "ZChat/apps/single-chat/api/internal/handler/record"
	"ZChat/apps/single-chat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/chatconn",
				Handler: chatconn.ChatConnHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/singlechat"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/recordlist",
				Handler: record.RecordListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/singlechat"),
	)
}