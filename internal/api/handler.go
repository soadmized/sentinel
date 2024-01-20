package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *Api) saveValues(ctx echo.Context) error {
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

func (a *Api) getLastValues(ctx echo.Context) error {
	req := new(lastValuesReq)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "convert request body to model")
	}

	values := a.Service.LastValues(req.Id)

	return ctx.JSON(http.StatusOK, values)
}

func (a *Api) helloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "HELLO THERE")
}
