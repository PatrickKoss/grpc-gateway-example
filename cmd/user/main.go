package main

import (
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/repository"
	"github.com/patrickkoss/grpc-gateway-example/internal/service"
)

func main() {
	userRepoLogger := logging.New(logging.String("component", "userRepository"))
	userServiceLogger := logging.New(logging.String("component", "userService"))
	userGrpcApiLogger := logging.New(logging.String("component", "userGrpcApi"))
	options := api.Options{
		Port:        "8002",
		MetricsPort: "9002",
	}

	userRepository := repository.NewUserRepository(userRepoLogger)
	userService := service.NewUserService(userServiceLogger, userRepository)
	userGrpcApi := api.NewUserGrpcApi(options, userService, userGrpcApiLogger)

	userGrpcApiLogger.Info("Starting userGrpcApi")

	err := userGrpcApi.Start()
	if err != nil {
		userGrpcApiLogger.Fatal("Failed to start userGrpcApi", logging.String("err", err.Error()))
	}
}
