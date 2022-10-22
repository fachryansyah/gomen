package app

import (
	"errors"
	"gomen/app/dtos"
	"gomen/utils"
	"time"
)

func LoginService(req dtos.LoginRequest) (utils.Response, error) {
	users := map[string]string{
		"username": "Jhon Doe",
		"email":    "jhondoe@email.com",
		"password": "password",
	}

	if req.Email != users["email"] || req.Password != users["password"] {
		return utils.ForbiddenResponse(nil, "Invalid email or password"), errors.New("Invalid email or password")
	}

	return utils.SuccessResponse(dtos.LoginResponse{
		Email:    users["email"],
		Username: users["username"],
		ApiKey:   time.Now().String(),
	}), nil
}
