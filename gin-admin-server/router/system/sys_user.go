package system

import (
	"gin-server/api/v1/system"
	"github.com/gin-gonic/gin"
)

func SysUserRouter(r *gin.RouterGroup)  {
	r = r.Group("/sysUser")
	r.POST("/register", system.Register)
	r.POST("/login", system.Login)
	r.POST("/changePassword", system.ChangePassword)
}
