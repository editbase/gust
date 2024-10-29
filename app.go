// /app.go
// App represents the main application structure containing core components

package stardust

// App represents the main application structure containing core components
type App struct {
	router     *Router          // Handles HTTP routing
	config     *Config          // Stores application configuration
	middleware []MiddlewareFunc // Slice of middleware functions
	template   *Template        // Manages HTML template rendering
}

// Create and initialize a new App instance with default configuration
// Returns a pointer to the newly created App
func New() *App {
	app := &App{
		router:   newRouter(),
		config:   defaultConfig(),
		template: NewTemplate("./templates"),
	}
	return app
}

// WithTemplateDir sets custom template directory and reinitialize template engine
func (a *App) WithTemplateDir(dir string) *App {
	a.config.TemplateDir = dir
	a.template = NewTemplate(dir)
	return a
}

// WithPort sets the server port
func (a *App) WithPort(port string) *App {
	a.config.Port = port
	return a
}

// WithStaticDir sets the static files directory
func (a *App) WithStaticDir(dir string) *App {
	a.config.StaticDir = dir
	return a
}

// GET registers a new GET route with the specified path and handler
func (a *App) GET(path string, handler HandlerFunc) {
	a.router.addRoute("GET", path, handler)
}

// POST registers a new POST route with the specified path and handler
func (a *App) POST(path string, handler HandlerFunc) {
	a.router.addRoute("POST", path, handler)
}

// Use adds one or more middleware functions to the application
func (a *App) Use(middleware ...MiddlewareFunc) {
	a.middleware = append(a.middleware, middleware...)
}

// Run starts the HTTP server and loads templates
func (a *App) Run() error {
	if err := a.template.Load(); err != nil {
		// Returns an error if server fails to start or templates fail to load
		return err
	}

	server := newServer(a)
	return server.start()
}
