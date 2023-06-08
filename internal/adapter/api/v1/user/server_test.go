package user

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/patrickkoss/grpc-gateway-example/domain"
	mockservice "github.com/patrickkoss/grpc-gateway-example/internal/service/mock"
	"github.com/stretchr/testify/assert"
)

func TestServer_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	s := server{userService: mockUserService}

	userToCreate := &CreateUserBody{
		Name:    "Test User",
		Email:   "test@example.com",
		Surname: "Test",
	}

	mockUserService.EXPECT().Create(gomock.Any()).Return(nil)

	_, err := s.CreateUser(context.Background(), userToCreate)

	assert.NoError(t, err)
}

func TestServer_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	s := server{userService: mockUserService}

	mockUserService.EXPECT().List().Return([]domain.User{}, nil)

	_, err := s.GetUsers(context.Background(), &UserParams{})

	assert.NoError(t, err)
}

func TestServer_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	s := server{userService: mockUserService}

	mockUserService.EXPECT().Get(gomock.Any()).Return(&domain.User{}, nil)

	_, err := s.GetUser(context.Background(), &UserId{})

	assert.NoError(t, err)
}

func TestServer_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	s := server{userService: mockUserService}

	mockUserService.EXPECT().Delete(gomock.Any()).Return(nil)

	_, err := s.DeleteUser(context.Background(), &UserId{})

	assert.NoError(t, err)
}

// Testing Error Cases

func TestServer_CreateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	s := server{userService: mockUserService}

	userToCreate := &CreateUserBody{
		Name:    "Test User",
		Email:   "test@example.com",
		Surname: "Test",
	}

	mockUserService.EXPECT().Create(gomock.Any()).Return(errors.New("create user error"))

	_, err := s.CreateUser(context.Background(), userToCreate)

	assert.Error(t, err)
}

func TestServer_GetUsers_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	s := server{userService: mockUserService}

	mockUserService.EXPECT().List().Return(nil, errors.New("get users error"))

	_, err := s.GetUsers(context.Background(), &UserParams{})

	assert.Error(t, err)
}

func TestServer_GetUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	s := server{userService: mockUserService}

	mockUserService.EXPECT().Get(gomock.Any()).Return(nil, errors.New("get user error"))

	_, err := s.GetUser(context.Background(), &UserId{})

	assert.Error(t, err)
}

func TestServer_DeleteUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	s := server{userService: mockUserService}

	mockUserService.EXPECT().Delete(gomock.Any()).Return(errors.New("delete user error"))

	_, err := s.DeleteUser(context.Background(), &UserId{})

	assert.Error(t, err)
}
