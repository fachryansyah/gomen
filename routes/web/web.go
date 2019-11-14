package web

import (
	. "gomen/app/Http/controllers"
	. "gomen/app/Http/middleware"

	"github.com/gorilla/mux"
)

// SetRoutes : setting route
func SetRoutes() *mux.Router {

	router := mux.NewRouter()

	/*
		Masukkan Route kamu disini.
	*/
	router.HandleFunc("/user", MainController{}.GetUser).Methods("GET")
	router.HandleFunc("/user/create", MainController{}.CreateUser).Methods("GET")
	router.HandleFunc("/test", MainController{}.Test).Methods("GET")

	/*
		if you want to use middleware.
	*/
	router.HandleFunc("/auth/user/get", ApiMiddleware{}.Auth(MainController{}.GetUser)).Methods("GET")

	return router
}
