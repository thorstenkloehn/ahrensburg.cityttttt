package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"log"
	"net/http"
)

func Artikel(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile("docs/index" + ".md")
	if err != nil {
		log.Fatal(err)
	}

	output := blackfriday.Run(content)
	fmt.Fprint(w, string(output))
}

func Artikels(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	content, err := ioutil.ReadFile("docs/" + vars["Artikel"] + ".md")
	if err != nil {
		log.Fatal(err)
	}

	output := blackfriday.Run(content)
	fmt.Fprint(w, string(output))

}
