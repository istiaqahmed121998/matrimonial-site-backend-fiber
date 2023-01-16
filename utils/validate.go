package utils

import "gopkg.in/go-playground/validator.v9"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(bodyJson interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(bodyJson)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
