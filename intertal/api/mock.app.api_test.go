// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package api_test is a generated GoMock package.
package api_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// GetNumber mocks base method.
func (m *MockApplication) GetNumber() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumber")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetNumber indicates an expected call of GetNumber.
func (mr *MockApplicationMockRecorder) GetNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumber", reflect.TypeOf((*MockApplication)(nil).GetNumber))
}

// PutNumber mocks base method.
func (m *MockApplication) PutNumber(num int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutNumber", num)
	ret0, _ := ret[0].(int64)
	return ret0
}

// PutNumber indicates an expected call of PutNumber.
func (mr *MockApplicationMockRecorder) PutNumber(num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutNumber", reflect.TypeOf((*MockApplication)(nil).PutNumber), num)
}
