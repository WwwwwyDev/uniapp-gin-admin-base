package middleware

import (
	"errors"
	"gin-server/global"
	"gin-server/util"
	"gin-server/util/app"
	"gin-server/util/e"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 j-token
		token := c.Request.Header.Get("j-token")
		if token == ""{
			err := errors.New("未携带j-token")
			global.LOG.Error(err)
			app.Error(c,e.ERROR,err,err.Error())
		}
		result, err := global.REDIS.Exists(token).Result()
		if err != nil {
			global.LOG.Error(err)
			app.Error(c,e.ERROR,err,err.Error())
		}
		if result == 0{
			err := errors.New("token过期")
			global.LOG.Error(err)
			app.Error(c,e.ERROR,err,err.Error())
		}
		secret, err1 := global.REDIS.Get(token).Result()
		if err1 != nil {
			global.LOG.Error(err1)
			app.Error(c,e.ERROR,err1,err1.Error())
		}
		claims := util.ParseToken(token, secret)
		c.Set("claims", claims)
		c.Next()
	}
}
