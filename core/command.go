package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Command : run command line interface
func Command(cmd []string) {
	switch cmd[0] {
	case "serve":
		serve()
		break
	case "test":
		fmt.Println("test")
		break
	}
}

func serve() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file, please provid it by copying .env.example")
	}

	go InitializeGRPCServer()

	restApi := InitializeRestAPIServer()
	log.Fatal(restApi.Listen(":" + os.Getenv("API_PORT")))
}
