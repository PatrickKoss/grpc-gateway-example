package validator

import (
	"testing"

	"github.com/patrickkoss/grpc-gateway-example/domain"
)

func TestProductValidation(t *testing.T) {
	v := NewValidator()

	testCases := []struct {
		name    string
		product domain.Product
		isValid bool
	}{
		{
			name: "empty id",
			product: domain.Product{
				Id:          "",
				Description: "Test product",
				Price:       99.99,
				Name:        "testproduct",
			},
			isValid: false,
		},
		{
			name: "name longer than 50 characters",
			product: domain.Product{
				Id:          "123",
				Description: "Test product",
				Price:       99.99,
				Name:        "testproducttestproducttestproducttestproducttestproducttestproduct",
			},
			isValid: false,
		},
		{
			name: "price below zero",
			product: domain.Product{
				Id:          "123",
				Description: "Test product",
				Price:       -0.01,
				Name:        "testproduct",
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			checkValidation(t, v, tc.product, tc.isValid)
		})
	}
}
