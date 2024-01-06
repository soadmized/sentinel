package api

import (
	"fmt"
	"sentry/internal/service"

	"github.com/labstack/echo/v4"
)

type Api struct {
	Service *service.Service
	Server  *echo.Echo
}

func (a *Api) Start(port int) error {
	addr := fmt.Sprintf(":%d", port)

	err := a.Server.Start(addr)
	if err != nil {
		return err
	}

	return nil
}

func (a *Api) Route() {
	a.Server.POST("/save_values", saveValues)
	a.Server.POST("/get_last_values", getLastValues)
	a.Server.GET("hello_world", helloWorld)
}
