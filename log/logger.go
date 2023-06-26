package log

import (
	"fmt"
	"github.com/gitkeng/ihttp/util/stringutil"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var instance ILogger

type (
	Option  func(*ZapLogger) error
	ILogger interface {
		Log(lvl zapcore.Level, msg string, fields ...zap.Field)
		Debug(message string, fields ...zap.Field)
		Debugf(format string, args ...any)
		Debugj(message string, key string, j JSON)
		Info(message string, fields ...zap.Field)
		Infof(format string, args ...any)
		Infoj(message string, key string, j JSON)
		Warn(message string, fields ...zap.Field)
		Warnf(format string, args ...any)
		Warnj(message string, key string, j JSON)
		Error(message string, fields ...zap.Field)
		Errorf(format string, args ...any)
		Errorj(message string, key string, j JSON)
		Fatal(message string, fields ...zap.Field)
		Fatalf(format string, args ...any)
		Fatalj(message string, key string, j JSON)
		Panic(message string, fields ...zap.Field)
		Panicf(format string, args ...any)
		Panicj(message string, key string, j JSON)
	}

	ZapLogger struct {
		logger            *zap.Logger
		logFileEnable     bool
		logFileLocation   string
		logFileMaxSize    int
		logFileMaxBackups int
		logFileMaxAge     int
		logLevel          zapcore.Level
	}

	JSON map[string]any
)

var (
	ErrInvalidLogLevel = func(level zapcore.Level) error {
		return fmt.Errorf("log level is invalid: %d", level)
	}
	ErrInvalidLogfileLocation = func(location string) error {
		return fmt.Errorf("log file location is invalid: %s", location)
	}
	ErrInvalidLogfileMaxSize = func(size int) error {
		return fmt.Errorf("log file max size is invalid: %d", size)
	}
	ErrInvalidLogfileMaxBackup = func(backup int) error {
		return fmt.Errorf("log file max backup is invalid: %d", backup)
	}
	ErrInvalidLogfileMaxAge = func(age int) error {
		return fmt.Errorf("log file max age is invalid: %d", age)
	}
)

// WithFileLocation is the option for setting log file location
func WithFileLocation(location string) Option {
	return func(zapLogger *ZapLogger) error {
		if stringutil.IsEmptyString(location) {
			zapLogger.logFileEnable = false
			return ErrInvalidLogfileLocation(location)
		}
		zapLogger.logFileLocation = location
		zapLogger.logFileEnable = true
		return nil
	}
}

// WithLevel is the option for setting log level
func WithLevel(level zapcore.Level) Option {
	return func(zapLogger *ZapLogger) error {
		switch level {
		case zapcore.DebugLevel,
			zapcore.InfoLevel,
			zapcore.WarnLevel,
			zapcore.ErrorLevel,
			zapcore.DPanicLevel,
			zapcore.PanicLevel,
			zapcore.FatalLevel:
			zapLogger.logLevel = level
			return nil
		default:
			return ErrInvalidLogLevel(level)
		}
		return nil
	}
}

// WithFileMaxSize is the option for setting log file max size (megabytes)
func WithFileMaxSize(size int) Option {
	return func(zapLogger *ZapLogger) error {
		if size <= 0 {
			return ErrInvalidLogfileMaxSize(size)
		}
		zapLogger.logFileMaxSize = size
		return nil
	}
}

// WithFileMaxBackups is the option for setting log file max backup
func WithFileMaxBackups(backup int) Option {
	return func(zapLogger *ZapLogger) error {
		if backup <= 0 {
			return ErrInvalidLogfileMaxBackup(backup)
		}
		zapLogger.logFileMaxBackups = backup
		return nil
	}
}

// WithFileMaxAge is the option for setting log file max age (days)
func WithFileMaxAge(age int) Option {
	return func(zapLogger *ZapLogger) error {
		if age <= 0 {
			return ErrInvalidLogfileMaxAge(age)
		}
		zapLogger.logFileMaxAge = age
		return nil
	}
}

func New(options ...Option) (*ZapLogger, error) {
	zapLogger := &ZapLogger{
		logFileEnable:     false,
		logLevel:          zapcore.DebugLevel,
		logFileMaxSize:    DefaultLogFileMaxSize,
		logFileMaxBackups: DefaultLogFileMaxBackups,
		logFileMaxAge:     DefaultLogFileMaxAge,
	}
	for _, setter := range options {
		if setter != nil {
			err := setter(zapLogger)
			if err != nil {
				return nil, err
			}
		}

	}

	//initial logger
	zapStdOutCfg := zap.NewProductionEncoderConfig()
	zapStdOutCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	zapStdOutCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapStdOutCfg.TimeKey = "timestamp"

	zapFileCfg := zap.NewProductionEncoderConfig()
	zapFileCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	zapFileCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	zapFileCfg.TimeKey = "timestamp"

	fe := zapcore.NewJSONEncoder(zapFileCfg)
	ce := zapcore.NewConsoleEncoder(zapStdOutCfg)

	//// lumberjack.Logger is already safe for concurrent use, so we don't need to
	//// lock it.

	if zapLogger.logFileEnable {
		logFile := zapcore.AddSync(&lumberjack.Logger{
			Filename:   zapLogger.logFileLocation,
			MaxSize:    zapLogger.logFileMaxSize, // megabytes
			MaxBackups: zapLogger.logFileMaxBackups,
			MaxAge:     zapLogger.logFileMaxAge, // days
		})
		logWriter := zapcore.AddSync(logFile)
		logCore := zapcore.NewTee(
			zapcore.NewCore(fe, logWriter, zapLogger.logLevel),
			zapcore.NewCore(ce, zapcore.AddSync(colorable.NewColorableStdout()), zapLogger.logLevel),
		)
		logger := zap.New(logCore, zap.AddStacktrace(zapcore.ErrorLevel))
		zapLogger.logger = logger
	} else {
		logCore := zapcore.NewTee(
			zapcore.NewCore(ce, zapcore.AddSync(colorable.NewColorableStdout()), zapLogger.logLevel),
		)
		logger := zap.New(logCore, zap.AddStacktrace(zapcore.ErrorLevel))
		zapLogger.logger = logger
	}
	instance = zapLogger
	return zapLogger, nil
}

func (log *ZapLogger) Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	log.logger.Log(lvl, msg, fields...)
}

func (log *ZapLogger) Debug(message string, fields ...zap.Field) {
	log.Log(zapcore.DebugLevel, message, fields...)
}

func (log *ZapLogger) Debugf(message string, args ...any) {
	log.Log(zapcore.DebugLevel, fmt.Sprintf(message, args...))
}

func (log *ZapLogger) Debugj(message string, key string, j JSON) {
	log.Log(zapcore.DebugLevel, message, zap.Any(key, j))
}

func (log *ZapLogger) Info(message string, fields ...zap.Field) {
	log.Log(zapcore.InfoLevel, message, fields...)
}

func (log *ZapLogger) Infof(message string, args ...any) {
	log.Log(zapcore.InfoLevel, fmt.Sprintf(message, args...))
}

func (log *ZapLogger) Infoj(message string, key string, j JSON) {
	log.Log(zapcore.InfoLevel, message, zap.Any(key, j))
}

func (log *ZapLogger) Warn(message string, fields ...zap.Field) {
	log.Log(zapcore.WarnLevel, message, fields...)
}

func (log *ZapLogger) Warnf(message string, args ...any) {
	log.Log(zapcore.WarnLevel, fmt.Sprintf(message, args...))
}

func (log *ZapLogger) Warnj(message string, key string, j JSON) {
	log.Log(zapcore.WarnLevel, message, zap.Any(key, j))
}

func (log *ZapLogger) Error(message string, fields ...zap.Field) {
	log.Log(zapcore.ErrorLevel, message, fields...)
}

func (log *ZapLogger) Errorf(message string, args ...any) {
	log.Log(zapcore.ErrorLevel, fmt.Sprintf(message, args...))
}

func (log *ZapLogger) Errorj(message string, key string, j JSON) {
	log.Log(zapcore.ErrorLevel, message, zap.Any(key, j))
}

func (log *ZapLogger) Fatal(message string, fields ...zap.Field) {
	log.Log(zapcore.FatalLevel, message, fields...)
}

func (log *ZapLogger) Fatalf(message string, args ...any) {
	log.Log(zapcore.FatalLevel, fmt.Sprintf(message, args...))
}

func (log *ZapLogger) Fatalj(message string, key string, j JSON) {
	log.Log(zapcore.FatalLevel, message, zap.Any(key, j))
}

func (log *ZapLogger) Panic(message string, fields ...zap.Field) {
	log.Log(zapcore.PanicLevel, message, fields...)
}

func (log *ZapLogger) Panicf(message string, args ...any) {
	log.Log(zapcore.PanicLevel, fmt.Sprintf(message, args...))
}

func (log *ZapLogger) Panicj(message string, key string, j JSON) {
	log.Log(zapcore.PanicLevel, message, zap.Any(key, j))
}

func Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	initialInstance()
	instance.Log(lvl, msg, fields...)
}

// Debug output message of debug level
func Debug(message string) {
	initialInstance()
	instance.Debug(message)

}

// Debugf output format message of debug level
func Debugf(format string, args ...interface{}) {
	initialInstance()
	instance.Debugf(format, args...)

}

// Debugj output json of debug level
func Debugj(message string, key string, j JSON) {
	initialInstance()
	instance.Debugj(message, key, j)

}

// Info output message of info level
func Info(message string, fields ...zap.Field) {
	initialInstance()
	instance.Info(message, fields...)

}

// Infof output format message of info level
func Infof(format string, args ...any) {
	initialInstance()
	instance.Infof(format, args...)
}

// Infoj output json of info level
func Infoj(message string, key string, j JSON) {
	initialInstance()
	instance.Infoj(message, key, j)
}

// Warn output message of warn level
func Warn(message string, fields ...zap.Field) {
	initialInstance()
	instance.Warn(message, fields...)
}

// Warnf output format message of warn level
func Warnf(format string, args ...any) {
	initialInstance()
	instance.Warnf(format, args...)
}

// Warnj output json of warn level
func Warnj(message string, key string, j JSON) {
	initialInstance()
	instance.Warnj(message, key, j)
}

// Error output message of error level
func Error(message string, fields ...zap.Field) {
	initialInstance()
	instance.Error(message, fields...)
}

// Errorf output format message of error level
func Errorf(format string, args ...any) {
	initialInstance()
	instance.Errorf(format, args...)
}

// Errorj output json of error level
func Errorj(message string, key string, j JSON) {
	initialInstance()
	instance.Errorj(message, key, j)
}

// Fatal output message of fatal level
func Fatal(message string, fields ...zap.Field) {
	initialInstance()
	instance.Fatal(message, fields...)
}

// Fatalf output format message of fatal level
func Fatalf(format string, args ...any) {
	initialInstance()
	instance.Fatalf(format, args...)
}

// Fatalj output json of fatal level
func Fatalj(message string, key string, j JSON) {
	initialInstance()
	instance.Fatalj(message, key, j)
}

// Panic output message of panic level
func Panic(message string, fields ...zap.Field) {
	initialInstance()
	instance.Panic(message, fields...)
}

// Panicf output format message of panic level
func Panicf(format string, args ...any) {
	initialInstance()
	instance.Panicf(format, args...)
}

// Panicj output json of panic level
func Panicj(message string, key string, j JSON) {
	initialInstance()
	instance.Fatalj(message, key, j)
}

func initialInstance() {
	if instance == nil {
		zapLogger, err := New()
		if err != nil {
			panic(err)
		}
		instance = zapLogger
	}
}
