package validation

import "github.com/go-playground/validator/v10"

type EchoValidator struct {
	validator *validator.Validate
}

func New() *EchoValidator {
	return &EchoValidator{validator: validator.New()}
}

func (v *EchoValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
