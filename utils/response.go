package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Message string      `json:"string"`
	Status  uint16      `json:"status"`
	Data    interface{} `json:"data"`
}

type ValidationErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func SuccessResponse(payload interface{}) Response {
	return Response{
		Message: "Request successfully!",
		Status:  fiber.StatusOK,
		Data:    payload,
	}
}

func InternalServerErrorResponse(payload interface{}) Response {
	return Response{
		Message: "Something went wrong!",
		Status:  fiber.StatusInternalServerError,
		Data:    payload,
	}
}

func ForbiddenResponse(payload interface{}, message string) Response {
	return Response{
		Message: message,
		Status:  fiber.StatusForbidden,
		Data:    payload,
	}
}

func BadRequestResponse(payload interface{}, message string) Response {
	return Response{
		Message: message,
		Status:  fiber.StatusBadRequest,
		Data:    payload,
	}
}
