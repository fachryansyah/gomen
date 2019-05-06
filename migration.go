package main

import (
	"log"
	"gomen/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.Connect()
	defer db.Close()

	_, err := db.Query("CREATE TABLE users ( id integer(12) primary key auto_increment, name varchar(191), email varchar(121), phone varchar(13) )")
	if err != nil {
		log.Print(err)
	}
}