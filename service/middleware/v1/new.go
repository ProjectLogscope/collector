package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/hardeepnarang10/collector/service/middleware/v1/internal/reqvalidator"
)

type middleware struct {
	validator reqvalidator.RequestValidator
}

func New() Middleware {
	return &middleware{
		validator: reqvalidator.RequestValidator{
			Validator: validator.New(validator.WithRequiredStructEnabled()),
		},
	}
}
