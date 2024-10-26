// template.go
package gust

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type Template struct {
	dir       string
	templates *template.Template
	mutex     sync.RWMutex
}

func NewTemplate(dir string) *Template {
	return &Template{
		dir: dir,
	}
}

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
