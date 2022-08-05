package logger

import "go.uber.org/zap"

var Zap *zap.Logger

func InitLog() {
	log, _ := zap.NewProduction()
	Zap = log
}

func SyncLog() {
	Zap.Sync()
}

func Sugar() *zap.SugaredLogger {
	return Zap.Sugar()
}
