package forms

type PasswordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式需要自定义validator
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}
