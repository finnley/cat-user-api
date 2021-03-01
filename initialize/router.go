package initialize

import (
	"cat-user-api/routers"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/u/v1")
	routers.InitUserRouter(ApiGroup)

	return Router
}