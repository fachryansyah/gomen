package main

import (
	"log"
	"fmt"
	"net/http"
	// "github.com/gorilla/handlers"
	"simple-golang-boilerplate-dev/routes"
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

	http.Handle("/", routes.SetRoutes())
	fmt.Println("Backend started on port : 9000")
	log.Fatal(http.ListenAndServe(":9000", routes.SetRoutes()))


	// log.Fatal(http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(routes.SetRoutes())))
}