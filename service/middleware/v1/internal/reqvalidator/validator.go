package reqvalidator

import (
	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	Validator *validator.Validate
}

func (v RequestValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	if errs := v.Validator.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ErrorResponse{
				Error:       true,
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			})
		}
	}
	return validationErrors
}
