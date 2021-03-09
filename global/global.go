package global

import (
	"cat-user-api/config"
	ut "github.com/go-playground/universal-translator"
)

var (
	Trans ut.Translator
	// 因为其他地方要来改变这个变量，所以设置为指针类型
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
