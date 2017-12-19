package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	templates = make(map[string]*template.Template)
)

func initTemplate(templatePath string) {
	layouts, err := filepath.Glob(templatePath + "layout/*.html")
	checkErr(err, "filepath.Glob err")
	for _, layout := range layouts {
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout))
	}
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, correct := templates[name]
	catchFalse(correct, name+" template not exist")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, name, data)
}

func renderNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
