package api

import (
	"cat-user-api/global"
	"cat-user-api/global/response"
	"cat-user-api/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

//将grpc的code转换成http的状态码
func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误" + e.Message(),
				})
			}
			return
		}
	}
}

func GetUserList(ctx *gin.Context) {
	zap.S().Debug("获取用户列表")
	//host := "127.0.0.1"
	//port := 50051
	host := global.ServerConfig.UserSrvInfo.Host
	port := global.ServerConfig.UserSrvInfo.Port
	// 1. 拨号连接用户 grpc 服务器
	userCon, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接【用户服务失败】",
			"msg", err.Error(),
		)
	}
	// 2. 生成grpc的client并调用接口
	userSrvClient := proto.NewUserClient(userCon)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Page:     0,
		PageSize: 0,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询【用户列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	// 返回数据
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		//生成map
		//data := make(map[string]interface{})
		//data["id"] = value.Id
		//data["name"] = value.Nickname
		//data["birthday"] = value.Birthday
		//data["gender"] = value.Gender
		//data["mobile"] = value.Mobile
		//result = append(result, data)
		//改写
		user := response.UserResponse{
			Id:       value.Id,
			Nickname: value.Nickname,
			//Birthday: time.Time(time.Unix(int64(value.Birthday), 0)),
			//方法一：
			//Birthday: time.Time(time.Unix(int64(value.Birthday), 0)).Format("2006 01-02"),
			//方法二：
			Birthday: response.JsonTime(time.Unix(int64(value.Birthday), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}
