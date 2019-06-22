package web

import (
	"github.com/gorilla/mux"
	. "gomen/app/Http/Middleware"
	. "gomen/app/Http/Controllers"
)

func SetRoutes() *mux.Router {

	router := mux.NewRouter()

	/*
		Masukkan Route kamu disini.
	*/
	router.HandleFunc("/user/get", MainController{}.GetUser).Methods("GET")
	router.HandleFunc("/test", MainController{}.Test).Methods("GET")

	/*
		if you want to use middleware.
	*/
 	router.HandleFunc("/auth/user/get", ApiMiddleware{}.Auth(MainController{}.GetUser)).Methods("GET")

	return router
}