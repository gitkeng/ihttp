package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/stringutil"
	"strings"
)

const (
	POSTGRES = "postgres"
	MYSQL    = "mysql"
)

type IDBConfig interface {
	IConfig
	// GetContextName is the option for setting database context name
	GetContextName() string
	// GetProvider is the option for setting database provider
	GetProvider() string
	// GetURL is the option for setting database url
	GetURL() string
	// GetUser is the option for setting database user
	GetUser() string
	// GetPassword is the option for setting database password
	GetPassword() string
	// GetDatabaseName is the option for setting database name
	GetDatabaseName() string
	// GetRetryLimits is the option for setting database retry limits
	GetRetryLimits() int
	// GetInitialScripts is the option for setting database initial scripts
	GetInitialScripts() []string
	// GetConnectionMaxLifeTime is the maximum amount of time a dbConn may be reused.
	//
	// Expired connections may be closed lazily before reuse.
	//
	// If d <= 0, connections are not closed due to a dbConn's age.
	GetConnectionMaxLifeTime() int
	// GetMaxIdleConns is the maximum number of connections in the idle
	// dbConn pool.
	//
	// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
	// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
	//
	// If n <= 0, no idle connections are retained.
	//
	// The default max idle connections is currently 2. This may change in
	// a future release.
	GetMaxIdleConns() int
	// GetMaxOpenConns the maximum number of open connections to the database.
	//
	// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
	// MaxIdleConns, then MaxIdleConns will be reduced to match the new
	// MaxOpenConns limit.
	//
	// If n <= 0, then there is no limit on the number of open connections.
	// The default is 0 (unlimited).
	GetMaxOpenConns() int
}

type DBConfig struct {
	ContextName           string   `mapstructure:"context-name" json:"context_name"`
	Provider              string   `mapstructure:"provider" json:"provider"`
	URL                   string   `mapstructure:"url" json:"url"`
	User                  string   `mapstructure:"user" json:"user"`
	Password              string   `mapstructure:"password" json:"password"`
	DatabaseName          string   `mapstructure:"database-name" json:"database_name"`
	RetryLimits           int      `mapstructure:"retry-limits" json:"retry_limits"`
	ConnectionMaxLifeTime int      `mapstructure:"dbConn-max-life-time" json:"connection_max_life_time"`
	MaxIdleConns          int      `mapstructure:"max-idle-conns" json:"max_idle_conns"`
	MaxOpenConns          int      `mapstructure:"max-open-conns" json:"max_open_conns"`
	InitialScripts        []string `mapstructure:"initial-scripts" json:"initial_scripts"`
}

func (db *DBConfig) Bind() error {
	db.ContextName = strings.TrimSpace(db.ContextName)
	db.Provider = strings.ToLower(db.Provider)
	if db.RetryLimits <= 0 {
		db.RetryLimits = DefaultDBGetRetryLimits
	}
	if db.ConnectionMaxLifeTime <= 0 {
		db.ConnectionMaxLifeTime = DefaultDBConnectionMaxLifeTime
	}
	if db.MaxIdleConns <= 0 {
		db.MaxIdleConns = DefaultDBConnectionMaxIdleConns
	}
	if db.MaxOpenConns <= 0 {
		db.MaxOpenConns = DefaultDBConnectionMaxOpenConns
	}
	return nil
}

func (db *DBConfig) Validate() error {
	if stringutil.IsEmptyString(db.ContextName) {
		return ErrDBContextNameIsRequire
	}

	if db.Provider != POSTGRES && db.Provider != MYSQL {
		return ErrInvalidDBProvider(db.Provider)
	}

	if stringutil.IsEmptyString(db.URL) {
		return ErrDBURLRequire
	}

	if stringutil.IsEmptyString(db.User) {
		return ErrDBUserRequire
	}

	if stringutil.IsEmptyString(db.Password) {
		return ErrDBPasswordRequire
	}

	if stringutil.IsEmptyString(db.DatabaseName) {
		return ErrDBNameRequire
	}

	return nil
}

func (db *DBConfig) String() string {
	return stringutil.Json(*db)
}

func (db *DBConfig) ToMap() map[string]any {
	return convutil.Obj2Map(*db)
}

func (db *DBConfig) GetContextName() string {
	return db.ContextName
}

func (db *DBConfig) GetProvider() string {
	return db.Provider
}

func (db *DBConfig) GetURL() string {
	return db.URL
}

func (db *DBConfig) GetUser() string {
	return db.User
}

func (db *DBConfig) GetPassword() string {
	return db.Password
}

func (db *DBConfig) GetDatabaseName() string {
	return db.DatabaseName
}

func (db *DBConfig) GetRetryLimits() int {
	return db.RetryLimits
}

func (db *DBConfig) GetInitialScripts() []string {
	return db.InitialScripts
}

func (db *DBConfig) GetConnectionMaxLifeTime() int {
	return db.ConnectionMaxLifeTime
}

func (db *DBConfig) GetMaxIdleConns() int {
	return db.MaxIdleConns
}

func (db *DBConfig) GetMaxOpenConns() int {
	return db.MaxOpenConns
}
