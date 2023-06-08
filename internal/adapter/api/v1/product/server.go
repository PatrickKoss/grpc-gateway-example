package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/common"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/domain/proto"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/patrickkoss/grpc-gateway-example/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	productService service.ProductService
	logger         logging.Logger
}

func (s server) CreateProduct(ctx context.Context, body *CreateProductBody) (*proto.Message, error) {
	domainProduct := domain.Product{
		Id:          uuid.NewString(),
		Description: body.Description,
		Price:       body.Price,
		Name:        body.Name,
	}

	err := s.productService.Create(&domainProduct)
	if err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

	return &proto.Message{Message: "Product created"}, nil
}

func (s server) GetProducts(ctx context.Context, params *ProductParams) (*ListProductsResponse, error) {
	domainProducts, err := s.productService.List()
	if err != nil {
		return nil, common.GetStatusErrorFromDomainError(err)
	}

	products := make([]*Product, len(domainProducts))
	for i, domainProduct := range domainProducts {
		products[i] = &Product{
			Id:          domainProduct.Id,
			Description: domainProduct.Description,
			Price:       domainProduct.Price,
			Name:        domainProduct.Name,
		}
	}

	return &ListProductsResponse{Products: products}, nil
}

func (s server) GetProduct(ctx context.Context, id *ProductId) (*GetProductResponse, error) {
	domainProduct, err := s.productService.Get(id.Id)
	if err != nil {
		return nil, common.GetStatusErrorFromDomainError(err)
	}

	product := Product{
		Id:          domainProduct.Id,
		Description: domainProduct.Description,
		Price:       domainProduct.Price,
		Name:        domainProduct.Name,
	}

	return &GetProductResponse{
		Product: &product,
	}, nil
}

func (s server) DeleteProduct(ctx context.Context, id *ProductId) (*proto.Message, error) {
	err := s.productService.Delete(id.Id)
	if err != nil {
		return nil, common.GetStatusErrorFromDomainError(err)
	}

	return &proto.Message{Message: "Product deleted"}, nil
}

func (s server) mustEmbedUnimplementedProductServiceServer() {}

func NewServer(productService service.ProductService, logger logging.Logger) ProductServiceServer {
	return &server{productService: productService, logger: logger}
}
