// /server.go
// Server manages the HTTP server and request handling

package stardust

import (
	"fmt"
	"net/http"
)

type Server struct {
	app        *App         // reference to the main application instance
	httpServer *http.Server // underlying HTTP server instance
}

// newServer creates a new Server instance with configured HTTP server
func newServer(app *App) *Server {
	return &Server{
		app: app,
		httpServer: &http.Server{
			Addr: fmt.Sprintf(":%s", app.config.Port),
		},
	}
}

// start initializes and starts the HTTP server
func (s *Server) start() error {
	// Set up the main handler
	s.httpServer.Handler = s.buildHandler()

	fmt.Printf("Server starting on http://localhost:%s\n", s.app.config.Port)
	return s.httpServer.ListenAndServe()
}

// buildHandler constructs the main HTTP handler with middleware and routes
func (s *Server) buildHandler() http.Handler {
	// Create the main mux
	mux := http.NewServeMux()

	// Add static file handler if static directory is configured
	if s.app.config.StaticDir != "" {
		staticHandler := http.FileServer(http.Dir(s.app.config.StaticDir))
		mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	}

	// Add the main application handler
	mux.HandleFunc("/", s.handleRequest)

	// Apply global middleware
	var handler http.Handler = mux
	for i := len(s.app.middleware) - 1; i >= 0; i-- {
		handler = s.wrapMiddleware(handler, s.app.middleware[i])
	}

	return handler
}

// handleRequest processes incoming HTTP requests
func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	// Create context for the request
	ctx := &Context{
		Request:  r,
		Response: w,
		engine:   s.app.template,
	}

	// Find the handler for this route
	handler := s.app.router.match(r.Method, r.URL.Path)
	if handler == nil {
		http.NotFound(w, r)
		return
	}

	// Execute the handler
	if err := handler(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// wrapMiddleware applies middleware to an http.Handler
func (s *Server) wrapMiddleware(handler http.Handler, middleware MiddlewareFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{
			Request:  r,
			Response: w,
			engine:   s.app.template,
		}

		nextHandler := func(c *Context) error {
			handler.ServeHTTP(c.Response, c.Request)
			return nil
		}

		if err := middleware(nextHandler)(ctx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
