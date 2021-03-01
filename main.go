package main

import (
	"cat-user-api/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	port := 8021

	// 初始化 logger
	initialize.InitLogger()

	// 初始化 routers
	Router := initialize.Routers()

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
