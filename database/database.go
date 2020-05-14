package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Connect : try to connect to database
func Connect() *gorm.DB {

	var username string = "username"
	var password string = "your-pass"
	var host string = "localhost"
	var database string = "gomen"

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:3306)/%s?parseTime=true", username, password, host, database))

	if err != nil {
		log.Fatal(err)
	}

	return db

}
