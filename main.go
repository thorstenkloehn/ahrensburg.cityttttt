package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/thorstenkloehn/ahrensburg.city/controller"
	"net/http"
)

func main() {

	var start controller.Website
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config/") // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.Set("DatenbankZugang", fmt.Sprintf("user=%s password=%s dbname=ahrensburg sslmode=disable", viper.Get("Postgres_User"), viper.Get("Postgress_Passwort")))
	viper.Set("MemberZugang", fmt.Sprintf("user=%s password=%s dbname=members sslmode=disable", viper.Get("Postgres_User"), viper.Get("Postgress_Passwort")))

	var dir string
	var gpx string
	var images string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")

	flag.StringVar(&gpx, "gpx", "./externe_daten/gpx", "the directory to serve files from. Defaults to the current dir")
	flag.StringVar(&images, "images", "./images/images", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	router := mux.NewRouter()
	router.PathPrefix("/images/images/").Handler(http.StripPrefix("/images/images/", http.FileServer(http.Dir(images))))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	router.PathPrefix("/gpx/").Handler(http.StripPrefix("/gpx/", http.FileServer(http.Dir(gpx))))

	router.HandleFunc("/", start.Startseite)
	router.HandleFunc("/docs/{Artikel}", start.Artikels)
	router.HandleFunc("/docs/", start.Artikel)
	router.HandleFunc("/javascript.js", start.Javascript)
	fmt.Println("http://localhost:5000")
	http.ListenAndServe(":5000", handlers.CompressHandler(router))

}
