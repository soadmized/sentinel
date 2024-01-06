package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func saveValues(ctx echo.Context) error {
	return nil
}

func getLastValues(ctx echo.Context) error {
	return nil
}

func helloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "HELLO THERE")
}
