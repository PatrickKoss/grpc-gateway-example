package repository

import (
	"testing"

	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	logger := logging.New()
	repo := NewUserRepository(logger)

	user := &domain.User{
		Id:    "1",
		Name:  "testuser",
		Email: "test@example.com",
	}

	// Test Create
	err := repo.Create(user)
	assert.NoError(t, err)

	// Test Get
	storedUser, err := repo.Get("1")
	assert.NoError(t, err)
	assert.NotNil(t, storedUser)
	assert.Equal(t, user, storedUser)

	// Test List
	users, err := repo.List()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, user, &users[0])

	// Test Delete
	err = repo.Delete("1")
	assert.NoError(t, err)

	// Test Get after Delete
	storedUser, err = repo.Get("1")
	assert.Error(t, err)
	assert.Nil(t, storedUser)

	// Test List after Delete
	users, err = repo.List()
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestUserRepositoryErrorCases(t *testing.T) {
	logger := logging.New()
	repo := NewUserRepository(logger)

	user := &domain.User{
		Id:    "1",
		Name:  "testuser",
		Email: "test@example.com",
	}

	// Test Get with non-existent user
	storedUser, err := repo.Get("non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNotFound, err)
	assert.Nil(t, storedUser)

	// Test Create
	err = repo.Create(user)
	assert.NoError(t, err)

	// Test creating a duplicate user
	duplicateUser := &domain.User{
		Id:    "1",
		Name:  "duplicateuser",
		Email: "duplicate@example.com",
	}
	err = repo.Create(duplicateUser)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrAlreadyExists, err)

	// Test Delete with non-existent user
	err = repo.Delete("non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNotFound, err)
}
