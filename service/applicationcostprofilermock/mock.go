// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/applicationcostprofiler/applicationcostprofileriface (interfaces: ApplicationCostProfilerAPI)

// Package applicationcostprofilermock is a generated GoMock package.
package applicationcostprofilermock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	applicationcostprofiler "github.com/aws/aws-sdk-go/service/applicationcostprofiler"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApplicationCostProfilerAPI is a mock of ApplicationCostProfilerAPI interface
type MockApplicationCostProfilerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationCostProfilerAPIMockRecorder
}

// MockApplicationCostProfilerAPIMockRecorder is the mock recorder for MockApplicationCostProfilerAPI
type MockApplicationCostProfilerAPIMockRecorder struct {
	mock *MockApplicationCostProfilerAPI
}

// NewMockApplicationCostProfilerAPI creates a new mock instance
func NewMockApplicationCostProfilerAPI(ctrl *gomock.Controller) *MockApplicationCostProfilerAPI {
	mock := &MockApplicationCostProfilerAPI{ctrl: ctrl}
	mock.recorder = &MockApplicationCostProfilerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationCostProfilerAPI) EXPECT() *MockApplicationCostProfilerAPIMockRecorder {
	return m.recorder
}

// DeleteReportDefinition mocks base method
func (m *MockApplicationCostProfilerAPI) DeleteReportDefinition(arg0 *applicationcostprofiler.DeleteReportDefinitionInput) (*applicationcostprofiler.DeleteReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReportDefinition", arg0)
	ret0, _ := ret[0].(*applicationcostprofiler.DeleteReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReportDefinition indicates an expected call of DeleteReportDefinition
func (mr *MockApplicationCostProfilerAPIMockRecorder) DeleteReportDefinition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReportDefinition", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).DeleteReportDefinition), arg0)
}

// DeleteReportDefinitionRequest mocks base method
func (m *MockApplicationCostProfilerAPI) DeleteReportDefinitionRequest(arg0 *applicationcostprofiler.DeleteReportDefinitionInput) (*request.Request, *applicationcostprofiler.DeleteReportDefinitionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReportDefinitionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationcostprofiler.DeleteReportDefinitionOutput)
	return ret0, ret1
}

// DeleteReportDefinitionRequest indicates an expected call of DeleteReportDefinitionRequest
func (mr *MockApplicationCostProfilerAPIMockRecorder) DeleteReportDefinitionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReportDefinitionRequest", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).DeleteReportDefinitionRequest), arg0)
}

// DeleteReportDefinitionWithContext mocks base method
func (m *MockApplicationCostProfilerAPI) DeleteReportDefinitionWithContext(arg0 context.Context, arg1 *applicationcostprofiler.DeleteReportDefinitionInput, arg2 ...request.Option) (*applicationcostprofiler.DeleteReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteReportDefinitionWithContext", varargs...)
	ret0, _ := ret[0].(*applicationcostprofiler.DeleteReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReportDefinitionWithContext indicates an expected call of DeleteReportDefinitionWithContext
func (mr *MockApplicationCostProfilerAPIMockRecorder) DeleteReportDefinitionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReportDefinitionWithContext", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).DeleteReportDefinitionWithContext), varargs...)
}

// GetReportDefinition mocks base method
func (m *MockApplicationCostProfilerAPI) GetReportDefinition(arg0 *applicationcostprofiler.GetReportDefinitionInput) (*applicationcostprofiler.GetReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportDefinition", arg0)
	ret0, _ := ret[0].(*applicationcostprofiler.GetReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportDefinition indicates an expected call of GetReportDefinition
func (mr *MockApplicationCostProfilerAPIMockRecorder) GetReportDefinition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportDefinition", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).GetReportDefinition), arg0)
}

// GetReportDefinitionRequest mocks base method
func (m *MockApplicationCostProfilerAPI) GetReportDefinitionRequest(arg0 *applicationcostprofiler.GetReportDefinitionInput) (*request.Request, *applicationcostprofiler.GetReportDefinitionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportDefinitionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationcostprofiler.GetReportDefinitionOutput)
	return ret0, ret1
}

// GetReportDefinitionRequest indicates an expected call of GetReportDefinitionRequest
func (mr *MockApplicationCostProfilerAPIMockRecorder) GetReportDefinitionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportDefinitionRequest", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).GetReportDefinitionRequest), arg0)
}

// GetReportDefinitionWithContext mocks base method
func (m *MockApplicationCostProfilerAPI) GetReportDefinitionWithContext(arg0 context.Context, arg1 *applicationcostprofiler.GetReportDefinitionInput, arg2 ...request.Option) (*applicationcostprofiler.GetReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReportDefinitionWithContext", varargs...)
	ret0, _ := ret[0].(*applicationcostprofiler.GetReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportDefinitionWithContext indicates an expected call of GetReportDefinitionWithContext
func (mr *MockApplicationCostProfilerAPIMockRecorder) GetReportDefinitionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportDefinitionWithContext", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).GetReportDefinitionWithContext), varargs...)
}

// ImportApplicationUsage mocks base method
func (m *MockApplicationCostProfilerAPI) ImportApplicationUsage(arg0 *applicationcostprofiler.ImportApplicationUsageInput) (*applicationcostprofiler.ImportApplicationUsageOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportApplicationUsage", arg0)
	ret0, _ := ret[0].(*applicationcostprofiler.ImportApplicationUsageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportApplicationUsage indicates an expected call of ImportApplicationUsage
func (mr *MockApplicationCostProfilerAPIMockRecorder) ImportApplicationUsage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportApplicationUsage", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).ImportApplicationUsage), arg0)
}

// ImportApplicationUsageRequest mocks base method
func (m *MockApplicationCostProfilerAPI) ImportApplicationUsageRequest(arg0 *applicationcostprofiler.ImportApplicationUsageInput) (*request.Request, *applicationcostprofiler.ImportApplicationUsageOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportApplicationUsageRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationcostprofiler.ImportApplicationUsageOutput)
	return ret0, ret1
}

// ImportApplicationUsageRequest indicates an expected call of ImportApplicationUsageRequest
func (mr *MockApplicationCostProfilerAPIMockRecorder) ImportApplicationUsageRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportApplicationUsageRequest", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).ImportApplicationUsageRequest), arg0)
}

// ImportApplicationUsageWithContext mocks base method
func (m *MockApplicationCostProfilerAPI) ImportApplicationUsageWithContext(arg0 context.Context, arg1 *applicationcostprofiler.ImportApplicationUsageInput, arg2 ...request.Option) (*applicationcostprofiler.ImportApplicationUsageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportApplicationUsageWithContext", varargs...)
	ret0, _ := ret[0].(*applicationcostprofiler.ImportApplicationUsageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportApplicationUsageWithContext indicates an expected call of ImportApplicationUsageWithContext
func (mr *MockApplicationCostProfilerAPIMockRecorder) ImportApplicationUsageWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportApplicationUsageWithContext", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).ImportApplicationUsageWithContext), varargs...)
}

// ListReportDefinitions mocks base method
func (m *MockApplicationCostProfilerAPI) ListReportDefinitions(arg0 *applicationcostprofiler.ListReportDefinitionsInput) (*applicationcostprofiler.ListReportDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReportDefinitions", arg0)
	ret0, _ := ret[0].(*applicationcostprofiler.ListReportDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReportDefinitions indicates an expected call of ListReportDefinitions
func (mr *MockApplicationCostProfilerAPIMockRecorder) ListReportDefinitions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportDefinitions", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).ListReportDefinitions), arg0)
}

// ListReportDefinitionsPages mocks base method
func (m *MockApplicationCostProfilerAPI) ListReportDefinitionsPages(arg0 *applicationcostprofiler.ListReportDefinitionsInput, arg1 func(*applicationcostprofiler.ListReportDefinitionsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReportDefinitionsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListReportDefinitionsPages indicates an expected call of ListReportDefinitionsPages
func (mr *MockApplicationCostProfilerAPIMockRecorder) ListReportDefinitionsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportDefinitionsPages", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).ListReportDefinitionsPages), arg0, arg1)
}

// ListReportDefinitionsPagesWithContext mocks base method
func (m *MockApplicationCostProfilerAPI) ListReportDefinitionsPagesWithContext(arg0 context.Context, arg1 *applicationcostprofiler.ListReportDefinitionsInput, arg2 func(*applicationcostprofiler.ListReportDefinitionsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReportDefinitionsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListReportDefinitionsPagesWithContext indicates an expected call of ListReportDefinitionsPagesWithContext
func (mr *MockApplicationCostProfilerAPIMockRecorder) ListReportDefinitionsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportDefinitionsPagesWithContext", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).ListReportDefinitionsPagesWithContext), varargs...)
}

// ListReportDefinitionsRequest mocks base method
func (m *MockApplicationCostProfilerAPI) ListReportDefinitionsRequest(arg0 *applicationcostprofiler.ListReportDefinitionsInput) (*request.Request, *applicationcostprofiler.ListReportDefinitionsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReportDefinitionsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationcostprofiler.ListReportDefinitionsOutput)
	return ret0, ret1
}

// ListReportDefinitionsRequest indicates an expected call of ListReportDefinitionsRequest
func (mr *MockApplicationCostProfilerAPIMockRecorder) ListReportDefinitionsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportDefinitionsRequest", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).ListReportDefinitionsRequest), arg0)
}

// ListReportDefinitionsWithContext mocks base method
func (m *MockApplicationCostProfilerAPI) ListReportDefinitionsWithContext(arg0 context.Context, arg1 *applicationcostprofiler.ListReportDefinitionsInput, arg2 ...request.Option) (*applicationcostprofiler.ListReportDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReportDefinitionsWithContext", varargs...)
	ret0, _ := ret[0].(*applicationcostprofiler.ListReportDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReportDefinitionsWithContext indicates an expected call of ListReportDefinitionsWithContext
func (mr *MockApplicationCostProfilerAPIMockRecorder) ListReportDefinitionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportDefinitionsWithContext", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).ListReportDefinitionsWithContext), varargs...)
}

// PutReportDefinition mocks base method
func (m *MockApplicationCostProfilerAPI) PutReportDefinition(arg0 *applicationcostprofiler.PutReportDefinitionInput) (*applicationcostprofiler.PutReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutReportDefinition", arg0)
	ret0, _ := ret[0].(*applicationcostprofiler.PutReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutReportDefinition indicates an expected call of PutReportDefinition
func (mr *MockApplicationCostProfilerAPIMockRecorder) PutReportDefinition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutReportDefinition", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).PutReportDefinition), arg0)
}

// PutReportDefinitionRequest mocks base method
func (m *MockApplicationCostProfilerAPI) PutReportDefinitionRequest(arg0 *applicationcostprofiler.PutReportDefinitionInput) (*request.Request, *applicationcostprofiler.PutReportDefinitionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutReportDefinitionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationcostprofiler.PutReportDefinitionOutput)
	return ret0, ret1
}

// PutReportDefinitionRequest indicates an expected call of PutReportDefinitionRequest
func (mr *MockApplicationCostProfilerAPIMockRecorder) PutReportDefinitionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutReportDefinitionRequest", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).PutReportDefinitionRequest), arg0)
}

// PutReportDefinitionWithContext mocks base method
func (m *MockApplicationCostProfilerAPI) PutReportDefinitionWithContext(arg0 context.Context, arg1 *applicationcostprofiler.PutReportDefinitionInput, arg2 ...request.Option) (*applicationcostprofiler.PutReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutReportDefinitionWithContext", varargs...)
	ret0, _ := ret[0].(*applicationcostprofiler.PutReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutReportDefinitionWithContext indicates an expected call of PutReportDefinitionWithContext
func (mr *MockApplicationCostProfilerAPIMockRecorder) PutReportDefinitionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutReportDefinitionWithContext", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).PutReportDefinitionWithContext), varargs...)
}

// UpdateReportDefinition mocks base method
func (m *MockApplicationCostProfilerAPI) UpdateReportDefinition(arg0 *applicationcostprofiler.UpdateReportDefinitionInput) (*applicationcostprofiler.UpdateReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReportDefinition", arg0)
	ret0, _ := ret[0].(*applicationcostprofiler.UpdateReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReportDefinition indicates an expected call of UpdateReportDefinition
func (mr *MockApplicationCostProfilerAPIMockRecorder) UpdateReportDefinition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReportDefinition", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).UpdateReportDefinition), arg0)
}

// UpdateReportDefinitionRequest mocks base method
func (m *MockApplicationCostProfilerAPI) UpdateReportDefinitionRequest(arg0 *applicationcostprofiler.UpdateReportDefinitionInput) (*request.Request, *applicationcostprofiler.UpdateReportDefinitionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReportDefinitionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationcostprofiler.UpdateReportDefinitionOutput)
	return ret0, ret1
}

// UpdateReportDefinitionRequest indicates an expected call of UpdateReportDefinitionRequest
func (mr *MockApplicationCostProfilerAPIMockRecorder) UpdateReportDefinitionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReportDefinitionRequest", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).UpdateReportDefinitionRequest), arg0)
}

// UpdateReportDefinitionWithContext mocks base method
func (m *MockApplicationCostProfilerAPI) UpdateReportDefinitionWithContext(arg0 context.Context, arg1 *applicationcostprofiler.UpdateReportDefinitionInput, arg2 ...request.Option) (*applicationcostprofiler.UpdateReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateReportDefinitionWithContext", varargs...)
	ret0, _ := ret[0].(*applicationcostprofiler.UpdateReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReportDefinitionWithContext indicates an expected call of UpdateReportDefinitionWithContext
func (mr *MockApplicationCostProfilerAPIMockRecorder) UpdateReportDefinitionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReportDefinitionWithContext", reflect.TypeOf((*MockApplicationCostProfilerAPI)(nil).UpdateReportDefinitionWithContext), varargs...)
}