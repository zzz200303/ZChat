package svc

import (
	"ZChat/apps/group-chat/internal/config"
	"ZChat/apps/group-chat/internal/handler/chatconn"
)

type ServiceContext struct {
	Config config.Config
	*chatconn.Hub
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Hub:    chatconn.NewHub(),
	}
}
