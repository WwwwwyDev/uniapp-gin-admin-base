package middleware

import (
	"errors"
	"gin-server/global"
	"gin-server/util"
	"gin-server/util/app"
	"gin-server/util/e"
	"github.com/gin-gonic/gin"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 j-token
		token := c.Request.Header.Get("j-token")
		if token == ""{
			err := errors.New("未携带j-token")
			app.Error(c,e.ERROR,err,err.Error())
			return
		}
		result, err := global.REDIS.Exists("SYSUSER_"+token).Result()
		if err != nil {
			global.LOG.Error(err)
			app.Error(c,e.ERROR,err,err.Error())
			return
		}
		if result == 0{
			err := errors.New("无效的j-token")
			global.LOG.Error(err)
			app.Error(c,e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT,err,err.Error())
			return
		}
		secret, err1 := global.REDIS.Get("SYSUSER_"+token).Result()
		if err1 != nil {
			global.LOG.Error(err1)
			app.Error(c,e.ERROR,err1,err1.Error())
			return
		}
		duration, err2 := global.REDIS.TTL("SYSUSER_"+token).Result()
		if err2 != nil {
			global.LOG.Error(err2)
			app.Error(c,e.ERROR,err2,err2.Error())
			return
		}
		if duration.Nanoseconds() < time.Second.Nanoseconds()*100{
			//jwt延寿
			_, err := global.REDIS.Set("SYSUSER_"+token, secret, time.Second*1800).Result()
			if err != nil {
				global.LOG.Error(err)
				app.Error(c,e.ERROR,err,err.Error())
				return
			}
		}
		claims := util.ParseToken(token, secret)
		if claims == nil {
			err := errors.New("jwt声明解析错误")
			app.Error(c,e.ERROR_AUTH_CHECK_TOKEN_FAIL,err,err.Error())
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
