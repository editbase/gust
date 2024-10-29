// /internal/engine/template.go
// TemplateEngine implements template rendering functionality
package engine

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateEngine struct {
	dir      string             // templates directory
	template *template.Template // parse template cache
}

// NewTemplateEngine creates a new template engine instance
func NewTemplateEngine(dir string) *TemplateEngine {
	return &TemplateEngine{
		dir: dir,
	}
}

// Render executes template rendering with lazy template loading
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
