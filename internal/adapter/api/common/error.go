package common

import (
	"errors"

	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/domain/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetStatusErrorFromDomainError(err error) error {
	if errors.Is(err, domain.ErrNotFound) {
		detailsError := &proto.ErrorMessage{
			Error:   err.Error(),
			Message: "Resource not found",
		}

		s, err2 := status.New(codes.Code(404), "Resource not found").WithDetails(detailsError)
		if err2 != nil {
			return err
		}

		return s.Err()
	}
	if errors.Is(err, domain.ErrAlreadyExists) {
		detailsError := &proto.ErrorMessage{
			Error:   err.Error(),
			Message: "Resource already exists",
		}

		s, err2 := status.New(codes.Code(409), "Resource already exists").WithDetails(detailsError)
		if err2 != nil {
			return err
		}

		return s.Err()
	}

	if errors.Is(err, &domain.InvalidInputError{}) {
		detailsError := &proto.ErrorMessage{
			Error:   err.Error(),
			Message: "Invalid input",
		}

		s, err2 := status.New(codes.Code(400), "Invalid input").WithDetails(detailsError)
		if err2 != nil {
			return err
		}

		return s.Err()
	}

	return err
}
