package app

import (
	"bytes"
	"net/http"
)

// html renders an html template
func (a *App) html(w http.ResponseWriter, name string, data M) {
	t, ok := a.templates[name+".html"]
	if !ok {
		a.logger.Printf("template %s does not exist", name)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// render the template to a buffer
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		a.serverError(w, err)
		return
	}

	_, err := buf.WriteTo(w)
	if err != nil {
		a.serverError(w, err)
		return
	}
}

// serverError logs a server related error, then send a 500 internal server error response
func (a *App) serverError(w http.ResponseWriter, err error) {
	a.logger.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
