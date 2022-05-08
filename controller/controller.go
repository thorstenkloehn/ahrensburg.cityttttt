package controller

import (
	"html/template"
	"net/http"
)

type Website struct {
	Titel string
}

var (
	view, _ = template.ParseGlob("view/*")
)

func (start *Website) Startseite(w http.ResponseWriter, r *http.Request) {
	start.Titel = "Hallo"
	view.ExecuteTemplate(w, "index.html", start)
}
