package initialize

import (
	"cat-user-api/middlewares"
	"cat-user-api/routers"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1")
	routers.InitBaseRouter(ApiGroup)
	routers.InitUserRouter(ApiGroup)

	return Router
}