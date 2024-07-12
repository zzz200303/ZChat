package svc

import (
	"ZeZeIM/apps/user/models"
	"ZeZeIM/apps/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
