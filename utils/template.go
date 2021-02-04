package utils

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// Template implemented the Renderer interface from echo.
type Template struct {
	Template *template.Template
}

// Render implements Renderer interface.
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	c.Logger().Infof("inside render :%#v", data)
	return t.Template.ExecuteTemplate(w, name, data)
}
