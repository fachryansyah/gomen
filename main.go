package main

import (

	// "github.com/gorilla/handlers"
	"gomen/core"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	args := os.Args[1:]
	core.Command(args)

}
