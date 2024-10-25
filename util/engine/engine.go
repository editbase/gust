// /util/engine/engine.go

package engine

import "net/http"

type Engine interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	RegisterRoutes()
}
