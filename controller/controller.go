package controller

import (
	"github.com/spf13/viper"
	"github.com/thorstenkloehn/ahrensburg.city/model"
	_ "github.com/thorstenkloehn/ahrensburg.city/model"
	"net/http"
	"text/template"
)

type Website struct {
	model.Website
}

var (
	view, _ = template.ParseGlob("view/*")
)

func (start *Website) Startseite(w http.ResponseWriter, r *http.Request) {
	start.Titel = viper.GetString("Website_Name")
	view.ExecuteTemplate(w, "startseite.html", start)
}

func (start *Website) Javascript(w http.ResponseWriter, r *http.Request) {
	view.ExecuteTemplate(w, "javascript.js", start)
}
