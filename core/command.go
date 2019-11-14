package core

import (
	"fmt"
	"gomen/database"
)

// Command : run command line interface
func Command(cmd []string) {
	switch cmd[0] {
	case "serve":
		Serve()
	case "migrate":
		database.Migrate()
	case "test":
		fmt.Println("test")
	}
}
