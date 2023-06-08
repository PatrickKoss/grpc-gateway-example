package middleware

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/domain/proto"
	"google.golang.org/grpc/status"
)

func CustomHTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	// Convert the error to a customError
	s := status.Convert(err)

	// 2xx requests are fine
	if s.Code() <= 300 {
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)

		return
	}

	if 200 > s.Code() && s.Code() > 505 {
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)

		return
	}

	errorMessage := proto.ErrorMessage{
		Message: s.Message(),
		Error:   s.Message(),
	}
	for _, d := range s.Details() {
		errorResponse, ok := d.(*proto.ErrorMessage)
		if ok {
			errorMessage.Error = errorResponse.Error
		}
	}

	w.WriteHeader(int(s.Code()))
	w.Header().Set("Content-Type", "application/json")

	// Write the error details to the response body.
	body, _ := marshaler.Marshal(proto.ErrorMessage{
		Message: errorMessage.Message,
		Error:   errorMessage.Error,
	})
	_, _ = w.Write(body)
}
