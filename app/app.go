package app

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	logger    *log.Logger
	templates map[string]*template.Template
}

// Start creates an App instance and starts an http server
func Start() {
	// create the app
	l := log.New(os.Stdout, "LOGGER: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		l.Println(err)
	}
	cache, err := newTemplateCache(os.Getenv("TEMPLATES_PATH"))
	if err != nil {
		l.Fatalln(err)
	}

	app := App{
		logger:    l,
		templates: cache,
	}

	// start the http server
	srv := &http.Server{
		Addr:    os.Getenv("ADDRESS"),
		Handler: app.routes(),
	}

	app.logger.Printf("starting server on %s", os.Getenv("ADDRESS"))
	err = srv.ListenAndServe()
	if err != nil {
		app.logger.Fatalln(err)
	}
}
