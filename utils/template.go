package utils

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.ExecuteTemplate(w, name, data)
}
