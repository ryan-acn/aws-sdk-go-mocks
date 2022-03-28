// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/marketplacecatalog/marketplacecatalogiface (interfaces: MarketplaceCatalogAPI)

// Package marketplacecatalogmock is a generated GoMock package.
package marketplacecatalogmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	marketplacecatalog "github.com/aws/aws-sdk-go/service/marketplacecatalog"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMarketplaceCatalogAPI is a mock of MarketplaceCatalogAPI interface
type MockMarketplaceCatalogAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMarketplaceCatalogAPIMockRecorder
}

// MockMarketplaceCatalogAPIMockRecorder is the mock recorder for MockMarketplaceCatalogAPI
type MockMarketplaceCatalogAPIMockRecorder struct {
	mock *MockMarketplaceCatalogAPI
}

// NewMockMarketplaceCatalogAPI creates a new mock instance
func NewMockMarketplaceCatalogAPI(ctrl *gomock.Controller) *MockMarketplaceCatalogAPI {
	mock := &MockMarketplaceCatalogAPI{ctrl: ctrl}
	mock.recorder = &MockMarketplaceCatalogAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarketplaceCatalogAPI) EXPECT() *MockMarketplaceCatalogAPIMockRecorder {
	return m.recorder
}

// CancelChangeSet mocks base method
func (m *MockMarketplaceCatalogAPI) CancelChangeSet(arg0 *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelChangeSet", arg0)
	ret0, _ := ret[0].(*marketplacecatalog.CancelChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelChangeSet indicates an expected call of CancelChangeSet
func (mr *MockMarketplaceCatalogAPIMockRecorder) CancelChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelChangeSet", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).CancelChangeSet), arg0)
}

// CancelChangeSetRequest mocks base method
func (m *MockMarketplaceCatalogAPI) CancelChangeSetRequest(arg0 *marketplacecatalog.CancelChangeSetInput) (*request.Request, *marketplacecatalog.CancelChangeSetOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelChangeSetRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*marketplacecatalog.CancelChangeSetOutput)
	return ret0, ret1
}

// CancelChangeSetRequest indicates an expected call of CancelChangeSetRequest
func (mr *MockMarketplaceCatalogAPIMockRecorder) CancelChangeSetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelChangeSetRequest", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).CancelChangeSetRequest), arg0)
}

// CancelChangeSetWithContext mocks base method
func (m *MockMarketplaceCatalogAPI) CancelChangeSetWithContext(arg0 context.Context, arg1 *marketplacecatalog.CancelChangeSetInput, arg2 ...request.Option) (*marketplacecatalog.CancelChangeSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelChangeSetWithContext", varargs...)
	ret0, _ := ret[0].(*marketplacecatalog.CancelChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelChangeSetWithContext indicates an expected call of CancelChangeSetWithContext
func (mr *MockMarketplaceCatalogAPIMockRecorder) CancelChangeSetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelChangeSetWithContext", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).CancelChangeSetWithContext), varargs...)
}

// DescribeChangeSet mocks base method
func (m *MockMarketplaceCatalogAPI) DescribeChangeSet(arg0 *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeChangeSet", arg0)
	ret0, _ := ret[0].(*marketplacecatalog.DescribeChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeChangeSet indicates an expected call of DescribeChangeSet
func (mr *MockMarketplaceCatalogAPIMockRecorder) DescribeChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeChangeSet", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).DescribeChangeSet), arg0)
}

// DescribeChangeSetRequest mocks base method
func (m *MockMarketplaceCatalogAPI) DescribeChangeSetRequest(arg0 *marketplacecatalog.DescribeChangeSetInput) (*request.Request, *marketplacecatalog.DescribeChangeSetOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeChangeSetRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*marketplacecatalog.DescribeChangeSetOutput)
	return ret0, ret1
}

// DescribeChangeSetRequest indicates an expected call of DescribeChangeSetRequest
func (mr *MockMarketplaceCatalogAPIMockRecorder) DescribeChangeSetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeChangeSetRequest", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).DescribeChangeSetRequest), arg0)
}

// DescribeChangeSetWithContext mocks base method
func (m *MockMarketplaceCatalogAPI) DescribeChangeSetWithContext(arg0 context.Context, arg1 *marketplacecatalog.DescribeChangeSetInput, arg2 ...request.Option) (*marketplacecatalog.DescribeChangeSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeChangeSetWithContext", varargs...)
	ret0, _ := ret[0].(*marketplacecatalog.DescribeChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeChangeSetWithContext indicates an expected call of DescribeChangeSetWithContext
func (mr *MockMarketplaceCatalogAPIMockRecorder) DescribeChangeSetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeChangeSetWithContext", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).DescribeChangeSetWithContext), varargs...)
}

// DescribeEntity mocks base method
func (m *MockMarketplaceCatalogAPI) DescribeEntity(arg0 *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEntity", arg0)
	ret0, _ := ret[0].(*marketplacecatalog.DescribeEntityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntity indicates an expected call of DescribeEntity
func (mr *MockMarketplaceCatalogAPIMockRecorder) DescribeEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntity", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).DescribeEntity), arg0)
}

// DescribeEntityRequest mocks base method
func (m *MockMarketplaceCatalogAPI) DescribeEntityRequest(arg0 *marketplacecatalog.DescribeEntityInput) (*request.Request, *marketplacecatalog.DescribeEntityOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEntityRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*marketplacecatalog.DescribeEntityOutput)
	return ret0, ret1
}

// DescribeEntityRequest indicates an expected call of DescribeEntityRequest
func (mr *MockMarketplaceCatalogAPIMockRecorder) DescribeEntityRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntityRequest", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).DescribeEntityRequest), arg0)
}

// DescribeEntityWithContext mocks base method
func (m *MockMarketplaceCatalogAPI) DescribeEntityWithContext(arg0 context.Context, arg1 *marketplacecatalog.DescribeEntityInput, arg2 ...request.Option) (*marketplacecatalog.DescribeEntityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEntityWithContext", varargs...)
	ret0, _ := ret[0].(*marketplacecatalog.DescribeEntityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntityWithContext indicates an expected call of DescribeEntityWithContext
func (mr *MockMarketplaceCatalogAPIMockRecorder) DescribeEntityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntityWithContext", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).DescribeEntityWithContext), varargs...)
}

// ListChangeSets mocks base method
func (m *MockMarketplaceCatalogAPI) ListChangeSets(arg0 *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChangeSets", arg0)
	ret0, _ := ret[0].(*marketplacecatalog.ListChangeSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChangeSets indicates an expected call of ListChangeSets
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListChangeSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChangeSets", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListChangeSets), arg0)
}

// ListChangeSetsPages mocks base method
func (m *MockMarketplaceCatalogAPI) ListChangeSetsPages(arg0 *marketplacecatalog.ListChangeSetsInput, arg1 func(*marketplacecatalog.ListChangeSetsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChangeSetsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListChangeSetsPages indicates an expected call of ListChangeSetsPages
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListChangeSetsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChangeSetsPages", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListChangeSetsPages), arg0, arg1)
}

// ListChangeSetsPagesWithContext mocks base method
func (m *MockMarketplaceCatalogAPI) ListChangeSetsPagesWithContext(arg0 context.Context, arg1 *marketplacecatalog.ListChangeSetsInput, arg2 func(*marketplacecatalog.ListChangeSetsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListChangeSetsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListChangeSetsPagesWithContext indicates an expected call of ListChangeSetsPagesWithContext
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListChangeSetsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChangeSetsPagesWithContext", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListChangeSetsPagesWithContext), varargs...)
}

// ListChangeSetsRequest mocks base method
func (m *MockMarketplaceCatalogAPI) ListChangeSetsRequest(arg0 *marketplacecatalog.ListChangeSetsInput) (*request.Request, *marketplacecatalog.ListChangeSetsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChangeSetsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*marketplacecatalog.ListChangeSetsOutput)
	return ret0, ret1
}

// ListChangeSetsRequest indicates an expected call of ListChangeSetsRequest
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListChangeSetsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChangeSetsRequest", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListChangeSetsRequest), arg0)
}

// ListChangeSetsWithContext mocks base method
func (m *MockMarketplaceCatalogAPI) ListChangeSetsWithContext(arg0 context.Context, arg1 *marketplacecatalog.ListChangeSetsInput, arg2 ...request.Option) (*marketplacecatalog.ListChangeSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListChangeSetsWithContext", varargs...)
	ret0, _ := ret[0].(*marketplacecatalog.ListChangeSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChangeSetsWithContext indicates an expected call of ListChangeSetsWithContext
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListChangeSetsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChangeSetsWithContext", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListChangeSetsWithContext), varargs...)
}

// ListEntities mocks base method
func (m *MockMarketplaceCatalogAPI) ListEntities(arg0 *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntities", arg0)
	ret0, _ := ret[0].(*marketplacecatalog.ListEntitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntities indicates an expected call of ListEntities
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListEntities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntities", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListEntities), arg0)
}

// ListEntitiesPages mocks base method
func (m *MockMarketplaceCatalogAPI) ListEntitiesPages(arg0 *marketplacecatalog.ListEntitiesInput, arg1 func(*marketplacecatalog.ListEntitiesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntitiesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListEntitiesPages indicates an expected call of ListEntitiesPages
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListEntitiesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntitiesPages", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListEntitiesPages), arg0, arg1)
}

// ListEntitiesPagesWithContext mocks base method
func (m *MockMarketplaceCatalogAPI) ListEntitiesPagesWithContext(arg0 context.Context, arg1 *marketplacecatalog.ListEntitiesInput, arg2 func(*marketplacecatalog.ListEntitiesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEntitiesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListEntitiesPagesWithContext indicates an expected call of ListEntitiesPagesWithContext
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListEntitiesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntitiesPagesWithContext", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListEntitiesPagesWithContext), varargs...)
}

// ListEntitiesRequest mocks base method
func (m *MockMarketplaceCatalogAPI) ListEntitiesRequest(arg0 *marketplacecatalog.ListEntitiesInput) (*request.Request, *marketplacecatalog.ListEntitiesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntitiesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*marketplacecatalog.ListEntitiesOutput)
	return ret0, ret1
}

// ListEntitiesRequest indicates an expected call of ListEntitiesRequest
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListEntitiesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntitiesRequest", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListEntitiesRequest), arg0)
}

// ListEntitiesWithContext mocks base method
func (m *MockMarketplaceCatalogAPI) ListEntitiesWithContext(arg0 context.Context, arg1 *marketplacecatalog.ListEntitiesInput, arg2 ...request.Option) (*marketplacecatalog.ListEntitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEntitiesWithContext", varargs...)
	ret0, _ := ret[0].(*marketplacecatalog.ListEntitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntitiesWithContext indicates an expected call of ListEntitiesWithContext
func (mr *MockMarketplaceCatalogAPIMockRecorder) ListEntitiesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntitiesWithContext", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).ListEntitiesWithContext), varargs...)
}

// StartChangeSet mocks base method
func (m *MockMarketplaceCatalogAPI) StartChangeSet(arg0 *marketplacecatalog.StartChangeSetInput) (*marketplacecatalog.StartChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartChangeSet", arg0)
	ret0, _ := ret[0].(*marketplacecatalog.StartChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartChangeSet indicates an expected call of StartChangeSet
func (mr *MockMarketplaceCatalogAPIMockRecorder) StartChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartChangeSet", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).StartChangeSet), arg0)
}

// StartChangeSetRequest mocks base method
func (m *MockMarketplaceCatalogAPI) StartChangeSetRequest(arg0 *marketplacecatalog.StartChangeSetInput) (*request.Request, *marketplacecatalog.StartChangeSetOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartChangeSetRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*marketplacecatalog.StartChangeSetOutput)
	return ret0, ret1
}

// StartChangeSetRequest indicates an expected call of StartChangeSetRequest
func (mr *MockMarketplaceCatalogAPIMockRecorder) StartChangeSetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartChangeSetRequest", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).StartChangeSetRequest), arg0)
}

// StartChangeSetWithContext mocks base method
func (m *MockMarketplaceCatalogAPI) StartChangeSetWithContext(arg0 context.Context, arg1 *marketplacecatalog.StartChangeSetInput, arg2 ...request.Option) (*marketplacecatalog.StartChangeSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartChangeSetWithContext", varargs...)
	ret0, _ := ret[0].(*marketplacecatalog.StartChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartChangeSetWithContext indicates an expected call of StartChangeSetWithContext
func (mr *MockMarketplaceCatalogAPIMockRecorder) StartChangeSetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartChangeSetWithContext", reflect.TypeOf((*MockMarketplaceCatalogAPI)(nil).StartChangeSetWithContext), varargs...)
}