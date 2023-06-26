package ihttp

import "github.com/labstack/echo/v4"

type HealthCheckFunc func(ms *Microservice) error

func (ms *Microservice) registerHealthCheck() {
	ms.echo.GET(ms.healthCheckEndpoint, func(c echo.Context) error {
		for _, health := range ms.healthCheckFuncs {
			if err := health(ms); err != nil {
				ms.responseUnHealthy(c.Response(), err.Error())
				return nil
			}
		}
		ms.responseHealthy(c.Response())
		return nil
	})
}
