package main

import (
	"cat-user-api/global"
	"cat-user-api/initialize"
	"cat-user-api/utils"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	myvalidator "cat-user-api/validator"
)

func main() {
	//port := 8021

	// 初始化 logger
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()

	// 初始化 routers
	Router := initialize.Routers()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// 初始化 srv 连接
	initialize.InitSrvConn()

	port := global.ServerConfig.Port
	viper.AutomaticEnv()
	// 如果是本地开发环境，端口是需要固定,要不然端口一变就要更改yapi，线上环境自动获取端口号
	debug := viper.GetBool("CAT_DEBUG")
	if !debug {
		tcpPort, err := utils.GetFreeTcpPort()
		if err == nil {
			port = tcpPort
		}
	}

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		// 翻译
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}为非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	/**
	1. S() 可以获取一个全局的 sugar,可以让我们自己设置一个全局的 logger
	2. 日志是分级别的，debug, info, warn, error, fetal
	3. S 函数和 L 函数，它提供了一个全局的安全访问 logger 的途径
	*/
	zap.S().Debugf("启动服务器, 端口: %d", port)
	//启动服务
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}
}
