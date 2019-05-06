package MainController

import (
	"log"
	"net/http"
	"encoding/json"
	"gomen/database"
	"gomen/app/model/UserModel"
)

func GetUser(res http.ResponseWriter, req *http.Request) {

	var users UserModel.User
	var dataUsers []UserModel.User
	var responseUser UserModel.Response

	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Fullname, &users.Email, &users.Phone); err != nil {
			log.Fatal(err)
		} else {
			dataUsers = append(dataUsers, users)
		}
	}

	responseUser.Status = 200
	responseUser.Message = "OKE!"
	responseUser.Data = dataUsers

	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(responseUser)

}