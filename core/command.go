package core

import (
	"fmt"
	"gomen/app/controllers"
	"gomen/database"
	"log"
	"net"
	"post-services/config"

	service "gomen/app/services"

	"google.golang.org/grpc"
)

// Command : run command line interface
func Command(cmd []string) {
	switch cmd[0] {
	case "serve":
		runRpcServer()
		break
	case "migrate":
		database.Migrate()
		break
	case "test":
		fmt.Println("test")
		break
	}
}

func runRpcServer() {
	srv := grpc.NewServer()
	var postSrv controllers.MainController
	service.RegisterUsersServer(srv, postSrv)

	log.Println("Starting GRPC Server at ", config.SERVICE_POST)

	listen, err := net.Listen("tcp", config.SERVICE_POST)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_POST, err)
	}

	log.Fatal(srv.Serve(listen))
}
