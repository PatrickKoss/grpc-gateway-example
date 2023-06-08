package main

import (
	"context"

	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/v1/product"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/v1/user"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := logging.New()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	params := api.RegistrationParams{
		Context:  context.Background(),
		Opts:     opts,
		Endpoint: "localhost:8001",
	}

	apiBuilder, err := api.NewRestApiBuilder().WithRegistration(product.RegisterProductServiceHandlerFromEndpoint, params)
	if err != nil {
		logger.Fatal("Failed to create rest restApi", logging.String("err", err.Error()))
	}

	params = api.RegistrationParams{
		Context:  context.Background(),
		Opts:     opts,
		Endpoint: "localhost:8002",
	}
	apiBuilder, err = apiBuilder.WithRegistration(user.RegisterUserServiceHandlerFromEndpoint, params)
	if err != nil {
		logger.Fatal("Failed to create rest restApi", logging.String("err", err.Error()))
	}

	restApi := apiBuilder.Build()

	logger.Info("Starting rest api")

	err = restApi.Start("8000", "9000")
	if err != nil {
		logger.Fatal("Failed to start rest restApi", logging.String("err", err.Error()))
	}
}
