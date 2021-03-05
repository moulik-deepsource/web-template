// Code generated by MockGen. DO NOT EDIT.
// Source: dbmodels/make/querier.go

// Package mock_make is a generated GoMock package.
package mock_make

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	make "github.com/rickypai/web-template/api/dbmodels/make"
	reflect "reflect"
)

// MockQuerier is a mock of Querier interface
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CountTotal mocks base method
func (m *MockQuerier) CountTotal(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountTotal", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountTotal indicates an expected call of CountTotal
func (mr *MockQuerierMockRecorder) CountTotal(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountTotal", reflect.TypeOf((*MockQuerier)(nil).CountTotal), ctx)
}

// CreateOne mocks base method
func (m *MockQuerier) CreateOne(ctx context.Context, arg make.CreateOneParams) (make.Make, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOne", ctx, arg)
	ret0, _ := ret[0].(make.Make)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOne indicates an expected call of CreateOne
func (mr *MockQuerierMockRecorder) CreateOne(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOne", reflect.TypeOf((*MockQuerier)(nil).CreateOne), ctx, arg)
}

// GetByID mocks base method
func (m *MockQuerier) GetByID(ctx context.Context, id int64) (make.Make, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(make.Make)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockQuerierMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockQuerier)(nil).GetByID), ctx, id)
}

// GetManyByIDs mocks base method
func (m *MockQuerier) GetManyByIDs(ctx context.Context, dollar_1 []int64) ([]make.Make, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManyByIDs", ctx, dollar_1)
	ret0, _ := ret[0].([]make.Make)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManyByIDs indicates an expected call of GetManyByIDs
func (mr *MockQuerierMockRecorder) GetManyByIDs(ctx, dollar_1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManyByIDs", reflect.TypeOf((*MockQuerier)(nil).GetManyByIDs), ctx, dollar_1)
}

// ListOffset mocks base method
func (m *MockQuerier) ListOffset(ctx context.Context, arg make.ListOffsetParams) ([]make.Make, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOffset", ctx, arg)
	ret0, _ := ret[0].([]make.Make)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOffset indicates an expected call of ListOffset
func (mr *MockQuerierMockRecorder) ListOffset(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOffset", reflect.TypeOf((*MockQuerier)(nil).ListOffset), ctx, arg)
}
