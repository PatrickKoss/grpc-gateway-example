package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	mockrepository "github.com/patrickkoss/grpc-gateway-example/internal/adapter/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepository.NewMockUserRepository(ctrl)
	logger := logging.New()
	service := NewUserService(logger, mockRepo)

	user := &domain.User{
		Id:      "1",
		Name:    "testuser",
		Email:   "test@example.com",
		Surname: "testsurname",
	}

	mockRepo.EXPECT().Create(user).Return(nil)
	err := service.Create(user)
	assert.NoError(t, err)

	mockRepo.EXPECT().Get("1").Return(user, nil)
	storedUser, err := service.Get("1")
	assert.NoError(t, err)
	assert.NotNil(t, storedUser)
	assert.Equal(t, user, storedUser)

	mockRepo.EXPECT().List().Return([]domain.User{*user}, nil)
	users, err := service.List()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, user, &users[0])

	mockRepo.EXPECT().Delete("1").Return(nil)
	err = service.Delete("1")
	assert.NoError(t, err)

	mockRepo.EXPECT().Get("1").Return(nil, errors.New("not found"))
	storedUser, err = service.Get("1")
	assert.Error(t, err)
	assert.Nil(t, storedUser)

	mockRepo.EXPECT().List().Return(nil, errors.New("empty list"))
	users, err = service.List()
	assert.Error(t, err)
	assert.Nil(t, users)
}

func TestUserServiceErrorCasesWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepository.NewMockUserRepository(ctrl)
	logger := logging.New()
	service := NewUserService(logger, mockRepo)

	user := &domain.User{
		Id:      "1",
		Name:    "testuser",
		Email:   "test@example.com",
		Surname: "testsurname",
	}

	// Test Create Error
	mockRepo.EXPECT().Create(user).Return(errors.New("create error"))
	err := service.Create(user)
	assert.Error(t, err)

	// Test Get Error
	mockRepo.EXPECT().Get("non-existent-id").Return(nil, errors.New("not found"))
	storedUser, err := service.Get("non-existent-id")
	assert.Error(t, err)
	assert.Nil(t, storedUser)

	// Test List Error
	mockRepo.EXPECT().List().Return(nil, errors.New("list error"))
	users, err := service.List()
	assert.Error(t, err)
	assert.Nil(t, users)

	// Test Delete Error
	mockRepo.EXPECT().Delete("non-existent-id").Return(errors.New("not found"))
	err = service.Delete("non-existent-id")
	assert.Error(t, err)
}
