// /render.go
// Engine defines the interface for template rendering

package stardust

import "net/http"

// Render: renders templates with given data to http.ResponseWriter
type Engine interface {
	Render(w http.ResponseWriter, template string, data interface{}) error
}

// HandlerFunc defines the signature for HTTP request handlers
type HandlerFunc func(*Context) error
