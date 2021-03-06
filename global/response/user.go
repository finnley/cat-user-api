package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

// 在生成 JsonTime 类型是调用json时内部自动MarshalJSON方法
func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006 01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       uint32 `json:"id"`
	Nickname string `json:"nickname"`
	//birthday 需要格式换换
	//Birthday time.Time `json:"birthday"`
	//方法一：
	//Birthday string `json:"birthday"`
	//方法二：
	Birthday JsonTime `json:"birthday"`
	Gender uint32 `json:"gender"`
	Mobile string `json:"mobile"`
}
