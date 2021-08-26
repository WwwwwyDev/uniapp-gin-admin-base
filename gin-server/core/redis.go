package core

import (
	"gin-server/global"
	"github.com/go-redis/redis"
)

func Redis() *redis.Client{
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.LOG.Error(err)
	} else {
		global.LOG.Info("redis connect ping response:"+pong)
	}
	return client
}
