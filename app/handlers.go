package app

import (
	"net/http"
)

func (a *App) HandleHome(w http.ResponseWriter, r *http.Request) {
	a.html(w, "home.page", M{})
}

func (a *App) HandlePrograms(w http.ResponseWriter, r *http.Request) {
	a.html(w, "programs.page", M{})
}
func (a *App)  HandleStructure(w http.ResponseWriter, r *http.Request) {
	a.html(w, "structure.page", M{})
}
