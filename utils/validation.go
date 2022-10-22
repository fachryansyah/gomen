package utils

import (
	"github.com/go-playground/validator"
)

func ValidateRequest(source interface{}) Response {
	var errors []ValidationErrorResponse
	validate := validator.New()

	if err := validate.Struct(source); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var field ValidationErrorResponse
			field.FailedField = err.StructField()
			field.Tag = err.Tag()
			field.Value = err.Param()
			errors = append(errors, field)
		}
	}

	return BadRequestResponse(errors, "Validation error")
}
