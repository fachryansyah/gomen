package api

import (
	"github.com/gofiber/fiber/v2"
)

func ProtectedAuth(c fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "TOKEN-EXAMPLE" {
		return c.Status(fiber.StatusForbidden).JSON(map[string]string{
			"message": "Forbidden Access",
		})
	}

	return c.Next()
}
