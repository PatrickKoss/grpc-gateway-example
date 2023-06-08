package validator

import "github.com/go-playground/validator/v10"

type Validator interface {
	ValidateStruct(s interface{}) error
}

type customValidator struct {
	validate *validator.Validate
}

func (c customValidator) ValidateStruct(s interface{}) error {
	return c.validate.Struct(s)
}

func NewValidator() Validator {
	validate := validator.New()
	_ = validate.RegisterValidation("onlySmallLetters", onlySmallLetters)
	_ = validate.RegisterValidation("isValidEmail", isValidEmail)

	return customValidator{validate: validate}
}
