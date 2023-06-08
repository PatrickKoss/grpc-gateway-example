//nolint:dupl // it's intended to be similar to other services
package service

import (
	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/repository"
	"github.com/patrickkoss/grpc-gateway-example/internal/service/validator"
)

type UserService interface {
	Get(id string) (*domain.User, error)
	List() ([]domain.User, error)
	Create(user *domain.User) error
	Delete(id string) error
}

type userService struct {
	logger         logging.Logger
	userRepository repository.UserRepository
	validator      validator.Validator
}

func (u userService) Get(id string) (*domain.User, error) {
	return u.userRepository.Get(id)
}

func (u userService) List() ([]domain.User, error) {
	return u.userRepository.List()
}

func (u userService) Create(user *domain.User) error {
	if err := u.validator.ValidateStruct(user); err != nil {
		return &domain.InvalidInputError{Message: err.Error()}
	}

	return u.userRepository.Create(user)
}

func (u userService) Delete(id string) error {
	return u.userRepository.Delete(id)
}

func NewUserService(logger logging.Logger, userRepository repository.UserRepository) UserService {
	return userService{
		logger:         logger,
		userRepository: userRepository,
		validator:      validator.NewValidator(),
	}
}
