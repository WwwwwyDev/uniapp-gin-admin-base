package core

import (
	"context"
	"gin-server/global"
	"github.com/go-redis/redis/v8"
)

func Redis() *redis.Client{
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error(err)
	} else {
		global.LOG.Info("redis connect ping response:"+pong)
	}
	return client
}
