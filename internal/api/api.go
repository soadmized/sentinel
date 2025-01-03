package api

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/soadmized/sentinel/pkg/dataset"
)

type Service interface {
	SaveValues(ctx context.Context, dataset dataset.Dataset) error
	LastValues(ctx context.Context, sensorID string) *dataset.Dataset
	SensorStatuses(ctx context.Context) map[string]bool
	SensorIDs() []string
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
	a.Server.POST("/last_values", a.lastValues)
	a.Server.GET("/status", a.status)
	a.Server.POST("/sensor_ids", a.sensorIDs)
}
