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
	bd, _ := time.Parse(time.RFC3339, r.FormValue("birthDate"))
	//if err != nil {
	//	a.serverError(w, err)
	//	return
	//}
	af := applicationForm{
		FullName:     r.FormValue("fullName"),
		BirthPlace:   r.FormValue("birthPlace"),
		EducAndJob:   r.FormValue("educAndJob"),
		Address:      r.FormValue("address"),
		Phone:        r.FormValue("phone"),
		Cin:          r.FormValue("cin"),
		Email:        r.FormValue("email"),
		TajweedLevel: r.FormValue("tajweedLevel"),
		HifdAmount:   r.FormValue("hifdAmount"),
		Reason:       r.FormValue("reason"),
		ImgName:      r.FormValue("imgName"),
		BirthDate:    bd,
	}

	errors := validate.Validate(&af)
	if errors.HasAny() {
		data := M{"af": af, "errors": errors}
		a.html(w, "appform.page", data)
	} else {
		http.Redirect(w, r, "/application-form/print", http.StatusOK)
	}
}
