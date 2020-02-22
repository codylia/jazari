package app

import (
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/static/*", Static(os.Getenv("STATIC_PATH")))
	r.Get("/", a.HandleHome)
	r.Get("/programs", a.HandlePrograms)

	return r
}
