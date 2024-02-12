// Code generated by MockGen. DO NOT EDIT.
// Source: stock.go

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	agg "github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	contextx "github.com/blackhorseya/monorepo-go/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockIStockRepo is a mock of IStockRepo interface.
type MockIStockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIStockRepoMockRecorder
}

// MockIStockRepoMockRecorder is the mock recorder for MockIStockRepo.
type MockIStockRepoMockRecorder struct {
	mock *MockIStockRepo
}

// NewMockIStockRepo creates a new mock instance.
func NewMockIStockRepo(ctrl *gomock.Controller) *MockIStockRepo {
	mock := &MockIStockRepo{ctrl: ctrl}
	mock.recorder = &MockIStockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStockRepo) EXPECT() *MockIStockRepoMockRecorder {
	return m.recorder
}

// BulkUpdateQuota mocks base method.
func (m *MockIStockRepo) BulkUpdateQuota(ctx contextx.Contextx, stocks []agg.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdateQuota", ctx, stocks)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdateQuota indicates an expected call of BulkUpdateQuota.
func (mr *MockIStockRepoMockRecorder) BulkUpdateQuota(ctx, stocks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdateQuota", reflect.TypeOf((*MockIStockRepo)(nil).BulkUpdateQuota), ctx, stocks)
}

// BulkUpsertInfo mocks base method.
func (m *MockIStockRepo) BulkUpsertInfo(ctx contextx.Contextx, stocks []agg.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpsertInfo", ctx, stocks)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpsertInfo indicates an expected call of BulkUpsertInfo.
func (mr *MockIStockRepoMockRecorder) BulkUpsertInfo(ctx, stocks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpsertInfo", reflect.TypeOf((*MockIStockRepo)(nil).BulkUpsertInfo), ctx, stocks)
}

// Get mocks base method.
func (m *MockIStockRepo) Get(ctx contextx.Contextx, symbol string) (agg.Stock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, symbol)
	ret0, _ := ret[0].(agg.Stock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIStockRepoMockRecorder) Get(ctx, symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIStockRepo)(nil).Get), ctx, symbol)
}

// List mocks base method.
func (m *MockIStockRepo) List(ctx contextx.Contextx) ([]agg.Stock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]agg.Stock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIStockRepoMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIStockRepo)(nil).List), ctx)
}