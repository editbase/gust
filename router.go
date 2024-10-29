// /router.go
// Router manages HTTP routes and their corresponding handlers

package stardust

type Router struct {
	routes map[string]map[string]HandlerFunc
}

// newRouter creates a new Router instance with initialized routes map
func newRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

// addRoute registers a new route with method, path and handler
func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
}

// match finds the handler for given method and path
// Returns nil if no matching route is found
func (r *Router) match(method, path string) HandlerFunc {
	if routes, ok := r.routes[method]; ok {
		if handler, ok := routes[path]; ok {
			return handler
		}
	}
	return nil
}
