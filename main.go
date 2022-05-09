package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/thorstenkloehn/ahrensburg.city/alphaFunktion"
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
	viper.Set("hallo", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", viper.Get("Postgres_User"), viper.Get("Postgress_Passwort"), viper.Get("Postgress_Datenbank")))

	var dir string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	router.HandleFunc("/", start.Startseite)
	fmt.Println("http://localhost:5000")
	alphaFunktion.Testseite()
	http.ListenAndServe(":5000", router)
}
