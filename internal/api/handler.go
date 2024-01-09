package api

import (
	"net/http"
	"sentinel/internal/model"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *Api) saveValues(ctx echo.Context) error {
	values := new(model.Dataset)
	values.UpdatedAt = time.Now()

	if err := ctx.Bind(values); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "convert request body to model")
	}

	if err := a.Service.SaveValues(*values); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "save values")
	}

	return ctx.JSON(http.StatusOK, "saved")
}

func (a *Api) getLastValues(ctx echo.Context) error {
	values := a.Service.LastValues()

	return ctx.JSON(http.StatusOK, values)
}

func (a *Api) helloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "HELLO THERE")
}
