package main

import (

	// "github.com/gorilla/handlers"

	"os"

	"gomen/core"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	args := os.Args[1:]
	core.Command(args)

}
