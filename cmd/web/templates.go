package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/nicodina/snippetbox/pkg/forms"
	"github.com/nicodina/snippetbox/pkg/models"
)

type templateData struct {
	AuthenticatedUser *models.User
	CurrentYear int
	Form *forms.Form
	Flash string
	Snippet *models.Snippet
	Snippets []*models.Snippet
	CSRFToken string
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		t, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		t, err = t.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		t, err = t.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		cache[name] = t
	}

	return cache, nil
}