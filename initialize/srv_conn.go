package initialize

import (
	"cat-user-api/proto"
	"fmt"

	"cat-user-api/global"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// 初始化 GRPC 服务
func InitSrvConn()  {
	// 从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	userSrvHost := ""
	userSrvPort := 0

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		// 只获取一个
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}

	zap.S().Debug("获取用户列表")
	//host := "127.0.0.1"
	//port := 50051
	//host := global.ServerConfig.UserSrvInfo.Host
	//port := global.ServerConfig.UserSrvInfo.Port
	host := userSrvHost
	port := userSrvPort
	// 1. 拨号连接用户 grpc 服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接【用户服务失败】",
			"msg", err.Error(),
		)
	}

	//todo 这里还有一些问题存在，比如1.后续的用户服务下线了，但是下面全局变量赋值了，这里就需要进行维护；2.端口变更了；3.IP变更了，上述问题在负载均衡的时候进行解决，这里不做解决

	// 这里已经事先建立连接，这里连接建立好了之后，在多个协程中使用，后续就不用再进行TCP的三次握手，此处性能相对来说较高了
	// TODO 但是还是有一些问题，比如一个连接多个 groutine 共用会不会存在性能问题，所以这里可以扩展成一个连接池 grpc connect pool
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}