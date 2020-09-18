package user_service

import (
	"go.uber.org/zap"
)

var log *zap.Logger
var sugar *zap.SugaredLogger

func Logger(env string) {
}

func Log() *zap.Logger {
	_ = log.Sync()
	return log
}

func Sugar() *zap.SugaredLogger {
	_ = sugar.Sync()
	return sugar
}