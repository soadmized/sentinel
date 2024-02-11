package api

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data) //nolint:wrapcheck
}
