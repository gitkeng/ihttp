package ihttp

import (
	"fmt"
	"github.com/gitkeng/ihttp/log"
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/dbutil"
	"github.com/gitkeng/ihttp/util/id"
	"github.com/gitkeng/ihttp/util/stringutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// ServiceHandleFunc is the handler for each Microservice
type ServiceHandleFunc func(ctx IContext) error
type CleanupFunc func(IMicroservice) error

// IMicroservice is interface for centralized service management
type IMicroservice interface {
	Start() error
	Stop()
	Cleanup() error
	Log(LogLevel, string, ...any)
	Logger() IContextLogger
	DefaultLogger() log.ILogger
	GetEngine() *echo.Echo
	GetHttpPort() int
	GetHttpsPort() int

	//HTTP Services

	// GET is the function to register GET method
	GET(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
	// POST is the function to register POST method
	POST(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
	// PUT is the function to register PUT method
	PUT(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
	// PATCH is the function to register PATCH method
	PATCH(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
	// DELETE is the function to register DELETE method
	DELETE(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)

	//DB return the DBStore
	DB(dbContextName string) (IDBStore, bool)
	//Cache return the RedisCache
	Cache(cacheContextName string) (IRedisCache, bool)
}

// Microservice is the centralized service management
type Microservice struct {
	echo *echo.Echo
	port int

	httpsOnly bool
	//ssl setting
	sslEnable   bool
	sslPort     int
	sslCertFile string
	sslKeyFile  string

	healthCheckEndpoint string
	healthCheckFuncs    []HealthCheckFunc
	middlewares         []echo.MiddlewareFunc

	exitChannel       chan bool
	logger            log.ILogger
	logFileEnable     bool
	logFileLocation   string
	logFileMaxSize    int
	logFileMaxBackups int
	logFileMaxAge     int
	logLevel          zapcore.Level

	dbStores    map[string]IDBStore
	redisCaches map[string]IRedisCache

	logConfig    ILogConfig
	apiConfig    IAPIConfig
	redisConfigs map[string]IRedisConfig
	dbConfigs    map[string]IDBConfig

	cleanupFuncs []CleanupFunc
}

// String returns a Microservice as a string
func (ms *Microservice) String() string {
	printObj := struct {
		HealthCheckEndpoint   string         `json:"health_check_endpoint"`
		TotalHealthCheckFuncs int            `json:"total_health_check_funcs"`
		TotalMiddlewares      int            `json:"total_middlewares"`
		Port                  int            `json:"port"`
		HttpsOnly             bool           `json:"https_only"`
		SSLEnable             bool           `json:"ssl_enable"`
		SSLPort               int            `json:"ssl_port"`
		SSLCertFile           string         `json:"ssl_cert_file"`
		SSLKeyFile            string         `json:"ssl_key_file"`
		LogFileEnable         bool           `json:"log_file_enable"`
		LogFileLocation       string         `json:"log_file_location"`
		LogFileMaxSize        int            `json:"log_file_max_size"`
		LogFileMaxBackups     int            `json:"log_file_max_backups"`
		LogFileMaxAge         int            `json:"log_file_max_age"`
		LogLevel              string         `json:"log_level"`
		RedisConfigs          map[string]any `json:"redis_configs"`
		DBConfigs             map[string]any `json:"db_configs"`
	}{
		HealthCheckEndpoint:   ms.healthCheckEndpoint,
		TotalHealthCheckFuncs: len(ms.healthCheckFuncs),
		TotalMiddlewares:      len(ms.middlewares),
		Port:                  ms.port,
		HttpsOnly:             ms.httpsOnly,
		SSLEnable:             ms.sslEnable,
		SSLPort:               ms.sslPort,
		SSLCertFile:           ms.sslCertFile,
		SSLKeyFile:            ms.sslKeyFile,
		LogFileEnable:         ms.logFileEnable,
		LogFileMaxSize:        ms.logFileMaxSize,
		LogFileMaxBackups:     ms.logFileMaxBackups,
		LogFileLocation:       ms.logFileLocation,
		LogFileMaxAge:         ms.logFileMaxAge,
		LogLevel:              ms.logLevel.String(),
		RedisConfigs:          convutil.Obj2Map(ms.redisConfigs),
		DBConfigs:             convutil.Obj2Map(ms.dbConfigs),
	}
	return stringutil.Json(printObj)
}

// ToMap returns a Microservice as a map
func (ms *Microservice) ToMap() map[string]any {
	printObj := struct {
		HealthCheckEndpoint   string         `json:"health_check_endpoint"`
		TotalHealthCheckFuncs int            `json:"total_health_check_funcs"`
		TotalMiddlewares      int            `json:"total_middlewares"`
		Port                  int            `json:"port"`
		HttpsOnly             bool           `json:"https_only"`
		SSLEnable             bool           `json:"ssl_enable"`
		SSLPort               int            `json:"ssl_port"`
		SSLCertFile           string         `json:"ssl_cert_file"`
		SSLKeyFile            string         `json:"ssl_key_file"`
		LogFileEnable         bool           `json:"log_file_enable"`
		LogFileLocation       string         `json:"log_file_location"`
		LogFileMaxSize        int            `json:"log_file_max_size"`
		LogFileMaxBackups     int            `json:"log_file_max_backups"`
		LogFileMaxAge         int            `json:"log_file_max_age"`
		LogLevel              string         `json:"log_level"`
		RedisConfigs          map[string]any `json:"redis_configs"`
		DBConfigs             map[string]any `json:"db_configs"`
	}{
		HealthCheckEndpoint:   ms.healthCheckEndpoint,
		TotalHealthCheckFuncs: len(ms.healthCheckFuncs),
		TotalMiddlewares:      len(ms.middlewares),
		Port:                  ms.port,
		HttpsOnly:             ms.httpsOnly,
		SSLEnable:             ms.sslEnable,
		SSLPort:               ms.sslPort,
		SSLCertFile:           ms.sslCertFile,
		SSLKeyFile:            ms.sslKeyFile,
		LogFileEnable:         ms.logFileEnable,
		LogFileMaxSize:        ms.logFileMaxSize,
		LogFileMaxBackups:     ms.logFileMaxBackups,
		LogFileLocation:       ms.logFileLocation,
		LogFileMaxAge:         ms.logFileMaxAge,
		LogLevel:              ms.logLevel.String(),
		RedisConfigs:          convutil.Obj2Map(ms.redisConfigs),
		DBConfigs:             convutil.Obj2Map(ms.dbConfigs),
	}
	return convutil.Obj2Map(printObj)
}

// New is the constructor function of Microservice
func New(options ...Option) (*Microservice, error) {
	ms := &Microservice{
		port:                DefaultPort,
		httpsOnly:           false,
		sslEnable:           false,
		sslPort:             DefaultSSLPort,
		healthCheckEndpoint: DefaultHealthCheckEndpoint,
		healthCheckFuncs:    make([]HealthCheckFunc, 0),
		middlewares:         make([]echo.MiddlewareFunc, 0),
		logFileMaxSize:      DefaultLogFileMaxSize,
		logFileMaxBackups:   DefaultLogFileMaxBackups,
		logFileMaxAge:       DefaultLogFileMaxAge,
		logLevel:            zapcore.DebugLevel,
		dbStores:            make(map[string]IDBStore),
		redisCaches:         make(map[string]IRedisCache),
		redisConfigs:        make(map[string]IRedisConfig),
		dbConfigs:           make(map[string]IDBConfig),
		cleanupFuncs:        make([]CleanupFunc, 0),
	}

	for _, setter := range options {
		if setter != nil {
			err := setter(ms)
			if err != nil {
				return nil, err
			}
		}

	}

	// connect to other services
	if err := ms.connectDB(); err != nil {
		ms.Cleanup()
		return nil, err
	}
	if err := ms.connectRedisCache(); err != nil {
		ms.Cleanup()
		return nil, err
	}

	ms.echo = echo.New()
	ms.echo.Validator = &Validator{}
	ms.echo.HideBanner = true
	// setup health check
	ms.registerHealthCheck()

	//initial logger
	if ms.logFileEnable {
		logger, err := log.New(
			log.WithLevel(ms.logLevel),
			log.WithFileLocation(ms.logFileLocation),
			log.WithFileMaxSize(ms.logFileMaxSize),
			log.WithFileMaxBackups(ms.logFileMaxBackups),
			log.WithFileMaxAge(ms.logFileMaxAge),
		)
		if err != nil {
			panic(err)
		}
		ms.logger = logger
	} else {
		logger, err := log.New(
			log.WithLevel(ms.logLevel),
			log.WithFileMaxSize(ms.logFileMaxSize),
			log.WithFileMaxBackups(ms.logFileMaxBackups),
			log.WithFileMaxAge(ms.logFileMaxAge),
		)
		if err != nil {
			panic(err)
		}
		ms.logger = logger

	}

	ms.echo.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return id.UUID()
		},
	}))

	if len(ms.middlewares) > 0 {
		ms.echo.Use(ms.middlewares...)
	}

	return ms, nil
}

// Start start all registered services
func (ms *Microservice) Start() error {
	routeCount := len(ms.echo.Routes())
	var exitHTTP chan bool
	if routeCount > 0 {
		exitHTTP = make(chan bool, 1)
		go func() {
			ms.startServer(exitHTTP)
		}()
	}

	// There are 2 ways to exit from Microservices
	// 1. The SigTerm can be sent from outside program such as from k8s
	// 2. Send true to ms.exitChannel
	osQuit := make(chan os.Signal, 1)
	ms.exitChannel = make(chan bool, 1)
	signal.Notify(osQuit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	exit := false
	for {
		if exit {
			break
		}
		select {
		case <-osQuit:
			// Exit from HTTP as well
			if exitHTTP != nil {
				exitHTTP <- true
			}
			exit = true
		case <-ms.exitChannel:
			// Exit from HTTP as well
			if exitHTTP != nil {
				exitHTTP <- true
			}
			exit = true
		}
	}

	ms.Cleanup()
	return nil
}

// Stop stop the services
func (ms *Microservice) Stop() {
	if ms.exitChannel == nil {
		return
	}
	ms.exitChannel <- true
}

// Cleanup clean resources up from every registered services before exit
func (ms *Microservice) Cleanup() error {
	if len(ms.cleanupFuncs) > 0 {
		for _, cleanupFunc := range ms.cleanupFuncs {
			if err := cleanupFunc(ms); err != nil {
				ms.Log(WarnLevel, fmt.Sprintf("call cleanup function fail  %s", err.Error()))
			}
		}
	}

	for key, dbStore := range ms.dbStores {
		err := dbStore.Close()
		if err != nil {
			ms.Log(WarnLevel, fmt.Sprintf("closing db store fail %s", err.Error()))
		}
		delete(ms.dbStores, key)
	}

	for key, redis := range ms.redisCaches {
		err := redis.Close()
		if err != nil {
			ms.Log(WarnLevel, fmt.Sprintf("closing redis fail %s", err.Error()))
		}
		delete(ms.redisCaches, key)
	}

	return nil
}

func (ms *Microservice) Logger() IContextLogger {
	return ms
}

func (ms *Microservice) DefaultLogger() log.ILogger {
	return ms.logger
}

func (ms *Microservice) connectDB() error {
	//setup dbStores from dbConfigs
	for key, _ := range ms.dbConfigs {
		dbStore, err := NewDBStore(ms.dbConfigs[key])
		if err != nil {
			return err
		}
		ms.dbStores[dbStore.Config().GetContextName()] = dbStore

		initScripts := ms.dbConfigs[key].GetInitialScripts()
		if len(initScripts) > 0 {
			for _, initScript := range initScripts {
				if sqlCmd, err := dbutil.SQLLoader(initScript); err != nil {
					return err
				} else {
					if _, err := dbStore.Conn().Exec(sqlCmd); err != nil {
						return err
					}
				}
			}
		}

	}
	return nil
}

func (ms *Microservice) connectRedisCache() error {
	//setup redisCaches from redisConfigs
	for key, _ := range ms.redisConfigs {
		cache := NewRedisCache(ms.redisConfigs[key])
		err := cache.Open()
		if err != nil {
			return err
		}
		ms.redisCaches[ms.redisConfigs[key].GetContextName()] = cache
	}
	return nil
}

func (ms *Microservice) responseHealthy(resp *echo.Response) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("ok"))
}

func (ms *Microservice) responseUnHealthy(resp *echo.Response, reason string) {
	errMsg := "Healthcheck failed because of " + reason
	resp.WriteHeader(http.StatusInternalServerError)
	resp.Write([]byte(errMsg))
}

func (ms *Microservice) DB(dbContextName string) (IDBStore, bool) {
	if len(ms.dbStores) > 0 {
		dbstore, found := ms.dbStores[dbContextName]
		return dbstore, found
	}
	return nil, false

}

func (ms *Microservice) Cache(cacheContextName string) (IRedisCache, bool) {
	if len(ms.redisCaches) > 0 {
		redis, found := ms.redisCaches[cacheContextName]
		return redis, found
	}
	return nil, false

}

func (ms *Microservice) GetEngine() *echo.Echo {
	return ms.echo
}

func (ms *Microservice) GetHttpPort() int {
	if ms.apiConfig.IsHttpsOnly() {
		return -1
	}
	return ms.apiConfig.GetPort()
}

func (ms *Microservice) GetHttpsPort() int {
	if ms.apiConfig.IsSSLEnable() {
		return ms.apiConfig.GetSSLPort()
	} else {
		return -1
	}
}
