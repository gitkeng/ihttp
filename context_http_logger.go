package ihttp

import (
	"fmt"
	"github.com/gitkeng/ihttp/log"
	"github.com/gitkeng/ihttp/util/stringutil"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (ctx *HTTPContext) Logger() IContextLogger {
	return ctx
}

func (ctx *HTTPContext) defaultLoggerFields() []any {
	if ctx.ctx != nil {
		req := ctx.ctx.Request()
		res := ctx.ctx.Response()

		logFields := []any{
			zap.String("remote_ip", ctx.ctx.RealIP()),
			zap.String("host", req.Host),
			zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
			zap.Int64("size", res.Size),
			zap.String("user_agent", req.UserAgent()),
		}

		id := req.Header.Get(echo.HeaderXRequestID)
		if stringutil.IsEmptyString(id) {
			id = res.Header().Get(echo.HeaderXRequestID)
		}
		if stringutil.IsNotEmptyString(id) {
			logFields = append(logFields, zap.String(RequestIdField, id))

		}
		return logFields
	}
	return nil

}

// ==================== Implement IContextLogger ====================

// Log log message
func (ctx *HTTPContext) Log(level LogLevel, message string, fields ...any) {
	logFields := ctx.defaultLoggerFields()
	if len(logFields) == 0 {
		logFields = make([]any, 0)
	}
	for _, field := range fields {
		if zapField, ok := field.(zap.Field); ok {
			logFields = append(logFields, zapField)
		}
	}

	if ctx.ms != nil {
		if len(logFields) > 0 {
			ctx.ms.Log(level, message, logFields...)
		} else {
			ctx.ms.Log(level, message)
		}
	} else {
		zapLogLevel := zapcore.DebugLevel
		switch level {
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

		zapFields := make([]zap.Field, 0)
		for _, field := range fields {
			if zapField, ok := field.(zap.Field); ok {
				zapFields = append(zapFields, zapField)
			}
		}

		if len(zapFields) > 0 {
			log.Log(zapLogLevel, message, zapFields...)
		} else {
			log.Log(zapLogLevel, message)

		}
	}

}

// Debug log debug message
func (ctx *HTTPContext) Debug(message string) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(DebugLevel, message, zapFields...)
		} else {
			ctx.ms.Log(DebugLevel, message)
		}
	} else {
		log.Debug(message)
	}
}

// Debugf log debug message with format
func (ctx *HTTPContext) Debugf(format string, args ...any) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(DebugLevel, fmt.Sprintf(format, args...), zapFields...)
		} else {
			ctx.ms.Log(DebugLevel, fmt.Sprintf(format, args...))
		}
	} else {
		log.Debugf(format, args...)
	}
}

// Debugj log debug message with json
func (ctx *HTTPContext) Debugj(message string, key string, j log.JSON) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			zapFields = append(zapFields, zap.Any(key, j))
			ctx.ms.Log(DebugLevel, message, zapFields...)
		} else {
			ctx.ms.Log(DebugLevel, message, zap.Any(key, j))
		}
	} else {
		log.Debugj(message, key, j)
	}
}

// Info log info message
func (ctx *HTTPContext) Info(message string) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(InfoLevel, message, zapFields...)
		} else {
			ctx.ms.Log(InfoLevel, message)
		}
	} else {
		log.Info(message)
	}
}

// Infof log info message with format
func (ctx *HTTPContext) Infof(format string, args ...any) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(InfoLevel, fmt.Sprintf(format, args...), zapFields...)
		} else {
			ctx.ms.Log(InfoLevel, fmt.Sprintf(format, args...))
		}
	} else {
		log.Infof(format, args...)
	}
}

// Infoj log info message with json
func (ctx *HTTPContext) Infoj(message string, key string, j log.JSON) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			zapFields = append(zapFields, zap.Any(key, j))
			ctx.ms.Log(InfoLevel, message, zapFields...)
		} else {
			ctx.ms.Log(InfoLevel, message, zap.Any(key, j))
		}
	} else {
		log.Infoj(message, key, j)
	}
}

// Warn log warn message
func (ctx *HTTPContext) Warn(message string) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(WarnLevel, message, zapFields...)
		} else {
			ctx.ms.Log(WarnLevel, message)
		}
	} else {
		log.Warn(message)
	}
}

// Warnf log warn message with format
func (ctx *HTTPContext) Warnf(format string, args ...any) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(WarnLevel, fmt.Sprintf(format, args...), zapFields...)
		} else {
			ctx.ms.Log(WarnLevel, fmt.Sprintf(format, args...))
		}
	} else {
		log.Warnf(format, args...)
	}
}

// Warnj log warn message with json
func (ctx *HTTPContext) Warnj(message string, key string, j log.JSON) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			zapFields = append(zapFields, zap.Any(key, j))
			ctx.ms.Log(WarnLevel, message, zapFields...)
		} else {
			ctx.ms.Log(WarnLevel, message, zap.Any(key, j))
		}
	} else {
		log.Warnj(message, key, j)
	}
}

// Error log error message
func (ctx *HTTPContext) Error(message string) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(ErrorLevel, message, zapFields...)
		} else {
			ctx.ms.Log(ErrorLevel, message)
		}
	} else {
		log.Error(message)
	}
}

// Errorf log error message with format
func (ctx *HTTPContext) Errorf(format string, args ...any) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(ErrorLevel, fmt.Sprintf(format, args...), zapFields...)
		} else {
			ctx.ms.Log(ErrorLevel, fmt.Sprintf(format, args...))
		}
	} else {
		log.Errorf(format, args...)
	}
}

// Errorj log error message with json
func (ctx *HTTPContext) Errorj(message string, key string, j log.JSON) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			zapFields = append(zapFields, zap.Any(key, j))
			ctx.ms.Log(ErrorLevel, message, zapFields...)
		} else {
			ctx.ms.Log(ErrorLevel, message, zap.Any(key, j))
		}
	} else {
		log.Errorj(message, key, j)
	}
}

// Fatal log fatal message
func (ctx *HTTPContext) Fatal(message string) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(FatalLevel, message, zapFields...)
		} else {
			ctx.ms.Log(FatalLevel, message)
		}
	} else {
		log.Fatal(message)
	}
}

// Fatalf log fatal message with format
func (ctx *HTTPContext) Fatalf(format string, args ...any) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(FatalLevel, fmt.Sprintf(format, args...), zapFields...)
		} else {
			ctx.ms.Log(FatalLevel, fmt.Sprintf(format, args...))
		}
	} else {
		log.Fatalf(format, args...)
	}
}

// Fatalj log fatal message with json
func (ctx *HTTPContext) Fatalj(message string, key string, j log.JSON) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			zapFields = append(zapFields, zap.Any(key, j))
			ctx.ms.Log(FatalLevel, message, zapFields...)
		} else {
			ctx.ms.Log(FatalLevel, message, zap.Any(key, j))
		}
	} else {
		log.Fatalj(message, key, j)
	}
}

// Panic log panic message
func (ctx *HTTPContext) Panic(message string) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(PanicLevel, message, zapFields...)
		} else {
			ctx.ms.Log(PanicLevel, message)
		}
	} else {
		log.Panic(message)
	}
}

// Panicf log panic message with format
func (ctx *HTTPContext) Panicf(format string, args ...any) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			ctx.ms.Log(PanicLevel, fmt.Sprintf(format, args...), zapFields...)
		} else {
			ctx.ms.Log(PanicLevel, fmt.Sprintf(format, args...))
		}
	} else {
		log.Panicf(format, args...)
	}
}

// Panicj log panic message with json
func (ctx *HTTPContext) Panicj(message string, key string, j log.JSON) {
	if ctx.ms != nil {
		zapFields := ctx.defaultLoggerFields()
		if len(zapFields) > 0 {
			zapFields = append(zapFields, zap.Any(key, j))
			ctx.ms.Log(PanicLevel, message, zapFields...)
		} else {
			ctx.ms.Log(PanicLevel, message, zap.Any(key, j))
		}
	} else {
		log.Panicj(message, key, j)
	}
}
