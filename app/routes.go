package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/static/*", Static(a.cfg.StaticPath))
	r.Get("/", a.HandleHome)
	r.Get("/programs", a.HandlePrograms)
	r.Get("/programs2", a.HandlePrograms2)
	r.Get("/announcement", a.HandleAnnouncement)
	r.Get("/programme", a.HandleProgramme)
	r.Get("/courses", a.HandleCourses)
	r.Get("/reading", a.HandleReading)
	r.Get("/structure", a.HandleStructure)
	r.Get("/application-form", a.HandleApplicationForm)
	r.Get("/application-form/print", a.HandleFormToPrint)
	r.Post("/application-form", a.HandleAppFormPost)

	return r
}
