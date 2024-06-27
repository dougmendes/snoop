package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("views", "index.html")))
	tmpl.Execute(w, nil)
}

func Echo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		input := r.FormValue("textInput")
		w.Write([]byte(input))
	}
}
