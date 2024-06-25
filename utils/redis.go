package utils

import (
	"ZeZeIM/config"
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisAddr,
		Password: config.AppConfig.RedisPwd,
		DB:       config.AppConfig.RedisDB,
	})

	return rdb
}
