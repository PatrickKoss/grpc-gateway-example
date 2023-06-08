package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/middleware"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/metrics_collector"
	"google.golang.org/grpc"
)

type RestApiBuilder interface {
	WithRegistration(registration func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error, params RegistrationParams) (RestApiBuilder, error)
	Build() RestApi
}

type RestApi interface {
	Start(port string, metricsPort string) error
}

type restApiBuilder struct {
	mux *runtime.ServeMux
}

type restApi struct {
	handler http.Handler
	logger  logging.Logger
}

type RegistrationParams struct {
	Context  context.Context
	Opts     []grpc.DialOption
	Endpoint string
}

func (r restApiBuilder) WithRegistration(registration func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error, params RegistrationParams) (RestApiBuilder, error) {
	err := registration(params.Context, r.mux, params.Endpoint, params.Opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r restApiBuilder) Build() RestApi {
	logger := logging.New()
	httpMetricsCollector := metrics_collector.NewHttpApiMetrics()

	// Wrap the mux with the middlewares
	loggingMiddleware := middleware.LoggingMiddleware(logger)
	metricsMiddleware := middleware.MetricsMiddleware(httpMetricsCollector, logger)
	securityMiddleware := middleware.SecurityMiddleware()

	handler := loggingMiddleware(r.mux)
	handler = metricsMiddleware(handler)
	handler = securityMiddleware(handler)

	return restApi{handler: handler, logger: logging.New()}
}

func (r restApi) Start(port string, metricsPort string) error {
	go func() {
		err := startMetricsServer(metricsPort)
		if err != nil {
			r.logger.Fatal(err.Error())
		}
	}()

	return http.ListenAndServe(fmt.Sprintf(":%s", port), r.handler)
}

func NewRestApiBuilder() RestApiBuilder {
	return restApiBuilder{mux: runtime.NewServeMux(runtime.WithErrorHandler(middleware.CustomHTTPError), runtime.WithForwardResponseOption(middleware.HttpResponseModifier))}
}
