package initialize

import (
	"cat-user-api/middlewares"
	"cat-user-api/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 用来给consul做健康检查使用
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1")
	routers.InitBaseRouter(ApiGroup)
	routers.InitUserRouter(ApiGroup)

	return Router
}
