package initialize

import (
	"cat-user-api/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 通过设置环境变量，本地如果想要生效 需要重启 goland
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig()  {
	debug := GetEnvInfo("CAT_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-prod.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	// 设置文件路径
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// serverConfig这个对象是在其他文件中使用到的-所以它是一个全局变量
	//serverConfig := config.ServerConfig{}
	//if err := v.Unmarshal(&serverConfig); err != nil {
	//	panic(err)
	//}
	//fmt.Println(serverConfig)
	//fmt.Println("%V", v.Get("name"))
	//
	////viper还可以动态监控变化
	//v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("config file changed: ", e.Name)
	//	_ = v.ReadInConfig()
	//	//重新读取配置文件
	//	_ = v.Unmarshal(&serverConfig)
	//	fmt.Println(serverConfig)
	//})

	// 修改
	// serverConfig这个对象是在其他文件中使用到的-所以它是一个全局变量
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global.ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化: %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息: %v", global.ServerConfig)
	})
}