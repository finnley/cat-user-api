package routers

import (
	"cat-user-api/api"
	"github.com/gin-gonic/gin"
)

//用户路由
func InitUserRouter(Router *gin.RouterGroup)  {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", api.GetUserList)
	}
}