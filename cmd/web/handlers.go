package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nicodina/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Snippets: snippets}
	app.render(w, r, "home.page.html", data)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Snippet: snippet}
	app.render(w, r, "show.page.html", data)
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request){
	app.render(w, r, "create.page.html", nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Retrieving form data
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	// Create a new snippet in the database
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}