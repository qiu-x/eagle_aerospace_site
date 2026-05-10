package tmpl

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

var pageFiles = map[string]string{
	"home":        "templates/pages/home.tmpl",
	"hangar":      "templates/pages/hangar.tmpl",
	"scenarios":   "templates/pages/scenarios.tmpl",
	"cooperation": "templates/pages/cooperation.tmpl",
	"crew":        "templates/pages/crew.tmpl",
	"contact":     "templates/pages/contact.tmpl",
}

// Renderer owns one parsed *template.Template per page key
// It is safe for concurrent use.
type Renderer struct {
	pages map[string]*template.Template
}

// tr looks up a string in the localized chrome map. If the key is
// missing it returns the key itself so missing translations are
// visible in the rendered HTML rather than silently empty.
func tr(strings map[string]string, key string) string {
	if s, ok := strings[key]; ok {
		return s
	}
	return key
}

var funcMap = template.FuncMap{
	"tr": tr,
}

func NewRenderer(fsys fs.FS) (*Renderer, error) {
	pages := make(map[string]*template.Template, len(pageFiles))

	for key, pageFile := range pageFiles {
		t, err := template.New(key).Funcs(funcMap).ParseFS(
			fsys,
			"templates/base.tmpl",
			"templates/partials/*.tmpl",
			pageFile,
		)
		if err != nil {
			return nil, fmt.Errorf("parse template set %q: %w", key, err)
		}
		pages[key] = t
	}

	return &Renderer{pages: pages}, nil
}

func (r *Renderer) Render(w http.ResponseWriter, key string, data any) error {
	t, ok := r.pages[key]
	if !ok {
		return fmt.Errorf("unknown template set %q", key)
	}

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "base", data); err != nil {
		return fmt.Errorf("execute %q: %w", key, err)
	}

	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
	_, err := buf.WriteTo(w)
	return err
}
