package ihttp

import (
	"errors"
	"fmt"
)

var (
	//Microservice errors
	ErrNotValidatable             = Error{Code: "ERR_NOT_VALIDATABLE", Message: "object is not validatable"}
	ErrInvalidPort                = errors.New("port is invalid")
	ErrInvalidSSLPort             = errors.New("ssl port is invalid")
	ErrSSLTypeNotSupport          = func(sslType SSLType) error { return fmt.Errorf("ssl type: %s not support", sslType) }
	ErrSSLCertificateFileRequire  = errors.New("ssl certificate file is require")
	ErrSSLCertificateFileNotfound = func(certFile string) error { return fmt.Errorf("ssl certificate file: %s not found", certFile) }
	ErrSSLKeyFileRequire          = errors.New("ssl key file is require")
	ErrSSLKeyFileNotfound         = func(keyFile string) error { return fmt.Errorf("ssl key file: %s not found", keyFile) }

	ErrDBConfigsIsRequire    = errors.New("database configs is require")
	ErrLogConfigIsRequire    = errors.New("log config is require")
	ErrRedisConfigsIsRequire = errors.New("redis configs is require")
	ErrAPIConfigIsRequire    = errors.New("api config is require")

	//Log Config errors
	ErrInvalidLogLevel         = func(level string) error { return fmt.Errorf("log level is invalid: %s" + level) }
	ErrInvalidLogfileLocation  = func(location string) error { return fmt.Errorf("log file location is invalid: %s", location) }
	ErrInvalidLogfileMaxSize   = func(size int) error { return fmt.Errorf("log file max size is invalid: %d", size) }
	ErrInvalidLogfileMaxBackup = func(backup int) error { return fmt.Errorf("log file max backup is invalid: %d", backup) }
	ErrInvalidLogfileMaxAge    = func(age int) error { return fmt.Errorf("log file max age is invalid: %d", age) }

	//Database Config errors
	ErrInvalidDBProvider      = func(provider string) error { return fmt.Errorf("database provider is invalid: %s", provider) }
	ErrDBContextNameIsRequire = errors.New("database context name is required")
	ErrDBURLRequire           = errors.New("database url is require")
	ErrDBUserRequire          = errors.New("database user is require")
	ErrDBPasswordRequire      = errors.New("database password is require")
	ErrDBNameRequire          = errors.New("database name is require")
	ErrDBURLPattern           = func(url string) error { return fmt.Errorf("database invalid url [%s] pattern <db uri>:<port>", url) }
	ErrDuplicateDBContextName = func(name string) error { return fmt.Errorf("database context name [%s] is duplicate", name) }

	//RedisCache Config errors
	ErrDuplicateRedisContextName = func(name string) error { return fmt.Errorf("redis context name [%s] is duplicate", name) }
	ErrRedisContextNameIsRequire = errors.New("redis context name is required")
	ErrRedisEndpointIsRequire    = errors.New("redis endpoint is require")
	ErrRedisContextNameNotfound  = func(name string) error { return fmt.Errorf("redis context name [%s] not found", name) }
)
