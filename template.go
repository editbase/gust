// /template.go
// Template manages HTML template parsing and rendering

package stardust

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type Template struct {
	dir       string             // template directory
	templates *template.Template // parsed templates cache
	mutex     sync.RWMutex       // ensures thread-safe template operations
}

// NewTemplate creates a new Template instance for the specified directory
func NewTemplate(dir string) *Template {
	return &Template{
		dir: dir,
	}
}

// Load parses all HTML templates in the template directory
func (t *Template) Load() error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// Find all template files
	pattern := filepath.Join(t.dir, "*.html")
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}

	t.templates = tmpl
	return nil
}

// Render executes a template with the given name and data
func (t *Template) Render(w http.ResponseWriter, name string, data interface{}) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.templates == nil {
		if err := t.Load(); err != nil {
			return err
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return t.templates.ExecuteTemplate(w, name, data)
}
