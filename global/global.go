package global

import (
	"cat-user-api/config"
)

var (
	// 因为其他地方要来改变这个变量，所以设置为指针类型
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)