package system

import (
	"gin-server/api/v1/system"
	"github.com/gin-gonic/gin"
)

func SysCaptchaRouter(r *gin.RouterGroup)  {
	r.POST("/sysCaptcha", system.Captcha)
}

