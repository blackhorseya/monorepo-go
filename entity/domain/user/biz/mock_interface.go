// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	model "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	contextx "github.com/blackhorseya/monorepo-go/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserBiz is a mock of IUserBiz interface.
type MockIUserBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIUserBizMockRecorder
}

// MockIUserBizMockRecorder is the mock recorder for MockIUserBiz.
type MockIUserBizMockRecorder struct {
	mock *MockIUserBiz
}

// NewMockIUserBiz creates a new mock instance.
func NewMockIUserBiz(ctrl *gomock.Controller) *MockIUserBiz {
	mock := &MockIUserBiz{ctrl: ctrl}
	mock.recorder = &MockIUserBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserBiz) EXPECT() *MockIUserBizMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockIUserBiz) Login(ctx contextx.Contextx, username, password string) (*model.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, username, password)
	ret0, _ := ret[0].(*model.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIUserBizMockRecorder) Login(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIUserBiz)(nil).Login), ctx, username, password)
}

// Logout mocks base method.
func (m *MockIUserBiz) Logout(ctx contextx.Contextx, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockIUserBizMockRecorder) Logout(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockIUserBiz)(nil).Logout), ctx, token)
}

// Signup mocks base method.
func (m *MockIUserBiz) Signup(ctx contextx.Contextx, username, password string) (*model.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", ctx, username, password)
	ret0, _ := ret[0].(*model.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockIUserBizMockRecorder) Signup(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockIUserBiz)(nil).Signup), ctx, username, password)
}
