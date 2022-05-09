package alphaFunktion

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func Testseite() {

	var thorsten, name string
	fmt.Println(viper.Get("hallo"))
	db, err := sql.Open("postgres", viper.GetString("hallo"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
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
		fmt.Print("", name, thorsten)
	}

}
