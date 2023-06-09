// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/user.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/patrickkoss/grpc-gateway-example/domain"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserService) Create(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserServiceMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserService)(nil).Create), user)
}

// Delete mocks base method.
func (m *MockUserService) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserService)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockUserService) Get(id string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserService)(nil).Get), id)
}

// List mocks base method.
func (m *MockUserService) List() ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserService)(nil).List))
}
