package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateError(err error) map[string]string {
	result := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, v := range validationErrors {
			field := strings.ToLower(v.Field())
			result[field] = ValidateErrorMessage(field, v)
		}
	}
	return result
}

func ValidateErrorMessage(field string, errField validator.FieldError) string {
	switch errField.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return "Invalid format for " + field
	case "min":
		return field + " must be at least " + errField.Param() + " characters long"
	case "max":
		return field + " must be at most " + errField.Param() + " characters long"
	default:
		return field + " is invalid"
	}
}