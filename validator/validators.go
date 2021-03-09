package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

//自定义手机号码验证器
func ValidateMobile(field validator.FieldLevel) bool {
	mobile := field.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.MatchString(`^1([39][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}