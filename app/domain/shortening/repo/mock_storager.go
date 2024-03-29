// Code generated by MockGen. DO NOT EDIT.
// Source: storager.go

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	model "github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	contextx "github.com/blackhorseya/monorepo-go/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockStorager is a mock of Storager interface.
type MockStorager struct {
	ctrl     *gomock.Controller
	recorder *MockStoragerMockRecorder
}

// MockStoragerMockRecorder is the mock recorder for MockStorager.
type MockStoragerMockRecorder struct {
	mock *MockStorager
}

// NewMockStorager creates a new mock instance.
func NewMockStorager(ctrl *gomock.Controller) *MockStorager {
	mock := &MockStorager{ctrl: ctrl}
	mock.recorder = &MockStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorager) EXPECT() *MockStoragerMockRecorder {
	return m.recorder
}

// CreateURLRecord mocks base method.
func (m *MockStorager) CreateURLRecord(ctx contextx.Contextx, record *model.ShortenedUrl) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateURLRecord", ctx, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateURLRecord indicates an expected call of CreateURLRecord.
func (mr *MockStoragerMockRecorder) CreateURLRecord(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateURLRecord", reflect.TypeOf((*MockStorager)(nil).CreateURLRecord), ctx, record)
}

// GetURLRecordByShortURL mocks base method.
func (m *MockStorager) GetURLRecordByShortURL(ctx contextx.Contextx, shortURL string) (*model.ShortenedUrl, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURLRecordByShortURL", ctx, shortURL)
	ret0, _ := ret[0].(*model.ShortenedUrl)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURLRecordByShortURL indicates an expected call of GetURLRecordByShortURL.
func (mr *MockStoragerMockRecorder) GetURLRecordByShortURL(ctx, shortURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURLRecordByShortURL", reflect.TypeOf((*MockStorager)(nil).GetURLRecordByShortURL), ctx, shortURL)
}
