package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	mockrepository "github.com/patrickkoss/grpc-gateway-example/internal/adapter/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestProductServiceWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepository.NewMockProductRepository(ctrl)
	logger := logging.New()
	service := NewProductService(logger, mockRepo)

	product := &domain.Product{
		Id:    "1",
		Name:  "Test Product",
		Price: 100,
	}

	mockRepo.EXPECT().Create(product).Return(nil)
	err := service.Create(product)
	assert.NoError(t, err)

	mockRepo.EXPECT().Get("1").Return(product, nil)
	storedProduct, err := service.Get("1")
	assert.NoError(t, err)
	assert.NotNil(t, storedProduct)
	assert.Equal(t, product, storedProduct)

	mockRepo.EXPECT().List().Return([]domain.Product{*product}, nil)
	products, err := service.List()
	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, product, &products[0])

	mockRepo.EXPECT().Delete("1").Return(nil)
	err = service.Delete("1")
	assert.NoError(t, err)

	mockRepo.EXPECT().Get("1").Return(nil, errors.New("not found"))
	storedProduct, err = service.Get("1")
	assert.Error(t, err)
	assert.Nil(t, storedProduct)

	mockRepo.EXPECT().List().Return(nil, errors.New("empty list"))
	products, err = service.List()
	assert.Error(t, err)
	assert.Nil(t, products)
}

func TestProductServiceErrorCasesWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepository.NewMockProductRepository(ctrl)
	logger := logging.New()
	service := NewProductService(logger, mockRepo)

	product := &domain.Product{
		Id:    "1",
		Name:  "Test Product",
		Price: 100,
	}

	// Test Create Error
	mockRepo.EXPECT().Create(product).Return(errors.New("create error"))
	err := service.Create(product)
	assert.Error(t, err)

	// Test Get Error
	mockRepo.EXPECT().Get("non-existent-id").Return(nil, errors.New("not found"))
	storedProduct, err := service.Get("non-existent-id")
	assert.Error(t, err)
	assert.Nil(t, storedProduct)

	// Test List Error
	mockRepo.EXPECT().List().Return(nil, errors.New("list error"))
	products, err := service.List()
	assert.Error(t, err)
	assert.Nil(t, products)

	// Test Delete Error
	mockRepo.EXPECT().Delete("non-existent-id").Return(errors.New("not found"))
	err = service.Delete("non-existent-id")
	assert.Error(t, err)
}
