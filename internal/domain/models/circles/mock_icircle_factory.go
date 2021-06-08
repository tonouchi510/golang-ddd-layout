// Code generated by MockGen. DO NOT EDIT.
// Source: icircle_factory.go

// Package circles is a generated GoMock package.
package circles

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	users "github.com/tonouchi510/golang-ddd-layout/internal/domain/models/users"
)

// MockICircleFactory is a mock of ICircleFactory interface.
type MockICircleFactory struct {
	ctrl     *gomock.Controller
	recorder *MockICircleFactoryMockRecorder
}

// MockICircleFactoryMockRecorder is the mock recorder for MockICircleFactory.
type MockICircleFactoryMockRecorder struct {
	mock *MockICircleFactory
}

// NewMockICircleFactory creates a new mock instance.
func NewMockICircleFactory(ctrl *gomock.Controller) *MockICircleFactory {
	mock := &MockICircleFactory{ctrl: ctrl}
	mock.recorder = &MockICircleFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICircleFactory) EXPECT() *MockICircleFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICircleFactory) Create(name CircleName, ownerId users.UserId) (*Circle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, ownerId)
	ret0, _ := ret[0].(*Circle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockICircleFactoryMockRecorder) Create(name, ownerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICircleFactory)(nil).Create), name, ownerId)
}
