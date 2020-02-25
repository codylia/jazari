package app

import (
	"net/http"
	"time"

	"github.com/oussama4/validate/v4"
)

func (a *App) HandleHome(w http.ResponseWriter, r *http.Request) {
	a.html(w, "home.page", M{})
}

func (a *App) HandlePrograms(w http.ResponseWriter, r *http.Request) {
	a.html(w, "programs.page", M{})
}
func (a *App) HandleStructure(w http.ResponseWriter, r *http.Request) {
	a.html(w, "structure.page", M{})
}

func (a *App) HandleApplicationForm(w http.ResponseWriter, r *http.Request) {
	a.html(w, "appform.page", M{})
}

func (a *App) HandleAppFormPost(w http.ResponseWriter, r *http.Request) {
	bd, err := time.Parse(time.RFC3339, r.FormValue("birthDate"))
	if err != nil {
		a.serverError(w, err)
		return
	}
	af := applicationForm{
		fullName:     r.FormValue("fullName"),
		birthPlace:   r.FormValue("birthPlace"),
		educAndJob:   r.FormValue("educAndJob"),
		address:      r.FormValue("address"),
		phone:        r.FormValue("phone"),
		cin:          r.FormValue("cin"),
		email:        r.FormValue("email"),
		tajweedLevel: r.FormValue("tajweedLevel"),
		hifdAmount:   r.FormValue("hifdAmount"),
		reason:       r.FormValue("reason"),
		// imgName:      r.FormValue("imgName"),
		birthDate: bd,
	}

	errors := validate.Validate(&af)
	if errors.HasAny() {
		data := M{"af": af, "errors": errors}
		a.html(w, "appform.page", data)
	} else {
		http.Redirect(w, r, "/application-form/print", http.StatusOK)
	}
}
