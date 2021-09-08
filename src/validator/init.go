package validator

import "github.com/go-playground/validator"

var validate *validator.Validate

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func validationErrorFormat(err error) []*ErrorResponse {
	var errors []*ErrorResponse
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = e.StructNamespace()
			element.Tag = e.Tag()
			element.Value = e.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
