// internal/handler/static.go
package handler

import (
	"net/http"
)

func Static(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}

func StaticWithPrefix(prefix, dir string) http.Handler {
	return http.StripPrefix(prefix, Static(dir))
}
