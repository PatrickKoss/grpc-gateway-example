package product

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/patrickkoss/grpc-gateway-example/domain"
	mockservice "github.com/patrickkoss/grpc-gateway-example/internal/service/mock"
	"github.com/stretchr/testify/assert"
)

func TestServer_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mockservice.NewMockProductService(ctrl)

	// Initialize the server with the mocked service
	s := server{productService: mockProductService}

	// Product to create
	productToCreate := &CreateProductBody{
		Name:        "Test Product",
		Price:       100,
		Description: "This is a test product",
	}

	// Expecting the service Create function to be called once with our product to create
	mockProductService.EXPECT().Create(gomock.Any()).Return(nil)

	_, err := s.CreateProduct(context.Background(), productToCreate)

	assert.NoError(t, err)
}

func TestServer_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mockservice.NewMockProductService(ctrl)

	s := server{productService: mockProductService}

	mockProductService.EXPECT().List().Return([]domain.Product{}, nil)

	_, err := s.GetProducts(context.Background(), &ProductParams{})

	assert.NoError(t, err)
}

func TestServer_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mockservice.NewMockProductService(ctrl)

	s := server{productService: mockProductService}

	mockProductService.EXPECT().Get(gomock.Any()).Return(&domain.Product{}, nil)

	_, err := s.GetProduct(context.Background(), &ProductId{})

	assert.NoError(t, err)
}

func TestServer_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mockservice.NewMockProductService(ctrl)

	s := server{productService: mockProductService}

	mockProductService.EXPECT().Delete(gomock.Any()).Return(nil)

	_, err := s.DeleteProduct(context.Background(), &ProductId{})

	assert.NoError(t, err)
}

func TestServer_GetProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mockservice.NewMockProductService(ctrl)

	s := server{productService: mockProductService}

	mockProductService.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error"))

	_, err := s.GetProduct(context.Background(), &ProductId{})

	assert.Error(t, err)
}

func TestServer_CreateProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mockservice.NewMockProductService(ctrl)
	s := server{productService: mockProductService}

	productToCreate := &CreateProductBody{
		Name:        "Test Product",
		Price:       100,
		Description: "This is a test product",
	}

	mockProductService.EXPECT().Create(gomock.Any()).Return(errors.New("create product error"))

	_, err := s.CreateProduct(context.Background(), productToCreate)

	assert.Error(t, err)
}

func TestServer_GetProducts_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mockservice.NewMockProductService(ctrl)
	s := server{productService: mockProductService}

	mockProductService.EXPECT().List().Return(nil, errors.New("get products error"))

	_, err := s.GetProducts(context.Background(), &ProductParams{})

	assert.Error(t, err)
}

func TestServer_DeleteProduct_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mockservice.NewMockProductService(ctrl)
	s := server{productService: mockProductService}

	mockProductService.EXPECT().Delete(gomock.Any()).Return(errors.New("delete product error"))

	_, err := s.DeleteProduct(context.Background(), &ProductId{})

	assert.Error(t, err)
}
