// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface (interfaces: MarketplaceEntitlementServiceAPI)

// Package marketplaceentitlementservicemock is a generated GoMock package.
package marketplaceentitlementservicemock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	marketplaceentitlementservice "github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMarketplaceEntitlementServiceAPI is a mock of MarketplaceEntitlementServiceAPI interface
type MockMarketplaceEntitlementServiceAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMarketplaceEntitlementServiceAPIMockRecorder
}

// MockMarketplaceEntitlementServiceAPIMockRecorder is the mock recorder for MockMarketplaceEntitlementServiceAPI
type MockMarketplaceEntitlementServiceAPIMockRecorder struct {
	mock *MockMarketplaceEntitlementServiceAPI
}

// NewMockMarketplaceEntitlementServiceAPI creates a new mock instance
func NewMockMarketplaceEntitlementServiceAPI(ctrl *gomock.Controller) *MockMarketplaceEntitlementServiceAPI {
	mock := &MockMarketplaceEntitlementServiceAPI{ctrl: ctrl}
	mock.recorder = &MockMarketplaceEntitlementServiceAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarketplaceEntitlementServiceAPI) EXPECT() *MockMarketplaceEntitlementServiceAPIMockRecorder {
	return m.recorder
}

// GetEntitlements mocks base method
func (m *MockMarketplaceEntitlementServiceAPI) GetEntitlements(arg0 *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntitlements", arg0)
	ret0, _ := ret[0].(*marketplaceentitlementservice.GetEntitlementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntitlements indicates an expected call of GetEntitlements
func (mr *MockMarketplaceEntitlementServiceAPIMockRecorder) GetEntitlements(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntitlements", reflect.TypeOf((*MockMarketplaceEntitlementServiceAPI)(nil).GetEntitlements), arg0)
}

// GetEntitlementsRequest mocks base method
func (m *MockMarketplaceEntitlementServiceAPI) GetEntitlementsRequest(arg0 *marketplaceentitlementservice.GetEntitlementsInput) (*request.Request, *marketplaceentitlementservice.GetEntitlementsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntitlementsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*marketplaceentitlementservice.GetEntitlementsOutput)
	return ret0, ret1
}

// GetEntitlementsRequest indicates an expected call of GetEntitlementsRequest
func (mr *MockMarketplaceEntitlementServiceAPIMockRecorder) GetEntitlementsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntitlementsRequest", reflect.TypeOf((*MockMarketplaceEntitlementServiceAPI)(nil).GetEntitlementsRequest), arg0)
}

// GetEntitlementsWithContext mocks base method
func (m *MockMarketplaceEntitlementServiceAPI) GetEntitlementsWithContext(arg0 aws.Context, arg1 *marketplaceentitlementservice.GetEntitlementsInput, arg2 ...request.Option) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEntitlementsWithContext", varargs...)
	ret0, _ := ret[0].(*marketplaceentitlementservice.GetEntitlementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntitlementsWithContext indicates an expected call of GetEntitlementsWithContext
func (mr *MockMarketplaceEntitlementServiceAPIMockRecorder) GetEntitlementsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntitlementsWithContext", reflect.TypeOf((*MockMarketplaceEntitlementServiceAPI)(nil).GetEntitlementsWithContext), varargs...)
}