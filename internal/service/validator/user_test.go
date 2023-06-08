package validator

import (
	"testing"

	"github.com/patrickkoss/grpc-gateway-example/domain"
)

func TestUserValidation(t *testing.T) {
	v := NewValidator()

	testCases := []struct {
		name    string
		user    domain.User
		isValid bool
	}{
		{
			name: "empty name",
			user: domain.User{
				Id:      "123",
				Name:    "",
				Email:   "testuser@gmail.com",
				Surname: "testsurname",
			},
			isValid: false,
		},
		{
			name: "email without '@'",
			user: domain.User{
				Id:      "123",
				Name:    "testuser",
				Email:   "testusergmail.com",
				Surname: "testsurname",
			},
			isValid: false,
		},
		{
			name: "id longer than 63 characters",
			user: domain.User{
				Id:      "1234567890123456789012345678901234567890123456789012345678901234567890",
				Name:    "testuser",
				Email:   "testuser@gmail.com",
				Surname: "testsurname",
			},
			isValid: false,
		},
		{
			name: "surname with capital letters",
			user: domain.User{
				Id:      "123",
				Name:    "testuser",
				Email:   "testuser@gmail.com",
				Surname: "TestSurname",
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			checkValidation(t, v, tc.user, tc.isValid)
		})
	}
}

func checkValidation(t *testing.T, v Validator, s interface{}, expectedValid bool) {
	t.Helper()

	err := v.ValidateStruct(s)

	if expectedValid && err != nil {
		t.Fatalf("expected valid, got error: %v", err)
	}

	if !expectedValid && err == nil {
		t.Fatalf("expected invalid, got no error")
	}
}
