package database

import(
	"fmt"
	"log"
	"database/sql"
)

func Connect() *sql.DB {

	var username string = "root"
	var password string = ""
	var host string = "localhost"
	var database string = "golang"

	db, err := sql.Open("mysql", fmt.Sprintf("%s@%stcp(%s:3306)/%s", username, password, host, database))

	if err != nil {
		log.Fatal(err)
	}

	return db

}