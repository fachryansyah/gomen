package grpc

import (
	"context"
	"errors"
	"gomen/app"
	"gomen/app/dtos"
	"gomen/utils"
	"reflect"
)

type GRPC struct {
	UnimplementedAuthServer
}

func (g *GRPC) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	validate := utils.ValidateRequest(req)
	if !reflect.ValueOf(validate.Data).IsNil() {
		return nil, errors.New(validate.Message)
	}

	result, err := app.LoginService(dtos.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, errors.New(result.Message)
	}

	return &LoginResponse{
		User: &User{
			Username: result.Data.(dtos.LoginResponse).Username,
			Email:    result.Data.(dtos.LoginResponse).Email,
			ApiKey:   result.Data.(dtos.LoginResponse).ApiKey,
		},
	}, nil
}
