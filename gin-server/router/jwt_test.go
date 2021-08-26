package router

import (
	"gin-server/api/v1"
	"github.com/gin-gonic/gin"
)

func jwtTestRouter(r *gin.RouterGroup)  {
	r = r.Group("/jwt")
	r.POST("", )
}
