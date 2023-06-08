package validator

import (
	"net/mail"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// onlySmallLetters allow only small letters in the field string.
func onlySmallLetters(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	pattern := "^[a-z]+$"
	match, err := regexp.MatchString(pattern, name)
	if err != nil {
		return false
	}

	return match
}

// isValidEmail validate for syntax correctness.
func isValidEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	_, err := mail.ParseAddress(email)

	return err == nil
}
