package ihttp

import (
	"github.com/gitkeng/ihttp/util/stringutil"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Option func(mc *Microservice) error

// WithAPIConfig is the option for setting api config
func WithAPIConfig(config IAPIConfig) Option {
	return func(ms *Microservice) error {
		if config == nil {
			return ErrAPIConfigIsRequire
		}
		if err := config.Bind(); err != nil {
			return err
		}
		if err := config.Validate(); err != nil {
			return err
		}

		ms.port = config.GetPort()
		ms.healthCheckEndpoint = config.GetHealthCheckEndpoint()
		ms.httpsOnly = config.IsHttpsOnly()
		if ms.httpsOnly {
			ms.sslEnable = true
		} else {
			ms.sslEnable = config.IsSSLEnable()
		}
		ms.sslPort = config.GetSSLPort()
		ms.sslCertFile = config.GetSSLCertFile()
		ms.sslKeyFile = config.GetSSLKeyFile()

		ms.apiConfig = config
		return nil
	}
}

// WithLogConfig is the option for setting log config
func WithLogConfig(config ILogConfig) Option {
	return func(ms *Microservice) error {
		if config == nil {
			return ErrLogConfigIsRequire
		}
		if err := config.Bind(); err != nil {
			return err
		}
		if err := config.Validate(); err != nil {
			return err
		}

		location := strings.TrimSpace(config.GetFileLocation())
		if stringutil.IsEmptyString(location) {
			ms.logFileLocation = ""
			ms.logFileEnable = false
		} else {
			ms.logFileLocation = location
			ms.logFileEnable = true
		}

		logLevel := config.GetLogLevel()
		switch logLevel {
		case DebugLevel:
			ms.logLevel = zapcore.DebugLevel
		case InfoLevel:
			ms.logLevel = zapcore.InfoLevel
		case WarnLevel:
			ms.logLevel = zapcore.WarnLevel
		case ErrorLevel:
			ms.logLevel = zapcore.ErrorLevel
		case DPanicLevel:
			ms.logLevel = zapcore.DPanicLevel
		case PanicLevel:
			ms.logLevel = zapcore.PanicLevel
		case FatalLevel:
			ms.logLevel = zapcore.FatalLevel
		default:
			return ErrInvalidLogLevel(string(logLevel))
		}

		size := config.GetMaxSize()
		if size <= 0 {
			return ErrInvalidLogfileMaxSize(size)
		}
		ms.logFileMaxSize = size

		backup := config.GetMaxBackups()
		if backup <= 0 {
			return ErrInvalidLogfileMaxBackup(backup)
		}
		ms.logFileMaxBackups = backup

		age := config.GetMaxAge()
		if age <= 0 {
			return ErrInvalidLogfileMaxAge(age)
		}
		ms.logFileMaxAge = age
		ms.logConfig = config

		return nil
	}
}

// WithDBConfigs is the option for setting database config
func WithDBConfigs(confs ...IDBConfig) Option {
	return func(ms *Microservice) error {
		if confs == nil || len(confs) == 0 {
			return ErrDBConfigsIsRequire
		}

		if len(confs) > 0 {
			if ms.dbConfigs == nil {
				ms.dbConfigs = make(map[string]IDBConfig, 0)
			}
			for idx, _ := range confs {
				if err := confs[idx].Bind(); err != nil {
					return err
				}
				if err := confs[idx].Validate(); err != nil {
					return err
				}
				ms.dbConfigs[confs[idx].GetContextName()] = confs[idx]
			}
		}

		return nil
	}
}

// WithRedisConfigs is the option for setting redis config
func WithRedisConfigs(confs ...IRedisConfig) Option {
	return func(ms *Microservice) error {
		if confs == nil || len(confs) == 0 {
			return ErrRedisConfigsIsRequire
		}

		if len(confs) > 0 {
			if ms.redisConfigs == nil {
				ms.redisConfigs = make(map[string]IRedisConfig, 0)
			}
			for idx, _ := range confs {
				if err := confs[idx].Bind(); err != nil {
					return err
				}
				if err := confs[idx].Validate(); err != nil {
					return err
				}
				ms.redisConfigs[confs[idx].GetContextName()] = confs[idx]
			}
		}
		return nil
	}
}

// WithHealthChecks is the option for setting health check functions
func WithHealthChecks(healthFuncs ...HealthCheckFunc) Option {
	return func(ms *Microservice) error {
		for idx, _ := range healthFuncs {
			ms.healthCheckFuncs = append(ms.healthCheckFuncs, healthFuncs[idx])
		}
		return nil
	}
}

// WithMiddleWares is the option for setting middlewares
func WithMiddleWares(middleWares ...echo.MiddlewareFunc) Option {
	return func(ms *Microservice) error {
		for idx, _ := range middleWares {
			ms.middlewares = append(ms.middlewares, middleWares[idx])
		}
		return nil
	}
}

func WithCleanupFuncs(cleanupFuncs ...CleanupFunc) Option {
	return func(ms *Microservice) error {
		for idx, _ := range cleanupFuncs {
			ms.cleanupFuncs = append(ms.cleanupFuncs, cleanupFuncs[idx])
		}
		return nil
	}
}
