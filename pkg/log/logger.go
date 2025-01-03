package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Logger() *zap.Logger {
	once.Do(func() {
		fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		stdOutEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		filePath := os.Getenv("LOG_FILE_PATH")
		var fileCore zapcore.Core
		if filePath != "" {
			file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			fileCore = zapcore.NewCore(fileEncoder, file, zap.LevelEnablerFunc(func(level zapcore.Level) bool {
				return level >= zapcore.InfoLevel
			}))
		}
		stdoutCore := zapcore.NewCore(stdOutEncoder, os.Stdout, zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zapcore.DebugLevel
		}))
		var core zapcore.Core
		if fileCore != nil {
			core = zapcore.NewTee(stdoutCore, fileCore)
		} else {
			core = stdoutCore
		}
		logger = zap.New(core, zap.AddCallerSkip(1))
	})
	return logger
}
