package api

import (
	"gomen/app"
	"gomen/app/dtos"
	"gomen/utils"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func (a *API) LoginHandler(c *fiber.Ctx) error {
	request := new(dtos.LoginRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BadRequestResponse(nil, err.Error()))
	}

	errors := utils.ValidateRequest(*request)
	if !reflect.ValueOf(errors.Data).IsNil() {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	result, err := app.LoginService(*request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
