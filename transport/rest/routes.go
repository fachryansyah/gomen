package api

import (
	"github.com/gofiber/fiber/v2"
)

type API struct{}

func NewRoute() *API {
	return &API{}
}

func (a *API) Routes(r fiber.Router) {
	r.Get("/", a.HelloWorld)
	r.Post("/login", a.LoginHandler)
}
