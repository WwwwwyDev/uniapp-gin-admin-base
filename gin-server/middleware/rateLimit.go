package middleware

import (
	"errors"
	"gin-server/global"
	"gin-server/util/app"
	"gin-server/util/e"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		IP := "RATE_"+ c.ClientIP()
		result, err := global.REDIS.Exists(IP).Result()
		if err != nil {
			global.LOG.Error(err)
			app.Error(c,e.ERROR,err,err.Error())
			c.Abort()
			return
		}
		if result == 0{
			_, err := global.REDIS.Set(IP, 0, time.Second*10).Result()
			if err != nil {
				global.LOG.Error(err)
				app.Error(c,e.ERROR,err,err.Error())
				c.Abort()
				return
			}
		}else{
			s, err := global.REDIS.Get(IP).Result()
			if err != nil {
				global.LOG.Error(err)
				app.Error(c,e.ERROR,err,err.Error())
				c.Abort()
				return
			}
			satoi, err1 := strconv.Atoi(s)
			if err1 != nil {
				global.LOG.Error(err1)
				app.Error(c,e.ERROR,err1,err1.Error())
				c.Abort()
				return
			}
			if satoi < 15{
				duration, err := global.REDIS.TTL(IP).Result()
				if err != nil {
					global.LOG.Error(err)
					app.Error(c,e.ERROR,err,err.Error())
					c.Abort()
					return
				}
				if duration == time.Second*-1{

					_, err := global.REDIS.Set(IP, 0, time.Second*10).Result()
					if err != nil {
						global.LOG.Error(err)
						app.Error(c,e.ERROR,err,err.Error())
						c.Abort()
						return
					}
				}else{
					_, err1 := global.REDIS.Set(IP, satoi+1, duration).Result()
					if err1 != nil {
						global.LOG.Error(err1)
						app.Error(c,e.ERROR,err1,err1.Error())
						c.Abort()
						return
					}
				}
			}else {
				app.Error(c,e.ERROR_TIME_LIMIT,errors.New("访问频率过高"),"访问频率过高")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}