// /util/render/render.go

package render

import (
	"html/template"
	"net/http"
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer() *Renderer {
	// First parse base templates
	tmpl := template.New("").Funcs(template.FuncMap{
		// Add any custom functions here
	})

	// Parse all templates
	tmpl = template.Must(tmpl.ParseGlob("templates/layouts/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("templates/pages/*.html"))

	return &Renderer{
		templates: tmpl,
	}
}

func (r *Renderer) Render(w http.ResponseWriter, name string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return r.templates.ExecuteTemplate(w, name, data)
}
