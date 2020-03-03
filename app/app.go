package app

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	TemplatesPath string
	StaticPath    string
}

type App struct {
	logger    *log.Logger
	templates map[string]*template.Template
	cfg       Config
}

// Start creates an App instance and starts an http server
func Start() {
	// config
	p := flag.String("p", "8080", "port number to listen to")
	s := flag.String("s", "./ui/static", "path to static files")
	t := flag.String("t", "./ui/templates", "path to templates")
	flag.Parse()

	c := Config{*p, *t, *s}

	// create the app
	l := log.New(os.Stdout, "LOGGER: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		l.Println(err)
	}
	cache, err := newTemplateCache(c.TemplatesPath)
	if err != nil {
		l.Fatalln(err)
	}

	app := App{
		logger:    l,
		templates: cache,
		cfg:       c,
	}

	// start the http server
	srv := &http.Server{
		Addr:    ":" + app.cfg.Port,
		Handler: app.routes(),
	}

	app.logger.Printf("starting server on %s", ":"+app.cfg.Port)
	err = srv.ListenAndServe()
	if err != nil {
		app.logger.Fatalln(err)
	}
}
