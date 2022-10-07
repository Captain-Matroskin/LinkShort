// Code generated by MockGen. DO NOT EDIT.
// Source: LinkShortening/internals/myerror (interfaces: MultiLoggerInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMultiLoggerInterface is a mock of MultiLoggerInterface interface.
type MockMultiLoggerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMultiLoggerInterfaceMockRecorder
}

// MockMultiLoggerInterfaceMockRecorder is the mock recorder for MockMultiLoggerInterface.
type MockMultiLoggerInterfaceMockRecorder struct {
	mock *MockMultiLoggerInterface
}

// NewMockMultiLoggerInterface creates a new mock instance.
func NewMockMultiLoggerInterface(ctrl *gomock.Controller) *MockMultiLoggerInterface {
	mock := &MockMultiLoggerInterface{ctrl: ctrl}
	mock.recorder = &MockMultiLoggerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMultiLoggerInterface) EXPECT() *MockMultiLoggerInterfaceMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockMultiLoggerInterface) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockMultiLoggerInterfaceMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockMultiLoggerInterface)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockMultiLoggerInterface) Errorf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockMultiLoggerInterfaceMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockMultiLoggerInterface)(nil).Errorf), varargs...)
}

// Infof mocks base method.
func (m *MockMultiLoggerInterface) Infof(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockMultiLoggerInterfaceMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockMultiLoggerInterface)(nil).Infof), varargs...)
}

// Sync mocks base method.
func (m *MockMultiLoggerInterface) Sync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockMultiLoggerInterfaceMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockMultiLoggerInterface)(nil).Sync))
}

// Warnf mocks base method.
func (m *MockMultiLoggerInterface) Warnf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockMultiLoggerInterfaceMockRecorder) Warnf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockMultiLoggerInterface)(nil).Warnf), varargs...)
}
