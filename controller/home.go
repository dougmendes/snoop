package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var tmpl = template.Must(template.ParseFiles(filepath.Join("views", "index.html")))

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func Echo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		input := r.FormValue("textInput")
		w.Write([]byte(input))
	}
}
