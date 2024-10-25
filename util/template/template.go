// /util/template/template.go

package template

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"sync"
)

// TemplateError represents a template error
type TemplateError struct {
	Name string
	Err  error
}

func (e *TemplateError) Error() string {
	return fmt.Sprintf("template error in %s: %v", e.Name, e.Err)
}

// Manager handles template loading and caching
type Manager struct {
	templates  map[string]*template.Template
	mutex      sync.RWMutex
	basePath   string
	funcMap    template.FuncMap
	extensions []string
}

// NewManager creates a new template manager
func NewManager(basePath string) *Manager {
	return &Manager{
		templates:  make(map[string]*template.Template),
		basePath:   basePath,
		extensions: []string{".html", ".tmpl"},
		funcMap:    make(template.FuncMap),
	}
}

// AddFunc adds a template function
func (m *Manager) AddFunc(name string, fn interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.funcMap == nil {
		m.funcMap = make(template.FuncMap)
	}
	m.funcMap[name] = fn
}

// Load loads all templates from the base path
func (m *Manager) Load() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Clear existing templates
	m.templates = make(map[string]*template.Template)

	// Load layouts first
	layouts, err := filepath.Glob(filepath.Join(m.basePath, "layouts/*.html"))
	if err != nil {
		return &TemplateError{Name: "layouts", Err: err}
	}

	// Load pages
	pages, err := filepath.Glob(filepath.Join(m.basePath, "pages/*.html"))
	if err != nil {
		return &TemplateError{Name: "pages", Err: err}
	}

	// Parse each page with layouts
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.New(name).Funcs(m.funcMap)

		// Parse layouts first
		for _, layout := range layouts {
			_, err := tmpl.ParseFiles(layout)
			if err != nil {
				return &TemplateError{Name: layout, Err: err}
			}
		}

		// Parse the page itself
		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return &TemplateError{Name: page, Err: err}
		}

		m.templates[name] = tmpl
	}

	return nil
}

// Render renders a template with the given name and data
func (m *Manager) Render(w io.Writer, name string, data interface{}) error {
	m.mutex.RLock()
	tmpl, exists := m.templates[name]
	m.mutex.RUnlock()

	if !exists {
		return &TemplateError{
			Name: name,
			Err:  fmt.Errorf("template not found"),
		}
	}

	return tmpl.Execute(w, data)
}

// RenderPartial renders a partial template
func (m *Manager) RenderPartial(w io.Writer, name string, data interface{}) error {
	m.mutex.RLock()
	tmpl, exists := m.templates[name]
	m.mutex.RUnlock()

	if !exists {
		return &TemplateError{
			Name: name,
			Err:  fmt.Errorf("template not found"),
		}
	}

	return tmpl.ExecuteTemplate(w, "content", data)
}

// Hot reload for development
func (m *Manager) ReloadOnChange() error {
	return m.Load()
}

// GetTemplate returns a template by name
func (m *Manager) GetTemplate(name string) (*template.Template, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	tmpl, exists := m.templates[name]
	if !exists {
		return nil, &TemplateError{
			Name: name,
			Err:  fmt.Errorf("template not found"),
		}
	}

	return tmpl, nil
}
