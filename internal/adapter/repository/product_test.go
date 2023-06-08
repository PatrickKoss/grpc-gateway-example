package repository

import (
	"testing"

	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {
	logger := logging.New()
	repo := NewProductRepository(logger)

	product := &domain.Product{
		Id:    "1",
		Name:  "Test Product",
		Price: 100,
	}

	// Test Create
	err := repo.Create(product)
	assert.NoError(t, err)

	// Test Get
	storedProduct, err := repo.Get("1")
	assert.NoError(t, err)
	assert.NotNil(t, storedProduct)
	assert.Equal(t, product, storedProduct)

	// Test List
	products, err := repo.List()
	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, product, &products[0])

	// Test Delete
	err = repo.Delete("1")
	assert.NoError(t, err)

	// Test Get after Delete
	storedProduct, err = repo.Get("1")
	assert.Error(t, err)
	assert.Nil(t, storedProduct)

	// Test List after Delete
	products, err = repo.List()
	assert.NoError(t, err)
	assert.Len(t, products, 0)
}

func TestProductRepositoryErrorCases(t *testing.T) {
	logger := logging.New()
	repo := NewProductRepository(logger)

	product := &domain.Product{
		Id:    "1",
		Name:  "Test Product",
		Price: 100,
	}

	// Test Get with non-existent product
	storedProduct, err := repo.Get("non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNotFound, err)
	assert.Nil(t, storedProduct)

	// Test Create
	err = repo.Create(product)
	assert.NoError(t, err)

	// Test creating a duplicate product
	duplicateProduct := &domain.Product{
		Id:    "1",
		Name:  "Duplicate Product",
		Price: 200,
	}
	err = repo.Create(duplicateProduct)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrAlreadyExists, err)

	// Test Delete with non-existent product
	err = repo.Delete("non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNotFound, err)
}
