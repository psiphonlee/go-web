package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

// GetErrorMsg 获取错误消息
func GetErrorMsg(req interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := req.(Validator)

		for _, v := range err.(validator.ValidationErrors) {
			// 若req 结构体实现Validator接口即可实现自定义错误消息
			if isValidator {
				if message, exist := req.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return message
				}
			}
			return v.Error()
		}
	}
	return "Parameter error"
}

// ValidateMobile 校验手机号
func ValidateMobile(fd validator.FieldLevel) bool {
	mobile := fd.Field().String()
	ok,_ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}