package controller

import (
	"bytes"
	"github.com/13rac1/goldmark-embed"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"io/ioutil"
	"net/http"
	"text/template"
)

var ()

var (
	view1, _ = template.ParseGlob("view/*")
)

func (start *Website) Artikel(w http.ResponseWriter, r *http.Request) {

	content, _ := ioutil.ReadFile("docs/index" + ".md")

	start.Titel = viper.GetString("Website_Name")
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			embed.New(),
		),
	)
	var buf bytes.Buffer

	markdown.Convert(content, &buf)
	start.Titel = viper.GetString("Website_Name")
	start.Inhalt = string(buf.Bytes())
	view.ExecuteTemplate(w, "docs.html", start)
}

func (start *Website) Artikels(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	content, _ := ioutil.ReadFile("docs/" + vars["Artikel"] + ".md")
	var buf1 bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			embed.New(),
		),
	)
	markdown.Convert(content, &buf1)
	start.Titel = viper.GetString("Website_Name")
	start.Inhalt = string(buf1.Bytes())
	view.ExecuteTemplate(w, "docs.html", start)

}
