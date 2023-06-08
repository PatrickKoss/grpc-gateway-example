//nolint:dupl // it's intended to be similar to other repositories
package repository

import (
	"sync"

	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
)

type UserRepository interface {
	Get(id string) (*domain.User, error)
	List() ([]domain.User, error)
	Create(user *domain.User) error
	Delete(id string) error
}

type userRepository struct {
	logger logging.Logger
	rwLock sync.RWMutex
	users  map[string]domain.User
}

func (u *userRepository) Get(id string) (*domain.User, error) {
	u.rwLock.RLock()
	defer u.rwLock.RUnlock()

	user, ok := u.users[id]
	if !ok {
		u.logger.Info("user not found", logging.String("id", id))

		return nil, domain.ErrNotFound
	}

	return &user, nil
}

func (u *userRepository) List() ([]domain.User, error) {
	u.rwLock.RLock()
	defer u.rwLock.RUnlock()

	users := make([]domain.User, 0, len(u.users))
	for _, user := range u.users {
		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) Create(user *domain.User) error {
	u.rwLock.Lock()
	defer u.rwLock.Unlock()

	_, ok := u.users[user.Id]
	if ok {
		u.logger.Info("user already exists", logging.String("id", user.Id))

		return domain.ErrAlreadyExists
	}

	u.users[user.Id] = *user

	return nil
}

func (u *userRepository) Delete(id string) error {
	u.rwLock.Lock()
	defer u.rwLock.Unlock()

	_, ok := u.users[id]
	if !ok {
		u.logger.Info("user not found", logging.String("id", id))

		return domain.ErrNotFound
	}

	delete(u.users, id)

	return nil
}

func NewUserRepository(logger logging.Logger) UserRepository {
	return &userRepository{
		logger: logger,
		rwLock: sync.RWMutex{},
		users:  make(map[string]domain.User),
	}
}
