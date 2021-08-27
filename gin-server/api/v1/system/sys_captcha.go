package system

import (
	"gin-server/global"
	"gin-server/util/app"
	"gin-server/util/captcha"
	"gin-server/util/e"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)


// @Tags SysCaptcha
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /api/v1/sysCaptcha [post]
func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	store := captcha.NewDefaultRedisStore()
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.LOG.Error("验证码获取失败", err)
		app.Error(c,e.ERROR,err,err.Error()+"验证码获取失败")
	} else {
		app.OK(c,map[string]interface{}{"captchaId": id,"picPath":b64s},"验证码获取成功")
	}
}
