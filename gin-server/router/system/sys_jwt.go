package system

import (
	"gin-server/api/v1/system"
	"github.com/gin-gonic/gin"
)

func SysJwtRouter(r *gin.RouterGroup)  {
	r.POST("/decodeJwt", system.DecodeJwt)
	r.POST("/isValidJwt", system.IsValidJwt)
	r.DELETE("/jwt",system.DelJwt)
}
