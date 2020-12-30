package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/nicodina/snippetbox/pkg/forms"
	"github.com/nicodina/snippetbox/pkg/models"
)

type templateData struct {
	CurrentYear int
	Form *forms.Form
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:02")
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