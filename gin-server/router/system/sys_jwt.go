package system

import (
	"gin-server/api/v1/system"
	"github.com/gin-gonic/gin"
)

func SysJwtRouter(r *gin.RouterGroup)  {
	r = r.Group("/jwt")
	r.POST("", system.DecodeJwt)
	r.DELETE("",system.DelJwt)
}
