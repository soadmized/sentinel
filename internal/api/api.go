package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"sentinel/internal/dataset"
)

type Service interface {
	SaveValues(dataset dataset.Dataset) error
	LastValues(sensorID string) *dataset.Dataset
	SensorStatuses() map[string]bool
}

type API struct {
	Service Service
	Server  *echo.Echo
}

func (a *API) Start(port int) error {
	addr := fmt.Sprintf(":%d", port)

	err := a.Server.Start(addr)
	if err != nil {
		return errors.Wrap(err, "start echo server is failed")
	}

	return nil
}

func (a *API) Route() {
	a.Server.POST("/save_values", a.saveValues)
	a.Server.POST("/last_values", a.getLastValues)
	a.Server.GET("/status", a.status)
}
