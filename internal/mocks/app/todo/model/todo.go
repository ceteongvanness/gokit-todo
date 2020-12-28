// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cage1016/todo/internal/app/todo/model (interfaces: TodoRepository)

// Package automocks is a generated GoMock package.
package automocks

import (
	context "context"
	model "github.com/cage1016/todo/internal/app/todo/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTodoRepository is a mock of TodoRepository interface
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockTodoRepository) Add(arg0 context.Context, arg1 *model.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockTodoRepositoryMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTodoRepository)(nil).Add), arg0, arg1)
}

// Clear mocks base method
func (m *MockTodoRepository) Clear(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear
func (mr *MockTodoRepositoryMockRecorder) Clear(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockTodoRepository)(nil).Clear), arg0)
}

// Complete mocks base method
func (m *MockTodoRepository) Complete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Complete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Complete indicates an expected call of Complete
func (mr *MockTodoRepositoryMockRecorder) Complete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockTodoRepository)(nil).Complete), arg0, arg1)
}

// CompleteAll mocks base method
func (m *MockTodoRepository) CompleteAll(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteAll indicates an expected call of CompleteAll
func (mr *MockTodoRepositoryMockRecorder) CompleteAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteAll", reflect.TypeOf((*MockTodoRepository)(nil).CompleteAll), arg0)
}

// Delete mocks base method
func (m *MockTodoRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTodoRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoRepository)(nil).Delete), arg0, arg1)
}

// List mocks base method
func (m *MockTodoRepository) List(arg0 context.Context, arg1 string) ([]*model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockTodoRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTodoRepository)(nil).List), arg0, arg1)
}

// Update mocks base method
func (m *MockTodoRepository) Update(arg0 context.Context, arg1 *model.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockTodoRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoRepository)(nil).Update), arg0, arg1)
}
