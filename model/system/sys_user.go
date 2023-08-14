package system

import "gomap/utils"

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

// GetMessages 自定义错误消息
func (register Register) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"name.required":     "用户名不能为空",
		"mobile.required":   "手机号不能为空",
		"mobile.mobile":     "手机号格式错误",
		"password.required": "密码不能为空",
	}
}
