package main

import (
	"os"

	"gomen/core"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	args := os.Args[1:]
	core.Command(args)
}
