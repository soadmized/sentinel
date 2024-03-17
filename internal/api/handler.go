package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

//nolint:wrapcheck
func (a *API) saveValues(ctx echo.Context) error {
	req := new(saveValuesReq)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "convert request body to model")
	}

	stamp := time.Now()
	set := req.toModel(stamp)

	if err := a.Service.SaveValues(set); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "save values")
	}

	return ctx.JSON(http.StatusOK, "saved")
}

//nolint:wrapcheck
func (a *API) lastValues(ctx echo.Context) error {
	req := new(lastValuesReq)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "convert request body to model")
	}

	values := a.Service.LastValues(req.ID)

	return ctx.JSON(http.StatusOK, values)
}

//nolint:wrapcheck
func (a *API) status(ctx echo.Context) error {
	statuses := a.Service.SensorStatuses()

	return ctx.Render(http.StatusOK, "status", statuses)
}

//nolint:wrapcheck
func (a *API) sensorIDs(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, a.Service.SensorIDs())
}
