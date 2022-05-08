package controller

import (
	"github.com/spf13/viper"
	_ "github.com/thorstenkloehn/ahrensburg.city/model"
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
	start.Titel = viper.GetString("Website_Name")
	view.ExecuteTemplate(w, "startseite.html", start)
}