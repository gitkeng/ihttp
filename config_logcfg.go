package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/stringutil"
	"strings"
)

type ILogConfig interface {
	IConfig
	// GetFileLocation is the option for setting log file location
	GetFileLocation() string
	// GetMaxSize is the option for setting log file max size (megabytes)
	GetMaxSize() int
	// GetMaxBackups is the option for setting log file max backup
	GetMaxBackups() int
	// GetMaxAge is the option for setting log file max backup
	GetMaxAge() int
	// GetLogLevel is the option for setting log level
	GetLogLevel() LogLevel
}

type LogConfig struct {
	FileLocation string   `mapstructure:"log-file-location" json:"file_location"`
	MaxSize      int      `mapstructure:"log-max-size" json:"max_size"`
	MaxBackups   int      `mapstructure:"log-max-backups" json:"max_backups"`
	MaxAge       int      `mapstructure:"log-max-age" json:"max_age"`
	LogLevel     LogLevel `mapstructure:"log-level" json:"log_level"`
}

func (logCfg *LogConfig) Bind() error {

	logCfg.FileLocation = strings.TrimSpace(logCfg.FileLocation)
	if logCfg.MaxSize < 50 {
		logCfg.MaxSize = DefaultLogFileMaxSize
	}
	if logCfg.MaxBackups <= 1 {
		logCfg.MaxBackups = DefaultLogFileMaxBackups
	}
	if logCfg.MaxAge <= 1 {
		logCfg.MaxAge = DefaultLogFileMaxAge
	}

	level := strings.TrimSpace(strings.ToLower(string(logCfg.LogLevel)))
	switch level {
	case "debug", "info", "warn", "error", "panic", "dpanic", "fatal":
		logCfg.LogLevel = LogLevel(level)
	default:
		logCfg.LogLevel = DebugLevel
	}
	return nil

}

func (logCfg *LogConfig) Validate() error {
	return nil
}

func (logCfg *LogConfig) String() string {
	return stringutil.Json(*logCfg)
}

func (logCfg *LogConfig) ToMap() map[string]any {
	return convutil.Obj2Map(*logCfg)
}

func (logCfg *LogConfig) GetFileLocation() string {
	return logCfg.FileLocation
}

func (logCfg *LogConfig) GetMaxSize() int {
	return logCfg.MaxSize
}

func (logCfg *LogConfig) GetMaxBackups() int {
	return logCfg.MaxBackups
}

func (logCfg *LogConfig) GetMaxAge() int {
	return logCfg.MaxAge
}

func (logCfg *LogConfig) GetLogLevel() LogLevel {
	return logCfg.LogLevel
}
