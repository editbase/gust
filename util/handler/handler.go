// /util/handler/handler.go

package handler

import (
	"net/http"

	"github.com/editbase/gust/util/engine"
)

// Handler represents the main HTTP handler for the application
type Handler struct {
	engine engine.Engine
}

// New creates a new Handler instance
func New(e engine.Engine) *Handler {
	return &Handler{
		engine: e,
	}
}

// ServeHTTP implements the http.Handler interface
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Add common middleware here
	h.addCommonHeaders(w)

	// Handle HTMX specific headers
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Push-Url", r.URL.String())
	}

	// Delegate to the engine
	h.engine.ServeHTTP(w, r)
}

// addCommonHeaders adds common response headers
func (h *Handler) addCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
}

// Middleware represents a function that wraps an http.Handler
type Middleware func(http.Handler) http.Handler

// Use applies middleware to the handler
func (h *Handler) Use(middleware ...Middleware) {
	for _, m := range middleware {
		h = &Handler{
			engine: wrapEngine(h.engine, m),
		}
	}
}

// wrapEngine wraps an engine with middleware
func wrapEngine(e engine.Engine, m Middleware) engine.Engine {
	return &wrappedEngine{
		engine:     e,
		middleware: m,
	}
}

// wrappedEngine represents an engine wrapped with middleware
type wrappedEngine struct {
	engine     engine.Engine
	middleware Middleware
}

func (w *wrappedEngine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	w.middleware(w.engine).ServeHTTP(resp, req)
}

func (w *wrappedEngine) RegisterRoutes() {
	w.engine.RegisterRoutes()
}
