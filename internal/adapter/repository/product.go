//nolint:dupl // it's intended to be similar to other repositories
package repository

import (
	"sync"

	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
)

type ProductRepository interface {
	Get(id string) (*domain.Product, error)
	List() ([]domain.Product, error)
	Create(user *domain.Product) error
	Delete(id string) error
}

type productRepository struct {
	logger   logging.Logger
	rwLock   sync.RWMutex
	products map[string]domain.Product
}

func (p *productRepository) Get(id string) (*domain.Product, error) {
	p.rwLock.RLock()
	defer p.rwLock.RUnlock()

	product, ok := p.products[id]
	if !ok {
		p.logger.Info("product not found", logging.String("id", id))

		return nil, domain.ErrNotFound
	}

	return &product, nil
}

func (p *productRepository) List() ([]domain.Product, error) {
	p.rwLock.RLock()
	defer p.rwLock.RUnlock()

	products := make([]domain.Product, 0, len(p.products))
	for _, product := range p.products {
		products = append(products, product)
	}

	return products, nil
}

func (p *productRepository) Create(product *domain.Product) error {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()

	_, ok := p.products[product.Id]
	if ok {
		p.logger.Info("product already exists", logging.String("id", product.Id))

		return domain.ErrAlreadyExists
	}

	p.products[product.Id] = *product

	return nil
}

func (p *productRepository) Delete(id string) error {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()

	_, ok := p.products[id]
	if !ok {
		p.logger.Info("product not found", logging.String("id", id))

		return domain.ErrNotFound
	}

	delete(p.products, id)

	return nil
}

func NewProductRepository(logger logging.Logger) ProductRepository {
	return &productRepository{
		logger:   logger,
		rwLock:   sync.RWMutex{},
		products: make(map[string]domain.Product),
	}
}
