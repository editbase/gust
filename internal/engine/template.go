// internal/engine/template.go
package engine

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateEngine struct {
	dir      string
	template *template.Template
}

func NewTemplateEngine(dir string) *TemplateEngine {
	return &TemplateEngine{
		dir: dir,
	}
}

func (e *TemplateEngine) Render(w http.ResponseWriter, name string, data interface{}) error {
	if e.template == nil {
		pattern := filepath.Join(e.dir, "*.html")
		tmpl, err := template.ParseGlob(pattern)
		if err != nil {
			return err
		}
		e.template = tmpl
	}
	return e.template.ExecuteTemplate(w, name, data)
}
