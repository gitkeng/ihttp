package ihttp

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// GET register service endpoint for HTTP GET
func (ms *Microservice) GET(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {
	ms.echo.GET(path, func(ctx echo.Context) error {
		return h(NewHTTPContext(ms, ctx))
	}, m...)
}

// POST register service endpoint for HTTP POST
func (ms *Microservice) POST(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {
	ms.echo.POST(path, func(ctx echo.Context) error {
		return h(NewHTTPContext(ms, ctx))
	}, m...)
}

// PUT register service endpoint for HTTP PUT
func (ms *Microservice) PUT(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {
	ms.echo.PUT(path, func(ctx echo.Context) error {
		return h(NewHTTPContext(ms, ctx))
	}, m...)
}

// PATCH register service endpoint for HTTP PATCH
func (ms *Microservice) PATCH(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {
	ms.echo.PATCH(path, func(ctx echo.Context) error {
		return h(NewHTTPContext(ms, ctx))
	}, m...)
}

// DELETE register service endpoint for HTTP DELETE
func (ms *Microservice) DELETE(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {
	ms.echo.DELETE(path, func(ctx echo.Context) error {
		return h(NewHTTPContext(ms, ctx))
	}, m...)
}

func (ms *Microservice) startServer(exitChannel chan bool) error {
	// Caller can exit by sending value to exitChannel
	go func() {
		<-exitChannel
		ms.stopHTTP()
	}()

	//start http server
	go func() {
		if err := ms.startHTTP(); err != nil {
			if err != http.ErrServerClosed {
				ms.logger.Warnf("start http server error %s", err.Error())
			}
		}
	}()
	go func() {
		if err := ms.startHTTPS(); err != nil {
			if err != http.ErrServerClosed {
				ms.logger.Warnf("start https server error %s", err.Error())
			}
		}
	}()

	return nil
}

// startHTTP will start HTTP service, this function will block thread
func (ms *Microservice) startHTTP() error {
	// Caller can exit by sending value to exitChannel
	if !ms.httpsOnly {
		return ms.echo.Start(fmt.Sprintf(":%d", ms.port))
	}
	return nil
}

func (ms *Microservice) startHTTPS() error {
	if ms.sslEnable {
		return ms.echo.StartTLS(fmt.Sprintf(":%d", ms.sslPort), ms.sslCertFile, ms.sslKeyFile)
	}
	return nil
}

// stopHTTP will stop HTTP service
func (ms *Microservice) stopHTTP() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ms.echo.Shutdown(ctx)
}
