package core

import (
	"gomen/app/dtos"
	"gomen/database"
	"os"
)

func InitializeDatabase() {
	params := database.DBModel{
		ServerMode:   os.Getenv("APP_MODE"),
		Driver:       os.Getenv("DB_DRIVER"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		Name:         os.Getenv("DB_NAME"),
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASS"),
		MaxIdleConn:  100, // Maximum number of connections in the idle connection pool
		MaxOpenConn:  200, // Maximum number of open connections to the database
		ConnLifeTime: 150, // Maximum amount of time a connection may be reused
	}

	db, err := params.OpenDB()
	if err != nil {
		panic(err)
	}

	params.NewMigration(dtos.User{})

	database.DB = db
}
