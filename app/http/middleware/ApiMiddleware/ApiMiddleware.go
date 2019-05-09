package ApiMiddleware

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type httpHandlerFunc func(http.ResponseWriter, *http.Request)
type Response struct{
	Status int `json:status`
	Message string `json:message`
}

func Auth(next httpHandlerFunc) httpHandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("token")
		var response Response
		if token == "fahri" {
			fmt.Println("Logged in")
			next(res, req)
		}else{
			fmt.Println("Not authenticated")
			response.Status = 304
			response.Message = "Not authenticated"

			res.Header().Set("Content-type", "application/json")
			json.NewEncoder(res).Encode(response)
			res.WriteHeader(http.StatusUnauthorized)
		}
	}
}