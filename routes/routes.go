package routes

import (
	"github.com/gorilla/mux"
	"simple-golang-boilerplate-dev/app/http/controller/MainController"
)

func SetRoutes() *mux.Router {

	router := mux.NewRouter()

	/*
		Masukkan Route kamu disini.
	*/
	router.HandleFunc("/user/get", MainController.GetUser).Methods("GET")

	return router
}