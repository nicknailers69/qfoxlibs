package logger

import "go.uber.org/zap"

var Log *zap.Logger

type Logger *zap.SugaredLogger

func init() {
	Log, _ := zap.NewProduction()
	defer Log.Sync()

}

func NewLogger() Logger {
	return Log.Sugar()
}
