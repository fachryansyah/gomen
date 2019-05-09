package web

import (
	"github.com/gorilla/mux"
	"gomen-dev/app/http/middleware/ApiMiddleware"
	"gomen-dev/app/http/controller/MainController"
)

func SetRoutes() *mux.Router {

	router := mux.NewRouter()

	/*
		Masukkan Route kamu disini.
	*/
	router.HandleFunc("/user/get", ApiMiddleware.Auth(MainController.GetUser)).Methods("GET")


	return router
}