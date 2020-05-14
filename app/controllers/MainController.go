package controllers

import (
	"context"
	service "gomen/app/services"
	"log"
)

type MainController struct{}

func (MainController) Register(ctx context.Context, param *service.User) (*service.Empty, error) {
	log.Println("User service has registered", param.String())
	return new(service.Empty), nil
}

func (MainController) List(ctx context.Context, void *service.Empty) (*service.UserList, error) {
	user := &service.User{
		Id:       12,
		Fullname: "Fahriansyah",
		Email:    "fachryansyah123@gmail.com",
		Phone:    "0813-****-****",
	}

	userList := &service.UserList{
		List: []*service.User{
			user,
		},
	}

	return userList, nil
}
