package main

import (
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/repository"
	"github.com/patrickkoss/grpc-gateway-example/internal/service"
)

func main() {
	productRepoLogger := logging.New(logging.String("component", "productRepository"))
	productServiceLogger := logging.New(logging.String("component", "productService"))
	productGrpcApiLogger := logging.New(logging.String("component", "productGrpcApi"))
	options := api.Options{
		Port:        "8001",
		MetricsPort: "9001",
	}

	productRepository := repository.NewProductRepository(productRepoLogger)
	productService := service.NewProductService(productServiceLogger, productRepository)
	productGrpcApi := api.NewProductGrpcApi(options, productService, productGrpcApiLogger)

	productGrpcApiLogger.Info("Starting productGrpcApi")

	err := productGrpcApi.Start()
	if err != nil {
		productGrpcApiLogger.Fatal("Failed to start productGrpcApi", logging.String("err", err.Error()))
	}
}
