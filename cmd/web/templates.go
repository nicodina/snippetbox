package main

import (
	"html/template"
	"path/filepath"
	
	"github.com/nicodina/snippetbox/pkg/models"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		t, err := template.ParseFiles(page)
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