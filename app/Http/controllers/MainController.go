package controller

import (
	"encoding/json"
	"gomen/app/models"
	"gomen/database"
	"log"
	"net/http"
)

type MainController struct{}

// GetUser : getting user data
func (MainController) GetUser(res http.ResponseWriter, req *http.Request) {

	user := []models.User{}
	var responseUser models.Response

	db := database.Connect()
	defer db.Close()

	db.Find(&user)

	for _, item := range user {
		log.Printf("%+v", item.Fullname)
	}

	responseUser.Status = 200
	responseUser.Message = "OKE!"
	responseUser.Data = user

	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(responseUser)

}

// CreateUser : Create new user to db
func (MainController) CreateUser(res http.ResponseWriter, req *http.Request) {

	var responseUser models.Response

	user := models.User{Fullname: "tony", Email: "tonystark@gmail.com", Phone: "081382355029"}

	db := database.Connect()
	defer db.Close()

	db.Create(&user)

	responseUser.Status = 200
	responseUser.Message = "OKE!"

	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(responseUser)
}

// Hello : just hello
type Hello struct {
	Message string
}

// Test : just test
func (MainController) Test(res http.ResponseWriter, req *http.Request) {
	var message Hello
	message.Message = "test gan"
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(message)
}
