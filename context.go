// context.go
package gust

import (
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	engine   Engine
}

func (c *Context) Render(template string, data interface{}) error {
	return c.engine.Render(c.Response, template, data)
}
