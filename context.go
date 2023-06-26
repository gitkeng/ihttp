package ihttp

import (
	"github.com/gitkeng/ihttp/log"
	"github.com/labstack/echo/v4"
	"time"
)

// IContext is the context for service
type IContext interface {
	Log(level LogLevel, message string, fields ...any)
	Logger() IContextLogger
	Param(name string) string
	QueryParam(name string) string
	Response(logLevel LogLevel, logTag string, httpStatus int, code, message string, err error, fields ...Field) error
	ReadRequest() string
	ReadRequests() []string
	WebContext() echo.Context
	Bind(request any) error

	// Now return current time
	Now() time.Time
	Epoch() int64

	//DB return the DBStore
	DB(contextName string) (IDBStore, bool)
	//Cache return the RedisCache
	Cache(contextName string) (IRedisCache, bool)

	//Requester return Requester
	Requester(baseURL string, timeout time.Duration, certFiles ...string) (IRequester, error)
	//Config return config

	//APIConfig return APIConfig
	APIConfig() (IAPIConfig, bool)
	//LogConfig return LogConfig
	LogConfig() (ILogConfig, bool)
	//DBConfig return DBConfig
	DBConfig(dbContextName string) (IDBConfig, bool)
	//RedisConfig return RedisConfig
	RedisConfig(cacheContextName string) (IRedisConfig, bool)
}

type IContextLogger interface {
	Log(lvl LogLevel, msg string, fields ...any)
	Debug(message string)
	Debugf(format string, args ...any)
	Debugj(message string, key string, j log.JSON)
	Info(message string)
	Infof(format string, args ...any)
	Infoj(message string, key string, j log.JSON)
	Warn(message string)
	Warnf(format string, args ...any)
	Warnj(message string, key string, j log.JSON)
	Error(message string)
	Errorf(format string, args ...any)
	Errorj(message string, key string, j log.JSON)
	Fatal(message string)
	Fatalf(format string, args ...any)
	Fatalj(message string, key string, j log.JSON)
	Panic(message string)
	Panicf(format string, args ...any)
	Panicj(message string, key string, j log.JSON)
}
