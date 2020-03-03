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

//program
func (a *App) HandleProgramme(w http.ResponseWriter, r *http.Request) {
	a.html(w, "programme.page", M{})
}

func (a *App) HandleCourses(w http.ResponseWriter, r *http.Request) {
	a.html(w, "courses.page", M{})
}
func (a *App) HandleReading(w http.ResponseWriter, r *http.Request) {
	a.html(w, "reading.page", M{})
}

func (a *App) HandleApplicationForm(w http.ResponseWriter, r *http.Request) {
	a.html(w, "appform.page", M{})
}

func (a *App) HandleFormToPrint(w http.ResponseWriter, r *http.Request) {
	a.html(w, "formToPrint.page", M{})
}

func (a *App) HandleAppFormPost(w http.ResponseWriter, r *http.Request) {
	fn := ""
	f, fh, err := r.FormFile("imgName")
	if err == nil {
		fn = fh.Filename
		defer f.Close()
	}
	af := applicationForm{
		FullName:     r.FormValue("fullName"),
		BirthPlace:   r.FormValue("birthPlace"),
		Job:          r.FormValue("job"),
		Educ:         r.FormValue("educ"),
		Address:      r.FormValue("address"),
		Phone:        r.FormValue("phone"),
		Cin:          r.FormValue("cin"),
		Email:        r.FormValue("email"),
		TajweedLevel: r.FormValue("tajweedLevel"),
		HifdAmount:   r.FormValue("hifdAmount"),
		Reason:       r.FormValue("reason"),
		ImgName:      fn,
		BirthDate:    r.FormValue("birthDate"),
	}
	errors := validate.Validate(&af)
	if err == http.ErrMissingFile {
		errors.Add("img_name", "المرجو اضافة الصورة")
	}
	if errors.HasAny() {
		data := M{"af": af, "errors": errors}
		a.html(w, "appform.page", data)
	} else {
		now := time.Now()
		af.TajweedLevel = tajweedLevel(af.TajweedLevel)
		img := imgToBase64(f, fh)
		data := M{"af": af, "now": now.Format("2006-01-02"), "img": img}
		a.html(w, "formToPrint.page", data)
	}
}
