package main

import (
	"ZChat/apps/group-chat/api/internal/config"
	"ZChat/apps/group-chat/api/internal/handler"
	"ZChat/apps/group-chat/api/internal/handler/chatconn"
	"ZChat/apps/group-chat/api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

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

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//初始化所有用户
	err := chatconn.InitAllUser(ctx)
	if err != nil {
		logx.Error(err)
		return
	}
	//初始化kafka消费者
	go func(ctx *svc.ServiceContext) {
		chatconn.StartMq(ctx)
	}(ctx)

	//初始化新用户检测消费者
	go func(ctx *svc.ServiceContext) {
		chatconn.StartNewUserMq(ctx)
	}(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
