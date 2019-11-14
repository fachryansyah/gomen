package database

import (
	"gomen/app/models"
)

// Migrate : migrate the table
func Migrate() {

	db := Connect()
	defer db.Close()

	db.AutoMigrate(&models.User{})
}
