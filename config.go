package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/stringutil"
	"sync"
	"time"
)

type IConfig interface {
	// Validate validate config
	Validate() error
	// String return string representation of config
	String() string
	// Bind process correct config value
	Bind() error
	//ToMap return map representation of config
	ToMap() map[string]any
}

type Config struct {
	logConf       ILogConfig
	apiConf       IAPIConfig
	dbConfs       map[string]IDBConfig
	redisConfs    map[string]IRedisConfig
	dbConfsMutex  sync.RWMutex
	logCfgMutex   sync.RWMutex
	apiCfgMutex   sync.RWMutex
	redisCfgMutex sync.RWMutex
}

func NewConfig(confFile string) (*Config, error) {
	conf := &Config{
		dbConfs:       make(map[string]IDBConfig),
		redisConfs:    make(map[string]IRedisConfig),
		dbConfsMutex:  sync.RWMutex{},
		logCfgMutex:   sync.RWMutex{},
		redisCfgMutex: sync.RWMutex{},
		apiCfgMutex:   sync.RWMutex{},
	}

	err := conf.apiConfigLoader(confFile)
	if err != nil {
		return nil, err
	}

	err = conf.logConfigLoader(confFile)
	if err != nil {
		return nil, err
	}

	err = conf.dbConfigLoader(confFile)
	if err != nil {
		return nil, err
	}

	err = conf.redisConfigLoader(confFile)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (conf *Config) Validate() error {
	if conf.apiConf != nil {
		if err := conf.apiConf.Validate(); err != nil {
			return err
		}
	}

	if conf.logConf != nil {
		if err := conf.logConf.Validate(); err != nil {
			return err
		}
	}

	if len(conf.dbConfs) > 0 {
		for _, dbConf := range conf.dbConfs {
			if err := dbConf.Validate(); err != nil {
				return err
			}
		}
	}

	if len(conf.redisConfs) > 0 {
		for _, redisConf := range conf.redisConfs {
			if err := redisConf.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

type ConfigPrinter struct {
	APIConfig struct {
		Port                int    `json:"port"`
		HealthCheckEndpoint string `json:"health-check-endpoint"`
	} `json:"api_config"`

	LogConfig struct {
		FileLocation string   `json:"file_location"`
		MaxSize      int      `json:"max_size"`
		MaxBackups   int      `json:"max_backups"`
		MaxAge       int      `json:"max_age"`
		LogLevel     LogLevel `json:"log_level"`
	} `json:"log_config"`

	AsyncTaskConfig struct {
		Partition    int `json:"partition"`
		Replication  int `json:"replication"`
		RetainPeriod int `json:"retain_period"`
	} `json:"async_task_config"`

	DBConfigs map[string]struct {
		ContextName           string   `json:"context_name"`
		Provider              string   `json:"provider"`
		URL                   string   `json:"url"`
		User                  string   `json:"user"`
		Password              string   `json:"password"`
		DatabaseName          string   `json:"database_name"`
		RetryLimits           int      `json:"retry_limits"`
		ConnectionMaxLifeTime int      `json:"connection_max_life_time"`
		MaxIdleConns          int      `json:"max_idle_conns"`
		MaxOpenConns          int      `json:"max_open_conns"`
		InitialScripts        []string `json:"initial_scripts"`
	} `json:"db_configs"`

	RedisConfigs map[string]struct {
		ContextName     string        `json:"context_name"`
		Endpoint        string        `json:"endpoint"`
		Password        string        `json:"password"`
		DB              int           `json:"db"`
		PoolSize        int           `json:"pool_size"`
		MinIdleConns    int           `json:"min_idle_conns"`
		MaxRetries      int           `json:"max_retries"`
		MinRetryBackoff time.Duration `json:"min_retry_backoff"`
		MaxRetryBackoff time.Duration `json:"max_retry_backoff"`
		MaxIdleTime     time.Duration `json:"max_idle_time"`
		PoolTimeout     time.Duration `json:"pool_timeout"`
		ReadTimeout     time.Duration `json:"read_timeout"`
		WriteTimeout    time.Duration `json:"write_timeout"`
	} `json:"redis_configs"`
}

func (conf *Config) String() string {
	printer := ConfigPrinter{}
	if conf.apiConf != nil {
		printer.APIConfig.Port = conf.apiConf.GetPort()
		printer.APIConfig.HealthCheckEndpoint = conf.apiConf.GetHealthCheckEndpoint()
	}

	if conf.logConf != nil {
		printer.LogConfig.FileLocation = conf.logConf.GetFileLocation()
		printer.LogConfig.MaxSize = conf.logConf.GetMaxSize()
		printer.LogConfig.MaxBackups = conf.logConf.GetMaxBackups()
		printer.LogConfig.MaxAge = conf.logConf.GetMaxAge()
		printer.LogConfig.LogLevel = conf.logConf.GetLogLevel()
	}

	if len(conf.dbConfs) > 0 {
		printer.DBConfigs = make(map[string]struct {
			ContextName           string   `json:"context_name"`
			Provider              string   `json:"provider"`
			URL                   string   `json:"url"`
			User                  string   `json:"user"`
			Password              string   `json:"password"`
			DatabaseName          string   `json:"database_name"`
			RetryLimits           int      `json:"retry_limits"`
			ConnectionMaxLifeTime int      `json:"connection_max_life_time"`
			MaxIdleConns          int      `json:"max_idle_conns"`
			MaxOpenConns          int      `json:"max_open_conns"`
			InitialScripts        []string `json:"initial_scripts"`
		})

		for key, dbConf := range conf.dbConfs {
			printer.DBConfigs[key] = struct {
				ContextName           string   `json:"context_name"`
				Provider              string   `json:"provider"`
				URL                   string   `json:"url"`
				User                  string   `json:"user"`
				Password              string   `json:"password"`
				DatabaseName          string   `json:"database_name"`
				RetryLimits           int      `json:"retry_limits"`
				ConnectionMaxLifeTime int      `json:"connection_max_life_time"`
				MaxIdleConns          int      `json:"max_idle_conns"`
				MaxOpenConns          int      `json:"max_open_conns"`
				InitialScripts        []string `json:"initial_scripts"`
			}{
				ContextName:           dbConf.GetContextName(),
				Provider:              dbConf.GetProvider(),
				URL:                   dbConf.GetURL(),
				User:                  dbConf.GetUser(),
				Password:              dbConf.GetPassword(),
				DatabaseName:          dbConf.GetDatabaseName(),
				RetryLimits:           dbConf.GetRetryLimits(),
				ConnectionMaxLifeTime: dbConf.GetConnectionMaxLifeTime(),
				MaxIdleConns:          dbConf.GetMaxIdleConns(),
				MaxOpenConns:          dbConf.GetMaxOpenConns(),
				InitialScripts:        dbConf.GetInitialScripts(),
			}
		}
	}

	if len(conf.redisConfs) > 0 {
		printer.RedisConfigs = make(map[string]struct {
			ContextName     string        `json:"context_name"`
			Endpoint        string        `json:"endpoint"`
			Password        string        `json:"password"`
			DB              int           `json:"db"`
			PoolSize        int           `json:"pool_size"`
			MinIdleConns    int           `json:"min_idle_conns"`
			MaxRetries      int           `json:"max_retries"`
			MinRetryBackoff time.Duration `json:"min_retry_backoff"`
			MaxRetryBackoff time.Duration `json:"max_retry_backoff"`
			MaxIdleTime     time.Duration `json:"max_idle_time"`
			PoolTimeout     time.Duration `json:"pool_timeout"`
			ReadTimeout     time.Duration `json:"read_timeout"`
			WriteTimeout    time.Duration `json:"write_timeout"`
		})

		for key, redisConf := range conf.redisConfs {
			printer.RedisConfigs[key] = struct {
				ContextName     string        `json:"context_name"`
				Endpoint        string        `json:"endpoint"`
				Password        string        `json:"password"`
				DB              int           `json:"db"`
				PoolSize        int           `json:"pool_size"`
				MinIdleConns    int           `json:"min_idle_conns"`
				MaxRetries      int           `json:"max_retries"`
				MinRetryBackoff time.Duration `json:"min_retry_backoff"`
				MaxRetryBackoff time.Duration `json:"max_retry_backoff"`
				MaxIdleTime     time.Duration `json:"max_idle_time"`
				PoolTimeout     time.Duration `json:"pool_timeout"`
				ReadTimeout     time.Duration `json:"read_timeout"`
				WriteTimeout    time.Duration `json:"write_timeout"`
			}{
				ContextName:     redisConf.GetContextName(),
				Endpoint:        redisConf.GetEndpoint(),
				Password:        redisConf.GetPassword(),
				DB:              redisConf.GetDB(),
				PoolSize:        redisConf.GetPoolSize(),
				MinIdleConns:    redisConf.GetMinIdleConns(),
				MaxRetries:      redisConf.GetMaxRetries(),
				MinRetryBackoff: redisConf.GetMinRetryBackoff(),
				MaxRetryBackoff: redisConf.GetMaxRetryBackoff(),
				MaxIdleTime:     redisConf.GetMaxIdleTime(),
				PoolTimeout:     redisConf.GetPoolTimeout(),
				ReadTimeout:     redisConf.GetReadTimeout(),
				WriteTimeout:    redisConf.GetWriteTimeout(),
			}
		}
	}

	return stringutil.Json(printer)
}

func (conf *Config) Bind() error {
	if conf.apiConf != nil {
		if err := conf.apiConf.Bind(); err != nil {
			return err
		}
	}

	if conf.logConf != nil {
		if err := conf.logConf.Bind(); err != nil {
			return err
		}
	}

	if len(conf.dbConfs) > 0 {
		for _, dbConf := range conf.dbConfs {
			if err := dbConf.Bind(); err != nil {
				return err
			}
		}
	}

	if len(conf.redisConfs) > 0 {
		for _, redisConf := range conf.redisConfs {
			if err := redisConf.Bind(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (conf *Config) ToMap() map[string]any {
	printer := ConfigPrinter{}
	if conf.apiConf != nil {
		printer.APIConfig.Port = conf.apiConf.GetPort()
		printer.APIConfig.HealthCheckEndpoint = conf.apiConf.GetHealthCheckEndpoint()
	}

	if conf.logConf != nil {
		printer.LogConfig.FileLocation = conf.logConf.GetFileLocation()
		printer.LogConfig.MaxSize = conf.logConf.GetMaxSize()
		printer.LogConfig.MaxBackups = conf.logConf.GetMaxBackups()
		printer.LogConfig.MaxAge = conf.logConf.GetMaxAge()
		printer.LogConfig.LogLevel = conf.logConf.GetLogLevel()
	}

	if len(conf.dbConfs) > 0 {
		printer.DBConfigs = make(map[string]struct {
			ContextName           string   `json:"context_name"`
			Provider              string   `json:"provider"`
			URL                   string   `json:"url"`
			User                  string   `json:"user"`
			Password              string   `json:"password"`
			DatabaseName          string   `json:"database_name"`
			RetryLimits           int      `json:"retry_limits"`
			ConnectionMaxLifeTime int      `json:"connection_max_life_time"`
			MaxIdleConns          int      `json:"max_idle_conns"`
			MaxOpenConns          int      `json:"max_open_conns"`
			InitialScripts        []string `json:"initial_scripts"`
		})

		for key, dbConf := range conf.dbConfs {
			printer.DBConfigs[key] = struct {
				ContextName           string   `json:"context_name"`
				Provider              string   `json:"provider"`
				URL                   string   `json:"url"`
				User                  string   `json:"user"`
				Password              string   `json:"password"`
				DatabaseName          string   `json:"database_name"`
				RetryLimits           int      `json:"retry_limits"`
				ConnectionMaxLifeTime int      `json:"connection_max_life_time"`
				MaxIdleConns          int      `json:"max_idle_conns"`
				MaxOpenConns          int      `json:"max_open_conns"`
				InitialScripts        []string `json:"initial_scripts"`
			}{
				ContextName:           dbConf.GetContextName(),
				Provider:              dbConf.GetProvider(),
				URL:                   dbConf.GetURL(),
				User:                  dbConf.GetUser(),
				Password:              dbConf.GetPassword(),
				DatabaseName:          dbConf.GetDatabaseName(),
				RetryLimits:           dbConf.GetRetryLimits(),
				ConnectionMaxLifeTime: dbConf.GetConnectionMaxLifeTime(),
				MaxIdleConns:          dbConf.GetMaxIdleConns(),
				MaxOpenConns:          dbConf.GetMaxOpenConns(),
				InitialScripts:        dbConf.GetInitialScripts(),
			}
		}
	}

	if len(conf.redisConfs) > 0 {
		printer.RedisConfigs = make(map[string]struct {
			ContextName     string        `json:"context_name"`
			Endpoint        string        `json:"endpoint"`
			Password        string        `json:"password"`
			DB              int           `json:"db"`
			PoolSize        int           `json:"pool_size"`
			MinIdleConns    int           `json:"min_idle_conns"`
			MaxRetries      int           `json:"max_retries"`
			MinRetryBackoff time.Duration `json:"min_retry_backoff"`
			MaxRetryBackoff time.Duration `json:"max_retry_backoff"`
			MaxIdleTime     time.Duration `json:"max_idle_time"`
			PoolTimeout     time.Duration `json:"pool_timeout"`
			ReadTimeout     time.Duration `json:"read_timeout"`
			WriteTimeout    time.Duration `json:"write_timeout"`
		})

		for key, redisConf := range conf.redisConfs {
			printer.RedisConfigs[key] = struct {
				ContextName     string        `json:"context_name"`
				Endpoint        string        `json:"endpoint"`
				Password        string        `json:"password"`
				DB              int           `json:"db"`
				PoolSize        int           `json:"pool_size"`
				MinIdleConns    int           `json:"min_idle_conns"`
				MaxRetries      int           `json:"max_retries"`
				MinRetryBackoff time.Duration `json:"min_retry_backoff"`
				MaxRetryBackoff time.Duration `json:"max_retry_backoff"`
				MaxIdleTime     time.Duration `json:"max_idle_time"`
				PoolTimeout     time.Duration `json:"pool_timeout"`
				ReadTimeout     time.Duration `json:"read_timeout"`
				WriteTimeout    time.Duration `json:"write_timeout"`
			}{
				ContextName:     redisConf.GetContextName(),
				Endpoint:        redisConf.GetEndpoint(),
				Password:        redisConf.GetPassword(),
				DB:              redisConf.GetDB(),
				PoolSize:        redisConf.GetPoolSize(),
				MinIdleConns:    redisConf.GetMinIdleConns(),
				MaxRetries:      redisConf.GetMaxRetries(),
				MinRetryBackoff: redisConf.GetMinRetryBackoff(),
				MaxRetryBackoff: redisConf.GetMaxRetryBackoff(),
				MaxIdleTime:     redisConf.GetMaxIdleTime(),
				PoolTimeout:     redisConf.GetPoolTimeout(),
				ReadTimeout:     redisConf.GetReadTimeout(),
				WriteTimeout:    redisConf.GetWriteTimeout(),
			}
		}
	}

	return convutil.Obj2Map(printer)
}

func (conf *Config) dbConfigLoader(fileLocation string) error {
	conf.dbConfsMutex.Lock()
	defer conf.dbConfsMutex.Unlock()
	dbConfs := DBConfigs{}
	if err := ReadConfigFile(fileLocation, &dbConfs); err != nil {
		return err
	}
	for idx, _ := range dbConfs.Databases {
		conf.dbConfs[dbConfs.Databases[idx].GetContextName()] = dbConfs.Databases[idx]
	}
	return nil
}

func (conf *Config) GetDBContextNames() []string {
	conf.dbConfsMutex.RLock()
	defer conf.dbConfsMutex.RUnlock()
	contextNames := make([]string, 0)
	for idx, _ := range conf.dbConfs {
		contextNames = append(contextNames, conf.dbConfs[idx].GetContextName())
	}
	return contextNames
}

func (conf *Config) GetDBConfig(contextName string) (IDBConfig, bool) {
	conf.dbConfsMutex.RLock()
	defer conf.dbConfsMutex.RUnlock()
	cfg, ok := conf.dbConfs[contextName]
	if !ok {
		return nil, false
	}
	return cfg, ok
}

func (conf *Config) GetDBConfigs() ([]IDBConfig, bool) {
	conf.dbConfsMutex.RLock()
	defer conf.dbConfsMutex.RUnlock()
	if len(conf.dbConfs) == 0 {
		return nil, false
	}
	configs := make([]IDBConfig, 0)
	for idx, _ := range conf.dbConfs {
		configs = append(configs, conf.dbConfs[idx])
	}
	return configs, true
}

func (conf *Config) logConfigLoader(fileLocation string) error {
	conf.logCfgMutex.Lock()
	defer conf.logCfgMutex.Unlock()
	logCfg := &LogConfig{}
	if err := ReadConfigFile(fileLocation, logCfg); err != nil {
		return err
	}
	conf.logConf = logCfg
	return nil
}

func (conf *Config) GetLogConfig() (ILogConfig, bool) {
	conf.logCfgMutex.RLock()
	defer conf.logCfgMutex.RUnlock()
	if conf.logConf == nil {
		return nil, false
	}
	return conf.logConf, true
}

func (conf *Config) apiConfigLoader(fileLocation string) error {
	conf.logCfgMutex.Lock()
	defer conf.logCfgMutex.Unlock()
	apiCfg := &APIConfig{}
	if err := ReadConfigFile(fileLocation, apiCfg); err != nil {
		return err
	}
	conf.apiConf = apiCfg
	return nil
}

func (conf *Config) GetAPIConfig() (IAPIConfig, bool) {
	conf.apiCfgMutex.RLock()
	defer conf.apiCfgMutex.RUnlock()
	if conf.apiConf == nil {
		return nil, false
	}
	return conf.apiConf, true
}

func (conf *Config) GetRedisContextNames() []string {
	conf.redisCfgMutex.RLock()
	defer conf.redisCfgMutex.RUnlock()
	contextNames := make([]string, 0)
	for idx, _ := range conf.redisConfs {
		contextNames = append(contextNames, conf.redisConfs[idx].GetContextName())
	}
	return contextNames
}

func (conf *Config) GetRedisConfig(contextName string) (IRedisConfig, bool) {
	conf.redisCfgMutex.RLock()
	defer conf.redisCfgMutex.RUnlock()
	cfg, ok := conf.redisConfs[contextName]
	if !ok {
		return nil, false
	}
	return cfg, ok
}

func (conf *Config) GetRedisConfigs() ([]IRedisConfig, bool) {
	conf.redisCfgMutex.RLock()
	defer conf.redisCfgMutex.RUnlock()
	if len(conf.redisConfs) == 0 {
		return nil, false
	}
	configs := make([]IRedisConfig, 0)
	for idx, _ := range conf.redisConfs {
		configs = append(configs, conf.redisConfs[idx])
	}
	return configs, true
}

func (conf *Config) redisConfigLoader(fileLocation string) error {
	conf.redisCfgMutex.Lock()
	defer conf.redisCfgMutex.Unlock()
	redisCfgs := &RedisConfigs{}
	if err := ReadConfigFile(fileLocation, redisCfgs); err != nil {
		return err
	}
	for idx, _ := range redisCfgs.Caches {
		conf.redisConfs[redisCfgs.Caches[idx].GetContextName()] = redisCfgs.Caches[idx]
	}
	return nil
}
