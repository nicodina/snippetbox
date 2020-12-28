package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"bytes"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.temaplateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	// Let's first write into a buffer to catch possible errors
	buff := new(bytes.Buffer)
	err := ts.Execute(buff, td)
	if err != nil {
		app.serverError(w, err)
	}

	buff.WriteTo(w)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Convenient wrapper for 404
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
