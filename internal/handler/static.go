// /internal/handler/static.go
package handler

import (
	"net/http"
)

// Static creates a file server for serving static files
func Static(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}

// StaticWithPrefix creates a file server with URL path prefix stripping
func StaticWithPrefix(prefix, dir string) http.Handler {
	return http.StripPrefix(prefix, Static(dir))
}
