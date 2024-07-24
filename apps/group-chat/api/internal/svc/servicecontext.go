package svc

import (
	"ZChat/apps/group-chat/api/internal/config"
	"ZChat/apps/group-chat/model"
	"ZChat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Redis    *redis.Redis
	KqPusher *kq.Pusher
	//RecodesModel   model.RecodesModel
	GmemberModel   model.GmemberModel
	UserRpcService userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{

		Config: c,
		//RecodesModel: model.NewRecodesModel(sqlx.NewMysql(c.Mysql.DataSource)),
		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		GmemberModel:   model.NewGmemberModel(sqlx.NewMysql(c.Mysql.DataSource)),
		KqPusher:       kq.NewPusher(c.GroupMsgKqConf.Brokers, c.GroupMsgKqConf.Topic),
		UserRpcService: userclient.NewUser(zrpc.MustNewClient(c.UserRpcService)),
	}
}
