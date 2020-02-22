package app

import (
	"net/http"
)

func (a *App) HandleHome(w http.ResponseWriter, r *http.Request) {
	a.html(w, "home.page", M{})
}

//HandlePrograms is good for your life
func (a *App) HandlePrograms(w http.ResponseWriter, r *http.Request) {
	a.html(w, "programs.page", M{})
}
