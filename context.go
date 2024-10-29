// /context.go
// Context holds request-specific data and provides helper methods

package stardust

import (
	"net/http"
)

// Context holds request-specific data and provides helper methods
type Context struct {
	Request  *http.Request       // Current HTTP request
	Response http.ResponseWriter // HTTP response writer
	engine   Engine              // HTML template engine
}

// Render executes template rendering with the given template name and data
func (c *Context) Render(template string, data interface{}) error {
	return c.engine.Render(c.Response, template, data)
}
