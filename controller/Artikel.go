package controller

import (
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
	"github.com/spf13/viper"
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

	output := blackfriday.Run(content)
	start.Titel = viper.GetString("Website_Name")
	start.Inhalt = string(output)
	view.ExecuteTemplate(w, "docs.html", start)
}

func (start *Website) Artikels(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	content, _ := ioutil.ReadFile("docs/" + vars["Artikel"] + ".md")

	output := blackfriday.Run(content)
	start.Titel = viper.GetString("Website_Name")
	start.Inhalt = string(output)
	view.ExecuteTemplate(w, "docs.html", start)

}
