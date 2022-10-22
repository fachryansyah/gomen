package core

import (
	"fmt"
	"gomen/transport/api"
	"net"
	"os"

	pb "gomen/transport/grpc"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func InitializeRestAPIServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Gomen 1.1",
	})

	apiEndpoint := app.Group("/api")
	v1Endpoint := apiEndpoint.Group("/v1")

	api := api.NewRoute()
	api.Routes(v1Endpoint)

	return app
}

func InitializeGRPCServer() {
	port := os.Getenv("GRPC_PORT")
	tcp, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("Failed to listen: " + port)
	}

	app := grpc.NewServer()
	pb.RegisterAuthServer(app, &pb.GRPC{})
	fmt.Printf("GRPC Listen at %v", tcp.Addr())
	if err := app.Serve(tcp); err != nil {
		fmt.Printf("Failed to serve GRPC %v \n", err)
	}
}
