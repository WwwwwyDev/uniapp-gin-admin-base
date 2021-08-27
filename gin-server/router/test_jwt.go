package router

import (
	v1 "gin-server/api/v1"
	"github.com/gin-gonic/gin"
)

func TestJwtRouter(r *gin.RouterGroup)  {
	r = r.Group("/jwt")
	r.POST("", v1.TestJwt)
}
