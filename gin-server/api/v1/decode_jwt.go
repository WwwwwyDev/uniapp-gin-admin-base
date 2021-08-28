package v1

import (
	"errors"
	"gin-server/global"
	service "gin-server/service/system"
	"gin-server/util"
	"gin-server/util/app"
	"gin-server/util/e"
	"github.com/gin-gonic/gin"
)

// @Tags DecodeJwt
// @Summary jwt解码,反馈用户信息
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"解码成功"}"
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
	err, inter := service.GetSysUserById(claims.UserID)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	app.OK(c, inter,"解码成功")
}
