package app

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"mime/multipart"
	"net/http"
	"os"
)

type staticFileServer struct {
	fs http.FileSystem
}

func (sfs staticFileServer) Open(name string) (http.File, error) {
	f, err := sfs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}

	return f, nil
}

// Static handles static files requests
func Static(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileserver := http.FileServer(staticFileServer{http.Dir(path)})
		http.StripPrefix("/static", fileserver).ServeHTTP(w, r)
	}
}

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

func tajweedLevel(v string) string {
	switch v {
	case "1":
		return "ضعيف"
	case "2":
		return "متوسط"
	case "3":
		return "مستحسن"
	case "4":
		return "ضعيف جدا"
	}
	return ""
}

func imgToBase64(f multipart.File, fh *multipart.FileHeader) string {
	buf := make([]byte, fh.Size)
	imgReader := bufio.NewReader(f)
	imgReader.Read(buf)
	imgBase64 := base64.StdEncoding.EncodeToString(buf)
	return imgBase64
}
