package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/stringutil"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	defaultAllowOrigins = []string{"*"}

	defaultAllowMethods = []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPut,
		http.MethodPatch,
		http.MethodPost,
		http.MethodDelete,
		http.MethodOptions,
	}

	defaultAllowHeader = []string{
		echo.HeaderOrigin,
		echo.HeaderContentType,
		echo.HeaderAccept,
		echo.HeaderAcceptEncoding,
		echo.HeaderAccessControlAllowHeaders,
		echo.HeaderAccessControlAllowMethods,
		echo.HeaderAccessControlAllowOrigin,
		echo.HeaderXRequestedWith,
		echo.HeaderAuthorization,
		echo.HeaderXCSRFToken,
		echo.HeaderXForwardedFor,
		echo.HeaderXRealIP,
	}

	defaultExposeHeaders = []string{}
	defaultMaxAge        = 0
)

type IAPICorsConfig interface {
	GetAllowOrigins() []string
	GetAllowMethods() []string
	GetAllowHeaders() []string
	IsAllowCredentials() bool
	GetExposeHeaders() []string
	GetMaxAge() int
}

type APICorsConfig struct {
	// AllowOrigin defines a list of origins that may access the resource.
	// Optional. Default value []string{"*"}.
	AllowOrigins []string `mapstructure:"allow-origins" json:"allow_origins"`
	// AllowMethods defines a list methods allowed when accessing the resource.
	// This is used in response to a preflight request.
	// Optional. Default value GET, HEAD, PUT, PATCH, POST, DELETE, OPTIONS.
	AllowMethods []string `mapstructure:"allow-methods" json:"allow_methods"`
	// AllowHeaders defines a list of request headers that can be used when
	// making the actual request. This is in response to a preflight request.
	// Optional. Default value Origin, Content-Type, Accept, Accept-Encoding, Access-Control-Allow-Headers,
	// Access-Control-Allow-Methods, Access-Control-Allow-Origin,
	// X-Requested-With, Authorization, X-CSRF-Token, X-Forwarded-For, X-Real-IP.
	AllowHeaders []string `mapstructure:"allow-headers" json:"allow_headers"`
	// AllowCredentials indicates whether or not the response to the request
	// can be exposed when the credentials flag is true. When used as part of
	// a response to a preflight request, this indicates whether or not the
	// actual request can be made using credentials.
	// Optional. Default value false.
	AllowCredentials bool `mapstructure:"allow-credentials" json:"allow_credentials"`
	// ExposeHeaders defines a whitelist headers that clients are allowed to
	// access.
	// Optional. Default value []string{}.
	ExposeHeaders []string `mapstructure:"expose-headers" json:"expose_headers"`
	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached.
	// Optional. Default value 0
	MaxAge int `mapstructure:"max-age" json:"max_age"`
}

func (corsCfg *APICorsConfig) Bind() error {
	//read cors config and check
	if len(corsCfg.AllowOrigins) <= 0 {
		corsCfg.AllowOrigins = defaultAllowOrigins
	}

	if len(corsCfg.AllowMethods) <= 0 {
		corsCfg.AllowMethods = defaultAllowMethods
	}

	if len(corsCfg.AllowHeaders) <= 0 {
		corsCfg.AllowHeaders = defaultAllowHeader
	}

	if len(corsCfg.ExposeHeaders) <= 0 {
		corsCfg.ExposeHeaders = defaultExposeHeaders
	}

	if corsCfg.MaxAge <= 0 {
		corsCfg.MaxAge = defaultMaxAge
	}
	return nil
}

func (corsCfg *APICorsConfig) Validate() error {
	return nil
}

func (corsCfg *APICorsConfig) String() string {
	return stringutil.Json(*corsCfg)
}

func (corsCfg *APICorsConfig) ToMap() map[string]any {
	return convutil.Obj2Map(*corsCfg)
}

func (corsCfg *APICorsConfig) GetAllowOrigins() []string {
	return corsCfg.AllowOrigins
}

func (corsCfg *APICorsConfig) GetAllowMethods() []string {
	return corsCfg.AllowMethods
}

func (corsCfg *APICorsConfig) GetAllowHeaders() []string {
	return corsCfg.AllowHeaders
}

func (corsCfg *APICorsConfig) IsAllowCredentials() bool {
	return corsCfg.AllowCredentials
}

func (corsCfg *APICorsConfig) GetExposeHeaders() []string {
	return corsCfg.ExposeHeaders
}

func (corsCfg *APICorsConfig) GetMaxAge() int {
	return corsCfg.MaxAge
}
