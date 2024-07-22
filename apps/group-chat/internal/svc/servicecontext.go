package svc

import (
	"ZChat/apps/group-chat/internal/config"
	"ZChat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Redis          *redis.Redis
	UserRpcService userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRpcService: userclient.NewUser(zrpc.MustNewClient(c.UserRpcService)),
	}
}
