package ihttp

import (
	"github.com/gitkeng/ihttp/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log log message to console
func (ms *Microservice) Log(lvl LogLevel, message string, fields ...any) {
	zapLogLevel := zapcore.DebugLevel
	switch lvl {
	case DebugLevel:
		zapLogLevel = zapcore.DebugLevel
	case InfoLevel:
		zapLogLevel = zapcore.InfoLevel
	case WarnLevel:
		zapLogLevel = zapcore.WarnLevel
	case ErrorLevel:
		zapLogLevel = zapcore.ErrorLevel
	case DPanicLevel:
		zapLogLevel = zapcore.DPanicLevel
	case PanicLevel:
		zapLogLevel = zapcore.PanicLevel
	case FatalLevel:
		zapLogLevel = zapcore.FatalLevel
	}

	// convert fields to zap.Field
	zapFields := make([]zap.Field, 0)
	for _, field := range fields {
		if zapField, ok := field.(zap.Field); ok {
			zapFields = append(zapFields, zapField)
		}
	}
	ms.logger.Log(zapLogLevel, message, zapFields...)
}

func (ms *Microservice) Debug(message string) {
	ms.logger.Debug(message)
}
func (ms *Microservice) Debugf(format string, args ...any) {
	ms.logger.Debugf(format, args...)
}
func (ms *Microservice) Debugj(message string, key string, j log.JSON) {
	ms.logger.Debugj(message, key, j)
}
func (ms *Microservice) Info(message string) {
	ms.logger.Info(message)
}
func (ms *Microservice) Infof(format string, args ...any) {
	ms.logger.Infof(format, args...)
}
func (ms *Microservice) Infoj(message string, key string, j log.JSON) {
	ms.logger.Infoj(message, key, j)
}
func (ms *Microservice) Warn(message string) {
	ms.logger.Warn(message)
}
func (ms *Microservice) Warnf(format string, args ...any) {
	ms.logger.Warnf(format, args...)
}
func (ms *Microservice) Warnj(message string, key string, j log.JSON) {
	ms.logger.Warnj(message, key, j)
}
func (ms *Microservice) Error(message string) {
	ms.logger.Error(message)
}
func (ms *Microservice) Errorf(format string, args ...any) {
	ms.logger.Errorf(format, args...)
}
func (ms *Microservice) Errorj(message string, key string, j log.JSON) {
	ms.logger.Errorj(message, key, j)
}
func (ms *Microservice) Fatal(message string) {
	ms.logger.Fatal(message)
}
func (ms *Microservice) Fatalf(format string, args ...any) {
	ms.logger.Fatalf(format, args...)
}
func (ms *Microservice) Fatalj(message string, key string, j log.JSON) {
	ms.logger.Fatalj(message, key, j)
}
func (ms *Microservice) Panic(message string) {
	ms.logger.Panic(message)
}
func (ms *Microservice) Panicf(format string, args ...any) {
	ms.logger.Panicf(format, args...)
}
func (ms *Microservice) Panicj(message string, key string, j log.JSON) {
	ms.logger.Panicj(message, key, j)
}
