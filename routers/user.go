package routers

import (
	"cat-user-api/api"
	"cat-user-api/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//用户路由
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	zap.S().Info("配置用户相关url")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
		UserRouter.POST("register", api.Register)
	}
}
