// Code generated by MockGen. DO NOT EDIT.
// Source: internal/users/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "github.com/martinyonatann/go-unit-test/internal/users/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockRepositories is a mock of Repositories interface.
type MockRepositories struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoriesMockRecorder
}

// MockRepositoriesMockRecorder is the mock recorder for MockRepositories.
type MockRepositoriesMockRecorder struct {
	mock *MockRepositories
}

// NewMockRepositories creates a new mock instance.
func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
	mock := &MockRepositories{ctrl: ctrl}
	mock.recorder = &MockRepositoriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositories) EXPECT() *MockRepositoriesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepositories) Create(ctx context.Context, request entities.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoriesMockRecorder) Create(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositories)(nil).Create), ctx, request)
}

// Detail mocks base method.
func (m *MockRepositories) Detail(ctx context.Context, userUUID string) (entities.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Detail", ctx, userUUID)
	ret0, _ := ret[0].(entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detail indicates an expected call of Detail.
func (mr *MockRepositoriesMockRecorder) Detail(ctx, userUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detail", reflect.TypeOf((*MockRepositories)(nil).Detail), ctx, userUUID)
}
