// /util/engine/standard.go

package engine

import (
	"net/http"

	"github.com/editbase/gust/util/render"
)

type StandardEngine struct {
	renderer *render.Renderer
	mux      *http.ServeMux
}

func NewStandardEngine() *StandardEngine {
	e := &StandardEngine{
		renderer: render.NewRenderer(),
		mux:      http.NewServeMux(),
	}
	e.RegisterRoutes()
	return e
}

func (e *StandardEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.mux.ServeHTTP(w, r)
}

func (e *StandardEngine) RegisterRoutes() {
	e.mux.HandleFunc("/", e.handleIndex())
	e.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func (e *StandardEngine) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e.renderer.Render(w, "index.html", nil)
	}
}
