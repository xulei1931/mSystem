package utils

import (
	"go.uber.org/zap"
	"movie/common/logtool"
)

var logger *zap.Logger

func NewLoggerServer() {
	logger = logtool.NewLogger(
		logtool.SetAppName("go-kit"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
}

func GetLogger() *zap.Logger {
	return logger
}
