package api

import (
	"fmt"
	"net"

	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/v1/product"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/v1/user"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/patrickkoss/grpc-gateway-example/internal/service"
	"google.golang.org/grpc"
)

type GrpcApi interface {
	Start() error
}

type Options struct {
	Port        string
	MetricsPort string
}

type grpcApi struct {
	server  *grpc.Server
	options Options
	logger  logging.Logger
}

type productGrpcApi struct {
	grpcApi
}

type userGrpcApi struct {
	grpcApi
}

func (g grpcApi) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", g.options.Port))
	if err != nil {
		return err
	}

	go func() {
		err := startMetricsServer(g.options.MetricsPort)
		if err != nil {
			g.logger.Fatal(err.Error())
		}
	}()

	return g.server.Serve(lis)
}

func NewProductGrpcApi(options Options, productService service.ProductService, logger logging.Logger) GrpcApi {
	grpcServer := grpc.NewServer()
	productServer := product.NewServer(productService, logger)
	product.RegisterProductServiceServer(grpcServer, productServer)

	return productGrpcApi{
		grpcApi{
			server:  grpcServer,
			options: options,
		},
	}
}

func NewUserGrpcApi(options Options, userService service.UserService, logger logging.Logger) GrpcApi {
	grpcServer := grpc.NewServer()
	userServer := user.NewServer(userService, logger)
	user.RegisterUserServiceServer(grpcServer, userServer)

	return userGrpcApi{
		grpcApi{
			server:  grpcServer,
			options: options,
			logger:  logging.New(),
		},
	}
}
