package captcha

import (
	"context"
	"gin-server/global"
	"time"
)

func NewDefaultRedisStore() RedisStore {
	return RedisStore{
		Expiration: time.Second * 100,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}


func (rs RedisStore) Set(id string, value string) error{
	err := global.REDIS.Set(rs.PreKey+id, value, rs.Expiration).Err()
	return err
}

func (rs RedisStore) Get(key string, clear bool) string {
	val, err := global.REDIS.Get(rs.PreKey+key).Result()
	if err != nil {
		global.LOG.Error("RedisStoreGetError!", err)
		return ""
	}
	if clear {
		err := global.REDIS.Del(rs.PreKey+key).Err()
		if err != nil {
			global.LOG.Error("RedisStoreClearError!", err)
			return ""
		}
	}
	return val
}

func (rs RedisStore) Verify(id, answer string, clear bool) bool {
	key := id
	v := rs.Get(key, clear)
	return v == answer
}
