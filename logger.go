package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   sync.Once
	logger *zap.Logger
)

func getLoggerMust() *zap.Logger {
	if logger == nil {
		panic("logger: logger is nil (forgotten register?)")
	}
	return logger
}

func WithOptions(opts ...zap.Option) *zap.Logger { return getLoggerMust().WithOptions(opts...) }

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return getLoggerMust().Check(lvl, msg)
}

func Core() zapcore.Core { return getLoggerMust().Core() }

func DPanic(msg string, fields ...zap.Field) { getLoggerMust().DPanic(msg, fields...) }

func Debug(msg string, fields ...zap.Field) { getLoggerMust().Debug(msg, fields...) }

func Error(msg string, fields ...zap.Field) { getLoggerMust().Error(msg, fields...) }

func Fatal(msg string, fields ...zap.Field) { getLoggerMust().Fatal(msg, fields...) }

func Info(msg string, fields ...zap.Field) { getLoggerMust().Info(msg, fields...) }

func Named(s string) *zap.Logger { return getLoggerMust().Named(s) }

func Panic(msg string, fields ...zap.Field) { getLoggerMust().Panic(msg, fields...) }

func Sugar() *zap.SugaredLogger { return getLoggerMust().Sugar() }

func Sync() error { return getLoggerMust().Sync() }

func Warn(msg string, fields ...zap.Field) { getLoggerMust().Warn(msg, fields...) }

func With(fields ...zap.Field) *zap.Logger { return getLoggerMust().With(fields...) }

func Register(l *zap.Logger) *zap.Logger { once.Do(func() { logger = l }) }
