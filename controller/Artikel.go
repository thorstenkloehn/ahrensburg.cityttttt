package controller

import (
	"bytes"
	embed "github.com/13rac1/goldmark-embed"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/yuin/goldmark"
	"io/ioutil"
	"net/http"
	"text/template"
)

var ()

var (
	view1, _ = template.ParseGlob("view/*")
)

func (start *Website) Artikel(w http.ResponseWriter, r *http.Request) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			embed.New(),
		),
	)

	content, _ := ioutil.ReadFile("docs/index" + ".md")

	start.Titel = viper.GetString("Website_Name")
	var buf bytes.Buffer
	markdown.Convert(content, &buf)
	start.Titel = viper.GetString("Website_Name")
	start.Inhalt = buf.String()
	view.ExecuteTemplate(w, "docs.html", start)
}

func (start *Website) Artikels(w http.ResponseWriter, r *http.Request) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			embed.New(),
		),
	)

	vars := mux.Vars(r)
	content, _ := ioutil.ReadFile("docs/" + vars["Artikel"] + ".md")
	var buf1 bytes.Buffer
	markdown.Convert(content, &buf1)
	start.Titel = viper.GetString("Website_Name")
	start.Inhalt = buf1.String()
	view.ExecuteTemplate(w, "docs.html", start)

}
