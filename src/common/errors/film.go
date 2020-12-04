package errors

import (
	"mSystem/src/common/config"
	"github.com/micro/go-micro/errors"
)


const (
	errorCodeFilmSuccess = 200

)

var (
	ErrorFilmFailed = errors.New(
		config.ServiceNameUser,"操作异常",errorCodeFilmSuccess,
	)
)
