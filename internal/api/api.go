package api

import (
	"fmt"
	"sentinel/internal/dataset"

	"github.com/labstack/echo/v4"
)

type Service interface {
	SaveValues(dataset dataset.Dataset) error
	LastValues(sensorID string) *dataset.Dataset
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
