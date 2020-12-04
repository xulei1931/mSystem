package errors

import (
	"mSystem/src/common/config"
	"github.com/micro/go-micro/errors"
)


const (
	errorCodeUserSuccess = 200
	errorCodeUserError= 400
)

var (
	ErrorUserSuccess = errors.New(
		config.ServiceNameUser,"操作成功",errorCodeUserSuccess,
	)

	ErrorUserFailed = errors.New(
		config.ServiceNameUser,"操作异常",errorCodeUserSuccess,
	)

	ErrorUserAlready = errors.New(
		config.ServiceNameUser,"该邮箱已经被注册过了~",errorCodeUserSuccess,
	)

	ErrorUserLoginFailed = errors.New(
		config.ServiceNameUser,"密码或者用户名错误~",errorCodeUserSuccess,
	)

	ErrorTokenEmpty = errors.New(
		config.ServiceNameUser,"token不能为空",errorCodeUserError,
	)
	ErrorTokenError = errors.New(
		config.ServiceNameUser,"token验证错误",errorCodeUserError,
	)
)

 func FormatError (name,err string) error{
 	return errors.New(
		name,err,errorCodeUserError,
	 )
 }