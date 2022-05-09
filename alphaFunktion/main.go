package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	var thorsten, name string
	connStr := "user=test dbname=test sslmode=disable password=test"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	error := db.Ping()
	if error != nil {
		panic(error)
	}
	quey := "select name , St_asgeojson(st_transform(geom,3857)) as geom from strassen;"

	rows, errtest := db.Query(quey)
	if errtest != nil {
		log.Fatal(errtest)
	}

	for rows.Next() {
		mama := rows.Scan(&name, &thorsten)
		if mama != nil {
			log.Fatal(mama)
		}
		fmt.Println(name, thorsten)
	}
}
