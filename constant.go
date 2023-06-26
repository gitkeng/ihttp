package ihttp

import "time"

const (
	// DefaultPort is the default port for the service
	DefaultPort                int    = 8080
	DefaultHealthCheckEndpoint string = "/health"
	DefaultSSLPort             int    = 8443

	//	DefaultLogFileMaxSize is the default max size of log file in MB
	DefaultLogFileMaxSize int = 500
	// DefaultLogFileMaxBackups is the default max number of log file backups
	DefaultLogFileMaxBackups int = 3
	// DefaultLogFileMaxAge is the default max age of log file in days
	DefaultLogFileMaxAge int = 30

	// DefaultDBGetRetryLimits is the default timeout for creating db dbConn loop
	DefaultDBGetRetryLimits int = 30
	//DefaultDBConnectionMaxOpenConns 0 (unlimited).
	DefaultDBConnectionMaxOpenConns int = 0
	//DefaultDBConnectionMaxIdleConns <= 0, no idle connections are retained. MaxIdleConns will be reduced to match the MaxOpenConns limit.
	DefaultDBConnectionMaxIdleConns int = 0
	//DefaultDBConnectionMaxLifeTime <= 0, connections are not closed due to a dbConn's age.
	DefaultDBConnectionMaxLifeTime int = 0

	//	Default for redis cache setting
	DefaultRedisCacheDB              int = 0
	DefaultRedisCachePoolSize        int = 10
	DefaultRedisCacheMinIdleConns    int = 5
	DefaultRedisCacheMaxRetries      int = 3
	DefaultRedisCacheMinRetryBackoff     = 8 * time.Millisecond
	DefaultRedisCacheMaxRetryBackoff     = 512 * time.Millisecond
	DefaultRedisCacheMaxIdleTime         = 30 * time.Minute
	DefaultRedisCacheReadTimeout         = 3 * time.Second
	DefaultRedisCachePoolTimeout         = DefaultRedisCacheReadTimeout + time.Second
	DefaultRedisCacheWriteTimeout        = 3 * time.Second
)

const (
	RequestIdField = "request_id"
)

type LogLevel string

const (
	//Loglevel
	DebugLevel  LogLevel = "debug"
	InfoLevel   LogLevel = "info"
	WarnLevel   LogLevel = "warn"
	ErrorLevel  LogLevel = "error"
	PanicLevel  LogLevel = "panic"
	DPanicLevel LogLevel = "dpanic"
	FatalLevel  LogLevel = "fatal"
)
