package main

import (
	"log"
	"fmt"
	"net/http"
	// "github.com/gorilla/handlers"
	"gomen/routes/web"
	"gomen/routes/socket"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	/*
		CORS ORIGIN HEADER.
		if you need this uncomment the code below, and uncomment "github.com/gorilla/handlers"
	*/
	
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// set for static file
    fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// set routes for api
	http.Handle("/api/", web.SetRoutes())

	// set routes for socket
	http.Handle("/socket.io/", socket.SetRoutes())

	fmt.Println("Backend server started on port : 9000")

	log.Fatal(http.ListenAndServe(":9000", nil))

	// log.Fatal(http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(routes.SetRoutes())))
}