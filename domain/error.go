package domain

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound      = fmt.Errorf("not found")
	ErrAlreadyExists = fmt.Errorf("already exists")
)

type InvalidInputError struct {
	Message string
}

func (e *InvalidInputError) Error() string {
	return e.Message
}

func (e *InvalidInputError) Is(target error) bool {
	return errors.As(target, &e)
}
