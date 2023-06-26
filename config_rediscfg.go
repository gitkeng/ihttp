package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/stringutil"
	"strings"
	"time"
)

// IRedisConfig is RedisCache configuration interface
type IRedisConfig interface {
	IConfig
	//GetContextName is the option for setting redis context name
	GetContextName() string
	//GetEndpoint is the option for setting redis server
	GetEndpoint() string
	//GetPassword is the option for setting redis password
	GetPassword() string
	//GetDB is the option for setting redis database
	GetDB() int
	//GetPoolSize is the option for setting redis dbConn pool size
	GetPoolSize() int
	//GetMinIdleConns is the option for setting redis dbConn minimum idle connections
	GetMinIdleConns() int
	//GetMaxRetries is the option for setting redis dbConn maximum retries
	GetMaxRetries() int
	//GetMinRetryBackoff is the option for setting redis dbConn minimum retry backoff
	GetMinRetryBackoff() time.Duration
	//GetMaxRetryBackoff is the option for setting redis dbConn maximum retry backoff
	GetMaxRetryBackoff() time.Duration
	//GetMaxIdleTime is the option for setting redis dbConn maximum idle time
	GetMaxIdleTime() time.Duration
	//GetPoolTimeout is the option for setting redis dbConn pool timeout
	GetPoolTimeout() time.Duration
	//GetReadTimeout is the option for setting redis dbConn read timeout
	GetReadTimeout() time.Duration
	//GetWriteTimeout is the option for setting redis dbConn write timeout
	GetWriteTimeout() time.Duration
}

type RedisConfig struct {
	//ContextName is the redis context name
	ContextName string `mapstructure:"context-name" json:"context_name"`
	//Endpoint is the redis server
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
	//Password is the redis password
	Password string `mapstructure:"password" json:"password"`
	// Database to be selected after connecting to the server.
	DB int `mapstructure:"db" json:"db"`
	// Maximum number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	PoolSize int `mapstructure:"pool-size" json:"pool_size"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int `mapstructure:"min-idle-conns" json:"min_idle_conns"`
	// Maximum number of retries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	MaxRetries int `mapstructure:"max-retries" json:"max_retries"`
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff time.Duration `mapstructure:"min-retry-backoff" json:"min_retry_backoff"`
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff time.Duration `mapstructure:"max-retry-backoff" json:"max_retry_backoff"`
	// MaxIdleTime is the maximum amount of time a connection may be idle.
	// Should be less than server's timeout.
	//
	// Expired connections may be closed lazily before reuse.
	// If d <= 0, connections are not closed due to a connection's idle time.
	//
	// Default is 30 minutes. -1 disables idle timeout check.
	MaxIdleTime time.Duration `mapstructure:"max-idle-time" json:"max_idle_time"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration `mapstructure:"pool-timeout" json:"pool_timeout"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetReadDeadline calls completely.
	ReadTimeout time.Duration `mapstructure:"read-timeout" json:"read_timeout"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.  Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetWriteDeadline calls completely.
	WriteTimeout time.Duration `mapstructure:"write-timeout" json:"write_timeout"`
}

func (cache *RedisConfig) Bind() error {
	cache.ContextName = strings.TrimSpace(cache.ContextName)
	cache.Endpoint = strings.TrimSpace(cache.Endpoint)
	cache.Password = strings.TrimSpace(cache.Password)
	if cache.DB < 0 {
		cache.DB = DefaultRedisCacheDB
	}

	if cache.PoolSize <= 0 {
		cache.PoolSize = DefaultRedisCachePoolSize
	}

	if cache.MinIdleConns <= 0 {
		cache.MinIdleConns = DefaultRedisCacheMinIdleConns
	}

	if cache.MaxRetries <= 0 &&
		cache.MaxRetries != -1 {
		cache.MaxRetries = DefaultRedisCacheMaxRetries
	}

	if cache.MinRetryBackoff < 8 &&
		cache.MinRetryBackoff != -1 {
		cache.MinRetryBackoff = DefaultRedisCacheMinRetryBackoff
	} else if cache.MinRetryBackoff != -1 {
		cache.MinRetryBackoff = cache.MinRetryBackoff * time.Millisecond
	}

	if cache.MaxRetryBackoff < 512 &&
		cache.MaxRetryBackoff != -1 {
		cache.MaxRetryBackoff = DefaultRedisCacheMaxRetryBackoff
	} else if cache.MaxRetryBackoff != -1 {
		cache.MaxRetryBackoff = cache.MaxRetryBackoff * time.Millisecond
	}

	if cache.MaxIdleTime < 30 &&
		cache.MaxIdleTime != -1 {
		cache.MaxIdleTime = DefaultRedisCacheMaxIdleTime
	} else if cache.MaxIdleTime != -1 &&
		cache.MaxIdleTime >= 30 {
		cache.MaxIdleTime = cache.MaxIdleTime * time.Minute

	}

	if cache.ReadTimeout <= 0 &&
		cache.ReadTimeout != -1 &&
		cache.ReadTimeout != -2 {
		cache.ReadTimeout = DefaultRedisCacheReadTimeout
	} else if cache.ReadTimeout != -1 &&
		cache.ReadTimeout != -2 {
		cache.ReadTimeout = cache.ReadTimeout * time.Second
	}

	if cache.PoolTimeout <= 0 {
		cache.PoolTimeout = DefaultRedisCachePoolTimeout
	} else {
		cache.PoolTimeout = cache.PoolTimeout * time.Second
	}

	if cache.WriteTimeout <= 0 &&
		cache.WriteTimeout != -1 &&
		cache.WriteTimeout != -2 {
		cache.WriteTimeout = DefaultRedisCacheWriteTimeout
	} else if cache.WriteTimeout != -1 &&
		cache.WriteTimeout != -2 {
		cache.WriteTimeout = cache.WriteTimeout * time.Second
	}
	return nil
}

func (cache *RedisConfig) Validate() error {
	if stringutil.IsEmptyString(cache.ContextName) {
		return ErrRedisContextNameIsRequire
	}
	if stringutil.IsEmptyString(cache.Endpoint) {
		return ErrRedisEndpointIsRequire
	}
	return nil
}

func (cache *RedisConfig) String() string {
	return stringutil.Json(cache)
}

func (cache *RedisConfig) ToMap() map[string]any {
	return convutil.Obj2Map(*cache)
}

func (cache *RedisConfig) GetContextName() string {
	return cache.ContextName
}

func (cache *RedisConfig) GetEndpoint() string {
	return cache.Endpoint
}

func (cache *RedisConfig) GetPassword() string {
	return cache.Password
}

func (cache *RedisConfig) GetDB() int {
	return cache.DB
}

func (cache *RedisConfig) GetPoolSize() int {
	return cache.PoolSize
}

func (cache *RedisConfig) GetMinIdleConns() int {
	return cache.MinIdleConns
}

func (cache *RedisConfig) GetMaxRetries() int {
	return cache.MaxRetries
}

func (cache *RedisConfig) GetMinRetryBackoff() time.Duration {
	return cache.MinRetryBackoff
}

func (cache *RedisConfig) GetMaxRetryBackoff() time.Duration {
	return cache.MaxRetryBackoff
}

func (cache *RedisConfig) GetMaxIdleTime() time.Duration {
	return cache.MaxIdleTime
}

func (cache *RedisConfig) GetPoolTimeout() time.Duration {
	return cache.PoolTimeout
}

func (cache *RedisConfig) GetReadTimeout() time.Duration {
	return cache.ReadTimeout
}

func (cache *RedisConfig) GetWriteTimeout() time.Duration {
	return cache.WriteTimeout
}
