package ihttp

import (
	"fmt"
	"github.com/gitkeng/ihttp/util/dateutil"
	"github.com/gitkeng/ihttp/util/stringutil"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io"
	"time"
)

// HTTPContext implement IContext it is context for HTTP
type HTTPContext struct {
	ms  *Microservice
	ctx echo.Context
}

// NewHTTPContext is the constructor function for HTTPContext
func NewHTTPContext(ms *Microservice, ctx echo.Context) *HTTPContext {
	if ms == nil {
		return nil
	}
	return &HTTPContext{
		ms:  ms,
		ctx: ctx,
	}
}

// WebContext return the echo.Context
func (ctx *HTTPContext) WebContext() echo.Context {
	return ctx.ctx
}

// Param return parameter by name
func (ctx *HTTPContext) Param(name string) string {
	if ctx.ctx != nil {
		return ctx.ctx.Param(name)
	}
	return ""
}

// QueryParam return query param
func (ctx *HTTPContext) QueryParam(name string) string {
	if ctx.ctx != nil {
		return ctx.ctx.QueryParam(name)
	}
	return ""
}

// ReadRequest read the request body and return it as string
func (ctx *HTTPContext) ReadRequest() string {
	if ctx.ctx != nil {
		body, err := io.ReadAll(ctx.ctx.Request().Body)
		if err != nil {
			return ""
		}
		return string(body)
	}
	return ""
}

// Bind bind the request body to the given struct
func (ctx *HTTPContext) Bind(request any) error {
	if err := ctx.ctx.Bind(request); err != nil {
		return err
	}
	return ctx.ctx.Validate(request)
}

// ReadRequests return nil in HTTP WebContext
func (ctx *HTTPContext) ReadRequests() []string {
	return nil
}

// Response return response to client
func (ctx *HTTPContext) Response(logLevel LogLevel, logTag string, httpStatus int, code, message string, err error, fields ...Field) error {
	var dataFields map[string]any
	if len(fields) > 0 {
		dataFields = make(map[string]any)
		for _, field := range fields {
			dataFields[field.Key] = field.Value
		}
	}

	var requestId string

	if len(dataFields) > 0 {
		if id, found := dataFields[RequestIdField]; found {
			if stringutil.IsNotEmptyString(id.(string)) {
				requestId = id.(string)
			}
		}
	}

	if ctx.WebContext() != nil && stringutil.IsEmptyString(requestId) {
		req := ctx.WebContext().Request()
		res := ctx.WebContext().Response()
		id := req.Header.Get(echo.HeaderXRequestID)
		if stringutil.IsEmptyString(id) {
			id = res.Header().Get(echo.HeaderXRequestID)
		}
		if stringutil.IsNotEmptyString(id) {
			requestId = id
		}
	}

	var response Response
	if err != nil {
		errors := ToErrors(err, code, "", nil)
		response = Response{
			RequestId:  requestId,
			StatusCode: httpStatus,
			Code:       code,
			Message:    message,
			Data:       dataFields,
			Error:      errors,
		}

		ctx.Logger().Log(logLevel, fmt.Sprintf("[%s] %s: %s", logTag, code, message), zap.Any("response", response.ToMap()))
		return ctx.WebContext().JSON(httpStatus, response)
	} else {
		response = Response{
			RequestId:  requestId,
			StatusCode: httpStatus,
			Code:       code,
			Message:    message,
			Data:       dataFields,
		}
		ctx.Logger().Log(logLevel, fmt.Sprintf("[%s] %s: %s", logTag, code, message), zap.Any("response", response.ToMap()))
		return ctx.WebContext().JSON(httpStatus, response)

	}

}

// Now return current time
func (ctx *HTTPContext) Now() time.Time {
	return time.Now()
}

// Epoch return current epoch time
func (ctx *HTTPContext) Epoch() int64 {
	return dateutil.GetCurrentEpochTime()
}

func (ctx *HTTPContext) DB(dbContextName string) (IDBStore, bool) {
	if ctx.ms != nil && len(ctx.ms.dbStores) > 0 {
		dbstore, found := ctx.ms.dbStores[dbContextName]
		return dbstore, found
	}
	return nil, false

}

func (ctx *HTTPContext) Cache(cacheContextName string) (IRedisCache, bool) {
	if ctx.ms != nil && len(ctx.ms.redisCaches) > 0 {
		redis, found := ctx.ms.redisCaches[cacheContextName]
		return redis, found
	}
	return nil, false

}

// APIConfig return APIConfig
func (ctx *HTTPContext) APIConfig() (IAPIConfig, bool) {
	if ctx.ms != nil && ctx.ms.apiConfig != nil {
		cfg := ctx.ms.apiConfig
		return cfg, cfg != nil
	}
	return nil, false
}

// LogConfig return LogConfig
func (ctx *HTTPContext) LogConfig() (ILogConfig, bool) {
	if ctx.ms != nil && ctx.ms.logConfig != nil {
		cfg := ctx.ms.logConfig
		return cfg, cfg != nil
	}
	return nil, false
}

// DBConfig return DBConfig
func (ctx *HTTPContext) DBConfig(contextName string) (IDBConfig, bool) {
	if ctx.ms != nil && len(ctx.ms.dbConfigs) > 0 {
		conf, found := ctx.ms.dbConfigs[contextName]
		return conf, found
	}
	return nil, false
}

// RedisConfig return RedisConfig
func (ctx *HTTPContext) RedisConfig(contextName string) (IRedisConfig, bool) {
	if ctx.ms != nil && len(ctx.ms.redisConfigs) > 0 {
		conf, found := ctx.ms.redisConfigs[contextName]
		return conf, found
	}
	return nil, false
}

// Requester return Requester
func (ctx *HTTPContext) Requester(baseURL string, timeout time.Duration, certFiles ...string) (IRequester, error) {
	return NewRequester(ctx.ms, baseURL, timeout, certFiles...)
}
