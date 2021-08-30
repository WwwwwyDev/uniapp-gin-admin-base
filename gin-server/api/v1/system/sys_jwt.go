package system

import (
	"errors"
	"gin-server/global"
	service "gin-server/service/system"
	"gin-server/util"
	"gin-server/util/app"
	"gin-server/util/e"
	"github.com/gin-gonic/gin"
)

// @Tags SysJwt
// @Summary jwt解码,反馈用户信息
// @Produce  application/json
// @Success 200 {string} string "{"code": 20000,"data":{},"msg":"解码成功"}"
// @Router /api/v1/c/jwt [post]
func DecodeJwt(c *gin.Context) {
	get, _ := c.Get("claims")
	if get == nil{
		app.Error(c, e.ERROR, errors.New("jwt签发人校验失败"), "jwt签发人校验失败")
		return
	}
	claims := get.(*util.Claims)
	if claims.Issuer != global.CONFIG.Jwt.Issuer{
		app.Error(c, e.ERROR, errors.New("jwt签发人校验失败"), "jwt签发人校验失败")
		return
	}
	if claims.IpAddress != c.ClientIP(){
		app.Error(c, e.ERROR, errors.New("异常的token"), "异常的token")
		return
	}
	err, inter := service.GetSysUserById(claims.UserID)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	inter.Password = "敏感信息，已屏蔽"
	inter.Salt = "敏感信息，已屏蔽"
	app.OK(c, inter,"解码成功")
}


// @Tags SysJwt
// @Summary 删除用户token
// @Produce  application/json
// @Success 200 {string} string "{"code": 20000,"data":{},"msg":"删除成功"}"
// @Router /api/v1/c/jwt [DELETE]
func DelJwt(c *gin.Context) {
	token := c.Request.Header.Get("j-token")
	_, err := global.REDIS.Del("SYSUSER_"+token).Result()
	if err != nil {
		global.LOG.Error(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "删除成功")
}
