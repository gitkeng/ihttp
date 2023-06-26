package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/fileutil"
	"github.com/gitkeng/ihttp/util/stringutil"
)

type IAPIConfig interface {
	IConfig
	// GetPort is the option for setting port
	GetPort() int
	// GetHealthCheckEndpoint is the option for setting health endpoint
	GetHealthCheckEndpoint() string
	// IsHttpsOnly is the option for setting https only
	IsHttpsOnly() bool
	//IsSSLEnable is the option for setting ssl enable
	IsSSLEnable() bool
	// GetSSLPort is the option for setting ssl port
	GetSSLPort() int
	// GetSSLCertFile is the option for setting ssl cert file
	GetSSLCertFile() string
	// GetSSLKeyFile is the option for setting ssl key file
	GetSSLKeyFile() string
}

type SSLType string

type APIConfig struct {
	// Port is the port for web api
	Port int `mapstructure:"port" json:"port"`
	// HealthPath is the path for health check
	HealthCheckEndpoint string `mapstructure:"health-check-endpoint" json:"health_check_endpoint"`
	// HttpsOnly is the option for setting https only
	HttpsOnly bool `mapstructure:"https-only" json:"https_only"`
	// SSLEnable is the option for setting ssl enable
	SSLEnable bool `mapstructure:"ssl-enable" json:"ssl_enable"`
	// SSLPort is the option for setting ssl port
	SSLPort int `mapstructure:"ssl-port" json:"ssl_port"`
	// SSLCertFile is the option for setting ssl cert file
	SSLCertFile string `mapstructure:"ssl-cert-file" json:"ssl_cert_file"`
	// SSLKeyFile is the option for setting ssl key file
	SSLKeyFile string `mapstructure:"ssl-key-file" json:"ssl_key_file"`
}

func (apiCfg *APIConfig) Bind() error {
	if apiCfg.Port <= 0 {
		apiCfg.Port = DefaultPort
	}
	if stringutil.IsEmptyString(apiCfg.HealthCheckEndpoint) {
		apiCfg.HealthCheckEndpoint = DefaultHealthCheckEndpoint
	}
	if apiCfg.SSLEnable {
		if apiCfg.SSLPort <= 0 {
			apiCfg.SSLPort = DefaultSSLPort
		}
	}
	return nil
}

func (apiCfg *APIConfig) Validate() error {
	if apiCfg.Port <= 0 {
		return ErrInvalidPort
	}
	if apiCfg.SSLEnable || apiCfg.HttpsOnly {
		//valid ssl port
		if apiCfg.SSLPort <= 0 {
			return ErrInvalidSSLPort
		}

		if stringutil.IsEmptyString(apiCfg.SSLCertFile) {
			return ErrSSLCertificateFileRequire
		}

		if found, _ := fileutil.IsFileExist(apiCfg.SSLCertFile); !found {
			return ErrSSLCertificateFileNotfound(apiCfg.SSLCertFile)
		}

		if stringutil.IsEmptyString(apiCfg.SSLKeyFile) {
			return ErrSSLKeyFileRequire
		}

		if found, _ := fileutil.IsFileExist(apiCfg.SSLKeyFile); !found {
			return ErrSSLKeyFileNotfound(apiCfg.SSLKeyFile)
		}

	}

	return nil
}

func (apiCfg *APIConfig) String() string {
	return stringutil.Json(*apiCfg)
}

func (apiCfg *APIConfig) ToMap() map[string]any {
	return convutil.Obj2Map(*apiCfg)
}

func (apiCfg *APIConfig) GetPort() int {
	return apiCfg.Port
}

func (apiCfg *APIConfig) GetHealthCheckEndpoint() string {
	return apiCfg.HealthCheckEndpoint
}

func (apiCfg *APIConfig) IsSSLEnable() bool {
	return apiCfg.SSLEnable
}

func (apiCfg *APIConfig) GetSSLPort() int {
	return apiCfg.SSLPort
}

func (apiCfg *APIConfig) GetSSLCertFile() string {
	return apiCfg.SSLCertFile
}

func (apiCfg *APIConfig) GetSSLKeyFile() string {
	return apiCfg.SSLKeyFile
}

func (apiCfg *APIConfig) IsHttpsOnly() bool {
	return apiCfg.HttpsOnly
}
