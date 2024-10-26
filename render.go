// render.go
package gust

import "net/http"

type Engine interface {
	Render(w http.ResponseWriter, template string, data interface{}) error
}

type HandlerFunc func(*Context) error
