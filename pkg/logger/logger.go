package logger

import (
	"sync"

	"github.com/Angstreminus/exchanger/pkg/config"
	"go.uber.org/zap"
)

type Logger struct {
	Zap *zap.Logger
}

var (
	Log  *Logger
	once sync.Once
)

func MustInitLogger(cfg *config.Config) *Logger {
	once.Do(
		func() {
			zlogger := zap.Must(zap.NewProduction())
			if cfg.LogLevel == "development" {
				zlogger = zap.Must(zap.NewDevelopment())
			}
			Log = &Logger{
				Zap: zlogger,
			}
		})
	return Log
}
