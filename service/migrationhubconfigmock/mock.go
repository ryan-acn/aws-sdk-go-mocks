// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/migrationhubconfig/migrationhubconfigiface (interfaces: MigrationHubConfigAPI)

// Package migrationhubconfigmock is a generated GoMock package.
package migrationhubconfigmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	migrationhubconfig "github.com/aws/aws-sdk-go/service/migrationhubconfig"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMigrationHubConfigAPI is a mock of MigrationHubConfigAPI interface
type MockMigrationHubConfigAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationHubConfigAPIMockRecorder
}

// MockMigrationHubConfigAPIMockRecorder is the mock recorder for MockMigrationHubConfigAPI
type MockMigrationHubConfigAPIMockRecorder struct {
	mock *MockMigrationHubConfigAPI
}

// NewMockMigrationHubConfigAPI creates a new mock instance
func NewMockMigrationHubConfigAPI(ctrl *gomock.Controller) *MockMigrationHubConfigAPI {
	mock := &MockMigrationHubConfigAPI{ctrl: ctrl}
	mock.recorder = &MockMigrationHubConfigAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMigrationHubConfigAPI) EXPECT() *MockMigrationHubConfigAPIMockRecorder {
	return m.recorder
}

// CreateHomeRegionControl mocks base method
func (m *MockMigrationHubConfigAPI) CreateHomeRegionControl(arg0 *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHomeRegionControl", arg0)
	ret0, _ := ret[0].(*migrationhubconfig.CreateHomeRegionControlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHomeRegionControl indicates an expected call of CreateHomeRegionControl
func (mr *MockMigrationHubConfigAPIMockRecorder) CreateHomeRegionControl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHomeRegionControl", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).CreateHomeRegionControl), arg0)
}

// CreateHomeRegionControlRequest mocks base method
func (m *MockMigrationHubConfigAPI) CreateHomeRegionControlRequest(arg0 *migrationhubconfig.CreateHomeRegionControlInput) (*request.Request, *migrationhubconfig.CreateHomeRegionControlOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHomeRegionControlRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*migrationhubconfig.CreateHomeRegionControlOutput)
	return ret0, ret1
}

// CreateHomeRegionControlRequest indicates an expected call of CreateHomeRegionControlRequest
func (mr *MockMigrationHubConfigAPIMockRecorder) CreateHomeRegionControlRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHomeRegionControlRequest", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).CreateHomeRegionControlRequest), arg0)
}

// CreateHomeRegionControlWithContext mocks base method
func (m *MockMigrationHubConfigAPI) CreateHomeRegionControlWithContext(arg0 context.Context, arg1 *migrationhubconfig.CreateHomeRegionControlInput, arg2 ...request.Option) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateHomeRegionControlWithContext", varargs...)
	ret0, _ := ret[0].(*migrationhubconfig.CreateHomeRegionControlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHomeRegionControlWithContext indicates an expected call of CreateHomeRegionControlWithContext
func (mr *MockMigrationHubConfigAPIMockRecorder) CreateHomeRegionControlWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHomeRegionControlWithContext", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).CreateHomeRegionControlWithContext), varargs...)
}

// DescribeHomeRegionControls mocks base method
func (m *MockMigrationHubConfigAPI) DescribeHomeRegionControls(arg0 *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeHomeRegionControls", arg0)
	ret0, _ := ret[0].(*migrationhubconfig.DescribeHomeRegionControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHomeRegionControls indicates an expected call of DescribeHomeRegionControls
func (mr *MockMigrationHubConfigAPIMockRecorder) DescribeHomeRegionControls(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHomeRegionControls", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).DescribeHomeRegionControls), arg0)
}

// DescribeHomeRegionControlsPages mocks base method
func (m *MockMigrationHubConfigAPI) DescribeHomeRegionControlsPages(arg0 *migrationhubconfig.DescribeHomeRegionControlsInput, arg1 func(*migrationhubconfig.DescribeHomeRegionControlsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeHomeRegionControlsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeHomeRegionControlsPages indicates an expected call of DescribeHomeRegionControlsPages
func (mr *MockMigrationHubConfigAPIMockRecorder) DescribeHomeRegionControlsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHomeRegionControlsPages", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).DescribeHomeRegionControlsPages), arg0, arg1)
}

// DescribeHomeRegionControlsPagesWithContext mocks base method
func (m *MockMigrationHubConfigAPI) DescribeHomeRegionControlsPagesWithContext(arg0 context.Context, arg1 *migrationhubconfig.DescribeHomeRegionControlsInput, arg2 func(*migrationhubconfig.DescribeHomeRegionControlsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeHomeRegionControlsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeHomeRegionControlsPagesWithContext indicates an expected call of DescribeHomeRegionControlsPagesWithContext
func (mr *MockMigrationHubConfigAPIMockRecorder) DescribeHomeRegionControlsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHomeRegionControlsPagesWithContext", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).DescribeHomeRegionControlsPagesWithContext), varargs...)
}

// DescribeHomeRegionControlsRequest mocks base method
func (m *MockMigrationHubConfigAPI) DescribeHomeRegionControlsRequest(arg0 *migrationhubconfig.DescribeHomeRegionControlsInput) (*request.Request, *migrationhubconfig.DescribeHomeRegionControlsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeHomeRegionControlsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*migrationhubconfig.DescribeHomeRegionControlsOutput)
	return ret0, ret1
}

// DescribeHomeRegionControlsRequest indicates an expected call of DescribeHomeRegionControlsRequest
func (mr *MockMigrationHubConfigAPIMockRecorder) DescribeHomeRegionControlsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHomeRegionControlsRequest", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).DescribeHomeRegionControlsRequest), arg0)
}

// DescribeHomeRegionControlsWithContext mocks base method
func (m *MockMigrationHubConfigAPI) DescribeHomeRegionControlsWithContext(arg0 context.Context, arg1 *migrationhubconfig.DescribeHomeRegionControlsInput, arg2 ...request.Option) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeHomeRegionControlsWithContext", varargs...)
	ret0, _ := ret[0].(*migrationhubconfig.DescribeHomeRegionControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHomeRegionControlsWithContext indicates an expected call of DescribeHomeRegionControlsWithContext
func (mr *MockMigrationHubConfigAPIMockRecorder) DescribeHomeRegionControlsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHomeRegionControlsWithContext", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).DescribeHomeRegionControlsWithContext), varargs...)
}

// GetHomeRegion mocks base method
func (m *MockMigrationHubConfigAPI) GetHomeRegion(arg0 *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHomeRegion", arg0)
	ret0, _ := ret[0].(*migrationhubconfig.GetHomeRegionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHomeRegion indicates an expected call of GetHomeRegion
func (mr *MockMigrationHubConfigAPIMockRecorder) GetHomeRegion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHomeRegion", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).GetHomeRegion), arg0)
}

// GetHomeRegionRequest mocks base method
func (m *MockMigrationHubConfigAPI) GetHomeRegionRequest(arg0 *migrationhubconfig.GetHomeRegionInput) (*request.Request, *migrationhubconfig.GetHomeRegionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHomeRegionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*migrationhubconfig.GetHomeRegionOutput)
	return ret0, ret1
}

// GetHomeRegionRequest indicates an expected call of GetHomeRegionRequest
func (mr *MockMigrationHubConfigAPIMockRecorder) GetHomeRegionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHomeRegionRequest", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).GetHomeRegionRequest), arg0)
}

// GetHomeRegionWithContext mocks base method
func (m *MockMigrationHubConfigAPI) GetHomeRegionWithContext(arg0 context.Context, arg1 *migrationhubconfig.GetHomeRegionInput, arg2 ...request.Option) (*migrationhubconfig.GetHomeRegionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHomeRegionWithContext", varargs...)
	ret0, _ := ret[0].(*migrationhubconfig.GetHomeRegionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHomeRegionWithContext indicates an expected call of GetHomeRegionWithContext
func (mr *MockMigrationHubConfigAPIMockRecorder) GetHomeRegionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHomeRegionWithContext", reflect.TypeOf((*MockMigrationHubConfigAPI)(nil).GetHomeRegionWithContext), varargs...)
}
