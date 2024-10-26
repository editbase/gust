// router.go
package gust

type Router struct {
	routes map[string]map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *Router) match(method, path string) HandlerFunc {
	if routes, ok := r.routes[method]; ok {
		if handler, ok := routes[path]; ok {
			return handler
		}
	}
	return nil
}
