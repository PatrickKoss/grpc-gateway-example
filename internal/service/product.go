//nolint:dupl // it's intended to be similar to other services
package service

import (
	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/repository"
	"github.com/patrickkoss/grpc-gateway-example/internal/service/validator"
)

type ProductService interface {
	Get(id string) (*domain.Product, error)
	List() ([]domain.Product, error)
	Create(user *domain.Product) error
	Delete(id string) error
}

type productService struct {
	logger            logging.Logger
	productRepository repository.ProductRepository
	validator         validator.Validator
}

func (p productService) Get(id string) (*domain.Product, error) {
	return p.productRepository.Get(id)
}

func (p productService) List() ([]domain.Product, error) {
	return p.productRepository.List()
}

func (p productService) Create(product *domain.Product) error {
	if err := p.validator.ValidateStruct(product); err != nil {
		return &domain.InvalidInputError{Message: err.Error()}
	}

	return p.productRepository.Create(product)
}

func (p productService) Delete(id string) error {
	return p.productRepository.Delete(id)
}

func NewProductService(logger logging.Logger, productRepository repository.ProductRepository) ProductService {
	return productService{
		logger:            logger,
		productRepository: productRepository,
		validator:         validator.NewValidator(),
	}
}
