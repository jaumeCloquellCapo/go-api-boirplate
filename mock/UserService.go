// Code generated by MockGen. DO NOT EDIT.
// Source: UserService.go

// Package mock is a generated GoMock package.
package mock

import (
	"ApiRest/app/model"
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockAuthServiceUseCase is a mock of UserUseCase interface
type MockServiceUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockAuthUseCaseMockRecorder is the mock recorder for MockAuthServiceUseCase
type MockUserUseCaseMockRecorder struct {
	mock *MockServiceUseCase
}

// NewMockUserServiceCase creates a new mock instance
func NewMockUserServiceCase(ctrl *gomock.Controller) *MockServiceUseCase {
	mock := &MockServiceUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockServiceUseCase) FindByID(id int) (user *model.User, err error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockUserUseCaseMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockServiceUseCase)(nil).FindByID), id)
}

// Register mocks base method
func (m *MockServiceUseCase) RemoveByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUserUseCaseMockRecorder) RemoveById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByID", reflect.TypeOf((*MockServiceUseCase)(nil).RemoveByID), id)
}

// Register mocks base method
func (m *MockServiceUseCase) UpdateByID(id int, user model.UpdateUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", id, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUserUseCaseMockRecorder) UpdateById(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockServiceUseCase)(nil).UpdateByID), id, user)
}

// Register mocks base method
func (m *MockServiceUseCase) Store(user model.CreateUser) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockUserUseCaseMockRecorder) Store(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockServiceUseCase)(nil).Store), user)
}
