package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"sentinel/internal/model"
)

type Service interface {
	SaveValues(dataset model.Dataset) error
	LastValues() *model.Dataset
}

type Api struct {
	Service Service
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
	a.Server.POST("/save_values", a.saveValues)
	a.Server.POST("/last_values", a.getLastValues)
	a.Server.GET("/hello_world", a.helloWorld)
}
