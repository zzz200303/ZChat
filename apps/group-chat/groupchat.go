package main

import (
	"ZChat/apps/group-chat/internal/config"
	"ZChat/apps/group-chat/internal/handler"
	"ZChat/apps/group-chat/internal/handler/chatconn"
	"ZChat/apps/group-chat/internal/svc"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/groupchat-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	hub := chatconn.NewHub()
	go hub.Run()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.Start()
}
