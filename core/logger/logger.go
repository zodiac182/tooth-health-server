// internal/logger/logger.go
package logger

import (
	"github.com/zodiac182/tooth-health/server/global"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	// 创建一个 AtomicLevel 对象，用于动态调整日志级别
	atomicLevel := zap.NewAtomicLevel()

	// 根据环境变量设置日志级别
	switch global.LogLevel {
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	case "fatal":
		atomicLevel.SetLevel(zap.FatalLevel)
	default:
		// default info level
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	// 创建 zap 配置并设置日志级别
	// config := zap.NewProductionConfig()
	config := zap.NewDevelopmentConfig()
	config.Level = atomicLevel // 设置动态级别

	if global.Mode == "production" {
		config.OutputPaths = []string{"stdout", "log/tooth-health.log"}
	}

	// 创建 logger
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// 使用 SugaredLogger 简化日志记录
	Log = logger.Sugar()
}

// 简化日志记录的包装函数
func Debug(msg string, args ...interface{}) {
	Log.Debugf(msg, args...)
}

func Info(msg string, args ...interface{}) {
	Log.Infof(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	Log.Warnf(msg, args...)
}

func Error(msg string, args ...interface{}) {
	Log.Errorf(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	Log.Fatalf(msg, args...)
}
