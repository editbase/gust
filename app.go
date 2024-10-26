// app.go
package gust

type App struct {
	router     *Router
	config     *Config
	middleware []MiddlewareFunc
	template   *Template
}

// app.go
func New() *App {
	app := &App{
		router:   newRouter(),
		config:   defaultConfig(),
		template: NewTemplate("./templates"), // Initialize template engine
	}
	return app
}

func (a *App) WithTemplateDir(dir string) *App {
	a.config.TemplateDir = dir
	a.template = NewTemplate(dir) // Create new template engine with updated dir
	return a
}

func (a *App) WithPort(port string) *App {
	a.config.Port = port
	return a
}
func (a *App) WithStaticDir(dir string) *App {
	a.config.StaticDir = dir
	return a
}

func (a *App) GET(path string, handler HandlerFunc) {
	a.router.addRoute("GET", path, handler)
}

func (a *App) POST(path string, handler HandlerFunc) {
	a.router.addRoute("POST", path, handler)
}

func (a *App) Use(middleware ...MiddlewareFunc) {
	a.middleware = append(a.middleware, middleware...)
}

func (a *App) Run() error {
	// Load templates before starting the server
	if err := a.template.Load(); err != nil {
		return err
	}

	server := newServer(a)
	return server.start()
}
