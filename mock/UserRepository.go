// Code generated by MockGen. DO NOT EDIT.
// Source: UserRepository.go

// Package mock is a generated GoMock package.
package mock

import (
	"ApiRest/app/model"
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockUserPGRepository is a mock of UserPGRepository interface
type MockUserPGRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserPGRepositoryMockRecorder
}

// MockUserPGRepositoryMockRecorder is the mock recorder for MockUserPGRepository
type MockUserPGRepositoryMockRecorder struct {
	mock *MockUserPGRepository
}

// NewMockUserPGRepository creates a new mock instance
func NewMockUserPGRepository(ctrl *gomock.Controller) *MockUserPGRepository {
	mock := &MockUserPGRepository{ctrl: ctrl}
	mock.recorder = &MockUserPGRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserPGRepository) EXPECT() *MockUserPGRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserPGRepository) Create(cuser model.CreateUser) (user *model.User, err error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cuser)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserPGRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserPGRepository)(nil).Create), user)
}

// FindById mocks base method
func (m *MockUserPGRepository) FindById(id int) (user *model.User, err error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockUserPGRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserPGRepository)(nil).FindById), id)
}

func (m *MockUserPGRepository) RemoveById(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserPGRepositoryMockRecorder) RemoveById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserPGRepository)(nil).RemoveById), id)
}

func (m *MockUserPGRepository) UpdateById(id int, user model.UpdateUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserPGRepositoryMockRecorder) UpdateById(id interface{}, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockUserPGRepository)(nil).UpdateById), id, user)
}
