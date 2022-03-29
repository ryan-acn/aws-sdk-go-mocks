// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/synthetics/syntheticsiface (interfaces: SyntheticsAPI)

// Package syntheticsmock is a generated GoMock package.
package syntheticsmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	synthetics "github.com/aws/aws-sdk-go/service/synthetics"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSyntheticsAPI is a mock of SyntheticsAPI interface
type MockSyntheticsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSyntheticsAPIMockRecorder
}

// MockSyntheticsAPIMockRecorder is the mock recorder for MockSyntheticsAPI
type MockSyntheticsAPIMockRecorder struct {
	mock *MockSyntheticsAPI
}

// NewMockSyntheticsAPI creates a new mock instance
func NewMockSyntheticsAPI(ctrl *gomock.Controller) *MockSyntheticsAPI {
	mock := &MockSyntheticsAPI{ctrl: ctrl}
	mock.recorder = &MockSyntheticsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSyntheticsAPI) EXPECT() *MockSyntheticsAPIMockRecorder {
	return m.recorder
}

// CreateCanary mocks base method
func (m *MockSyntheticsAPI) CreateCanary(arg0 *synthetics.CreateCanaryInput) (*synthetics.CreateCanaryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCanary", arg0)
	ret0, _ := ret[0].(*synthetics.CreateCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCanary indicates an expected call of CreateCanary
func (mr *MockSyntheticsAPIMockRecorder) CreateCanary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCanary", reflect.TypeOf((*MockSyntheticsAPI)(nil).CreateCanary), arg0)
}

// CreateCanaryRequest mocks base method
func (m *MockSyntheticsAPI) CreateCanaryRequest(arg0 *synthetics.CreateCanaryInput) (*request.Request, *synthetics.CreateCanaryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCanaryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.CreateCanaryOutput)
	return ret0, ret1
}

// CreateCanaryRequest indicates an expected call of CreateCanaryRequest
func (mr *MockSyntheticsAPIMockRecorder) CreateCanaryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCanaryRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).CreateCanaryRequest), arg0)
}

// CreateCanaryWithContext mocks base method
func (m *MockSyntheticsAPI) CreateCanaryWithContext(arg0 context.Context, arg1 *synthetics.CreateCanaryInput, arg2 ...request.Option) (*synthetics.CreateCanaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCanaryWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.CreateCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCanaryWithContext indicates an expected call of CreateCanaryWithContext
func (mr *MockSyntheticsAPIMockRecorder) CreateCanaryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCanaryWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).CreateCanaryWithContext), varargs...)
}

// DeleteCanary mocks base method
func (m *MockSyntheticsAPI) DeleteCanary(arg0 *synthetics.DeleteCanaryInput) (*synthetics.DeleteCanaryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCanary", arg0)
	ret0, _ := ret[0].(*synthetics.DeleteCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCanary indicates an expected call of DeleteCanary
func (mr *MockSyntheticsAPIMockRecorder) DeleteCanary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCanary", reflect.TypeOf((*MockSyntheticsAPI)(nil).DeleteCanary), arg0)
}

// DeleteCanaryRequest mocks base method
func (m *MockSyntheticsAPI) DeleteCanaryRequest(arg0 *synthetics.DeleteCanaryInput) (*request.Request, *synthetics.DeleteCanaryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCanaryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.DeleteCanaryOutput)
	return ret0, ret1
}

// DeleteCanaryRequest indicates an expected call of DeleteCanaryRequest
func (mr *MockSyntheticsAPIMockRecorder) DeleteCanaryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCanaryRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).DeleteCanaryRequest), arg0)
}

// DeleteCanaryWithContext mocks base method
func (m *MockSyntheticsAPI) DeleteCanaryWithContext(arg0 context.Context, arg1 *synthetics.DeleteCanaryInput, arg2 ...request.Option) (*synthetics.DeleteCanaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCanaryWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.DeleteCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCanaryWithContext indicates an expected call of DeleteCanaryWithContext
func (mr *MockSyntheticsAPIMockRecorder) DeleteCanaryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCanaryWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).DeleteCanaryWithContext), varargs...)
}

// DescribeCanaries mocks base method
func (m *MockSyntheticsAPI) DescribeCanaries(arg0 *synthetics.DescribeCanariesInput) (*synthetics.DescribeCanariesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCanaries", arg0)
	ret0, _ := ret[0].(*synthetics.DescribeCanariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCanaries indicates an expected call of DescribeCanaries
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanaries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanaries", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanaries), arg0)
}

// DescribeCanariesLastRun mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesLastRun(arg0 *synthetics.DescribeCanariesLastRunInput) (*synthetics.DescribeCanariesLastRunOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCanariesLastRun", arg0)
	ret0, _ := ret[0].(*synthetics.DescribeCanariesLastRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCanariesLastRun indicates an expected call of DescribeCanariesLastRun
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesLastRun(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesLastRun", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesLastRun), arg0)
}

// DescribeCanariesLastRunPages mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesLastRunPages(arg0 *synthetics.DescribeCanariesLastRunInput, arg1 func(*synthetics.DescribeCanariesLastRunOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCanariesLastRunPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeCanariesLastRunPages indicates an expected call of DescribeCanariesLastRunPages
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesLastRunPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesLastRunPages", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesLastRunPages), arg0, arg1)
}

// DescribeCanariesLastRunPagesWithContext mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesLastRunPagesWithContext(arg0 context.Context, arg1 *synthetics.DescribeCanariesLastRunInput, arg2 func(*synthetics.DescribeCanariesLastRunOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCanariesLastRunPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeCanariesLastRunPagesWithContext indicates an expected call of DescribeCanariesLastRunPagesWithContext
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesLastRunPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesLastRunPagesWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesLastRunPagesWithContext), varargs...)
}

// DescribeCanariesLastRunRequest mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesLastRunRequest(arg0 *synthetics.DescribeCanariesLastRunInput) (*request.Request, *synthetics.DescribeCanariesLastRunOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCanariesLastRunRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.DescribeCanariesLastRunOutput)
	return ret0, ret1
}

// DescribeCanariesLastRunRequest indicates an expected call of DescribeCanariesLastRunRequest
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesLastRunRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesLastRunRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesLastRunRequest), arg0)
}

// DescribeCanariesLastRunWithContext mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesLastRunWithContext(arg0 context.Context, arg1 *synthetics.DescribeCanariesLastRunInput, arg2 ...request.Option) (*synthetics.DescribeCanariesLastRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCanariesLastRunWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.DescribeCanariesLastRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCanariesLastRunWithContext indicates an expected call of DescribeCanariesLastRunWithContext
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesLastRunWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesLastRunWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesLastRunWithContext), varargs...)
}

// DescribeCanariesPages mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesPages(arg0 *synthetics.DescribeCanariesInput, arg1 func(*synthetics.DescribeCanariesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCanariesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeCanariesPages indicates an expected call of DescribeCanariesPages
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesPages", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesPages), arg0, arg1)
}

// DescribeCanariesPagesWithContext mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesPagesWithContext(arg0 context.Context, arg1 *synthetics.DescribeCanariesInput, arg2 func(*synthetics.DescribeCanariesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCanariesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeCanariesPagesWithContext indicates an expected call of DescribeCanariesPagesWithContext
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesPagesWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesPagesWithContext), varargs...)
}

// DescribeCanariesRequest mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesRequest(arg0 *synthetics.DescribeCanariesInput) (*request.Request, *synthetics.DescribeCanariesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCanariesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.DescribeCanariesOutput)
	return ret0, ret1
}

// DescribeCanariesRequest indicates an expected call of DescribeCanariesRequest
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesRequest), arg0)
}

// DescribeCanariesWithContext mocks base method
func (m *MockSyntheticsAPI) DescribeCanariesWithContext(arg0 context.Context, arg1 *synthetics.DescribeCanariesInput, arg2 ...request.Option) (*synthetics.DescribeCanariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCanariesWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.DescribeCanariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCanariesWithContext indicates an expected call of DescribeCanariesWithContext
func (mr *MockSyntheticsAPIMockRecorder) DescribeCanariesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCanariesWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeCanariesWithContext), varargs...)
}

// DescribeRuntimeVersions mocks base method
func (m *MockSyntheticsAPI) DescribeRuntimeVersions(arg0 *synthetics.DescribeRuntimeVersionsInput) (*synthetics.DescribeRuntimeVersionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRuntimeVersions", arg0)
	ret0, _ := ret[0].(*synthetics.DescribeRuntimeVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRuntimeVersions indicates an expected call of DescribeRuntimeVersions
func (mr *MockSyntheticsAPIMockRecorder) DescribeRuntimeVersions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRuntimeVersions", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeRuntimeVersions), arg0)
}

// DescribeRuntimeVersionsPages mocks base method
func (m *MockSyntheticsAPI) DescribeRuntimeVersionsPages(arg0 *synthetics.DescribeRuntimeVersionsInput, arg1 func(*synthetics.DescribeRuntimeVersionsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRuntimeVersionsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeRuntimeVersionsPages indicates an expected call of DescribeRuntimeVersionsPages
func (mr *MockSyntheticsAPIMockRecorder) DescribeRuntimeVersionsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRuntimeVersionsPages", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeRuntimeVersionsPages), arg0, arg1)
}

// DescribeRuntimeVersionsPagesWithContext mocks base method
func (m *MockSyntheticsAPI) DescribeRuntimeVersionsPagesWithContext(arg0 context.Context, arg1 *synthetics.DescribeRuntimeVersionsInput, arg2 func(*synthetics.DescribeRuntimeVersionsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRuntimeVersionsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeRuntimeVersionsPagesWithContext indicates an expected call of DescribeRuntimeVersionsPagesWithContext
func (mr *MockSyntheticsAPIMockRecorder) DescribeRuntimeVersionsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRuntimeVersionsPagesWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeRuntimeVersionsPagesWithContext), varargs...)
}

// DescribeRuntimeVersionsRequest mocks base method
func (m *MockSyntheticsAPI) DescribeRuntimeVersionsRequest(arg0 *synthetics.DescribeRuntimeVersionsInput) (*request.Request, *synthetics.DescribeRuntimeVersionsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRuntimeVersionsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.DescribeRuntimeVersionsOutput)
	return ret0, ret1
}

// DescribeRuntimeVersionsRequest indicates an expected call of DescribeRuntimeVersionsRequest
func (mr *MockSyntheticsAPIMockRecorder) DescribeRuntimeVersionsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRuntimeVersionsRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeRuntimeVersionsRequest), arg0)
}

// DescribeRuntimeVersionsWithContext mocks base method
func (m *MockSyntheticsAPI) DescribeRuntimeVersionsWithContext(arg0 context.Context, arg1 *synthetics.DescribeRuntimeVersionsInput, arg2 ...request.Option) (*synthetics.DescribeRuntimeVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRuntimeVersionsWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.DescribeRuntimeVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRuntimeVersionsWithContext indicates an expected call of DescribeRuntimeVersionsWithContext
func (mr *MockSyntheticsAPIMockRecorder) DescribeRuntimeVersionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRuntimeVersionsWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).DescribeRuntimeVersionsWithContext), varargs...)
}

// GetCanary mocks base method
func (m *MockSyntheticsAPI) GetCanary(arg0 *synthetics.GetCanaryInput) (*synthetics.GetCanaryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCanary", arg0)
	ret0, _ := ret[0].(*synthetics.GetCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCanary indicates an expected call of GetCanary
func (mr *MockSyntheticsAPIMockRecorder) GetCanary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanary", reflect.TypeOf((*MockSyntheticsAPI)(nil).GetCanary), arg0)
}

// GetCanaryRequest mocks base method
func (m *MockSyntheticsAPI) GetCanaryRequest(arg0 *synthetics.GetCanaryInput) (*request.Request, *synthetics.GetCanaryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCanaryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.GetCanaryOutput)
	return ret0, ret1
}

// GetCanaryRequest indicates an expected call of GetCanaryRequest
func (mr *MockSyntheticsAPIMockRecorder) GetCanaryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanaryRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).GetCanaryRequest), arg0)
}

// GetCanaryRuns mocks base method
func (m *MockSyntheticsAPI) GetCanaryRuns(arg0 *synthetics.GetCanaryRunsInput) (*synthetics.GetCanaryRunsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCanaryRuns", arg0)
	ret0, _ := ret[0].(*synthetics.GetCanaryRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCanaryRuns indicates an expected call of GetCanaryRuns
func (mr *MockSyntheticsAPIMockRecorder) GetCanaryRuns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanaryRuns", reflect.TypeOf((*MockSyntheticsAPI)(nil).GetCanaryRuns), arg0)
}

// GetCanaryRunsPages mocks base method
func (m *MockSyntheticsAPI) GetCanaryRunsPages(arg0 *synthetics.GetCanaryRunsInput, arg1 func(*synthetics.GetCanaryRunsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCanaryRunsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetCanaryRunsPages indicates an expected call of GetCanaryRunsPages
func (mr *MockSyntheticsAPIMockRecorder) GetCanaryRunsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanaryRunsPages", reflect.TypeOf((*MockSyntheticsAPI)(nil).GetCanaryRunsPages), arg0, arg1)
}

// GetCanaryRunsPagesWithContext mocks base method
func (m *MockSyntheticsAPI) GetCanaryRunsPagesWithContext(arg0 context.Context, arg1 *synthetics.GetCanaryRunsInput, arg2 func(*synthetics.GetCanaryRunsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCanaryRunsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetCanaryRunsPagesWithContext indicates an expected call of GetCanaryRunsPagesWithContext
func (mr *MockSyntheticsAPIMockRecorder) GetCanaryRunsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanaryRunsPagesWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).GetCanaryRunsPagesWithContext), varargs...)
}

// GetCanaryRunsRequest mocks base method
func (m *MockSyntheticsAPI) GetCanaryRunsRequest(arg0 *synthetics.GetCanaryRunsInput) (*request.Request, *synthetics.GetCanaryRunsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCanaryRunsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.GetCanaryRunsOutput)
	return ret0, ret1
}

// GetCanaryRunsRequest indicates an expected call of GetCanaryRunsRequest
func (mr *MockSyntheticsAPIMockRecorder) GetCanaryRunsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanaryRunsRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).GetCanaryRunsRequest), arg0)
}

// GetCanaryRunsWithContext mocks base method
func (m *MockSyntheticsAPI) GetCanaryRunsWithContext(arg0 context.Context, arg1 *synthetics.GetCanaryRunsInput, arg2 ...request.Option) (*synthetics.GetCanaryRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCanaryRunsWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.GetCanaryRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCanaryRunsWithContext indicates an expected call of GetCanaryRunsWithContext
func (mr *MockSyntheticsAPIMockRecorder) GetCanaryRunsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanaryRunsWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).GetCanaryRunsWithContext), varargs...)
}

// GetCanaryWithContext mocks base method
func (m *MockSyntheticsAPI) GetCanaryWithContext(arg0 context.Context, arg1 *synthetics.GetCanaryInput, arg2 ...request.Option) (*synthetics.GetCanaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCanaryWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.GetCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCanaryWithContext indicates an expected call of GetCanaryWithContext
func (mr *MockSyntheticsAPIMockRecorder) GetCanaryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCanaryWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).GetCanaryWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockSyntheticsAPI) ListTagsForResource(arg0 *synthetics.ListTagsForResourceInput) (*synthetics.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*synthetics.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockSyntheticsAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockSyntheticsAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockSyntheticsAPI) ListTagsForResourceRequest(arg0 *synthetics.ListTagsForResourceInput) (*request.Request, *synthetics.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockSyntheticsAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockSyntheticsAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *synthetics.ListTagsForResourceInput, arg2 ...request.Option) (*synthetics.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockSyntheticsAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// StartCanary mocks base method
func (m *MockSyntheticsAPI) StartCanary(arg0 *synthetics.StartCanaryInput) (*synthetics.StartCanaryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCanary", arg0)
	ret0, _ := ret[0].(*synthetics.StartCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartCanary indicates an expected call of StartCanary
func (mr *MockSyntheticsAPIMockRecorder) StartCanary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCanary", reflect.TypeOf((*MockSyntheticsAPI)(nil).StartCanary), arg0)
}

// StartCanaryRequest mocks base method
func (m *MockSyntheticsAPI) StartCanaryRequest(arg0 *synthetics.StartCanaryInput) (*request.Request, *synthetics.StartCanaryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCanaryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.StartCanaryOutput)
	return ret0, ret1
}

// StartCanaryRequest indicates an expected call of StartCanaryRequest
func (mr *MockSyntheticsAPIMockRecorder) StartCanaryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCanaryRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).StartCanaryRequest), arg0)
}

// StartCanaryWithContext mocks base method
func (m *MockSyntheticsAPI) StartCanaryWithContext(arg0 context.Context, arg1 *synthetics.StartCanaryInput, arg2 ...request.Option) (*synthetics.StartCanaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartCanaryWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.StartCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartCanaryWithContext indicates an expected call of StartCanaryWithContext
func (mr *MockSyntheticsAPIMockRecorder) StartCanaryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCanaryWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).StartCanaryWithContext), varargs...)
}

// StopCanary mocks base method
func (m *MockSyntheticsAPI) StopCanary(arg0 *synthetics.StopCanaryInput) (*synthetics.StopCanaryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopCanary", arg0)
	ret0, _ := ret[0].(*synthetics.StopCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopCanary indicates an expected call of StopCanary
func (mr *MockSyntheticsAPIMockRecorder) StopCanary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopCanary", reflect.TypeOf((*MockSyntheticsAPI)(nil).StopCanary), arg0)
}

// StopCanaryRequest mocks base method
func (m *MockSyntheticsAPI) StopCanaryRequest(arg0 *synthetics.StopCanaryInput) (*request.Request, *synthetics.StopCanaryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopCanaryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.StopCanaryOutput)
	return ret0, ret1
}

// StopCanaryRequest indicates an expected call of StopCanaryRequest
func (mr *MockSyntheticsAPIMockRecorder) StopCanaryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopCanaryRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).StopCanaryRequest), arg0)
}

// StopCanaryWithContext mocks base method
func (m *MockSyntheticsAPI) StopCanaryWithContext(arg0 context.Context, arg1 *synthetics.StopCanaryInput, arg2 ...request.Option) (*synthetics.StopCanaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopCanaryWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.StopCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopCanaryWithContext indicates an expected call of StopCanaryWithContext
func (mr *MockSyntheticsAPIMockRecorder) StopCanaryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopCanaryWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).StopCanaryWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockSyntheticsAPI) TagResource(arg0 *synthetics.TagResourceInput) (*synthetics.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*synthetics.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockSyntheticsAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockSyntheticsAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockSyntheticsAPI) TagResourceRequest(arg0 *synthetics.TagResourceInput) (*request.Request, *synthetics.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockSyntheticsAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockSyntheticsAPI) TagResourceWithContext(arg0 context.Context, arg1 *synthetics.TagResourceInput, arg2 ...request.Option) (*synthetics.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockSyntheticsAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockSyntheticsAPI) UntagResource(arg0 *synthetics.UntagResourceInput) (*synthetics.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*synthetics.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockSyntheticsAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockSyntheticsAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockSyntheticsAPI) UntagResourceRequest(arg0 *synthetics.UntagResourceInput) (*request.Request, *synthetics.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockSyntheticsAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockSyntheticsAPI) UntagResourceWithContext(arg0 context.Context, arg1 *synthetics.UntagResourceInput, arg2 ...request.Option) (*synthetics.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockSyntheticsAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).UntagResourceWithContext), varargs...)
}

// UpdateCanary mocks base method
func (m *MockSyntheticsAPI) UpdateCanary(arg0 *synthetics.UpdateCanaryInput) (*synthetics.UpdateCanaryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCanary", arg0)
	ret0, _ := ret[0].(*synthetics.UpdateCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCanary indicates an expected call of UpdateCanary
func (mr *MockSyntheticsAPIMockRecorder) UpdateCanary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCanary", reflect.TypeOf((*MockSyntheticsAPI)(nil).UpdateCanary), arg0)
}

// UpdateCanaryRequest mocks base method
func (m *MockSyntheticsAPI) UpdateCanaryRequest(arg0 *synthetics.UpdateCanaryInput) (*request.Request, *synthetics.UpdateCanaryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCanaryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*synthetics.UpdateCanaryOutput)
	return ret0, ret1
}

// UpdateCanaryRequest indicates an expected call of UpdateCanaryRequest
func (mr *MockSyntheticsAPIMockRecorder) UpdateCanaryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCanaryRequest", reflect.TypeOf((*MockSyntheticsAPI)(nil).UpdateCanaryRequest), arg0)
}

// UpdateCanaryWithContext mocks base method
func (m *MockSyntheticsAPI) UpdateCanaryWithContext(arg0 context.Context, arg1 *synthetics.UpdateCanaryInput, arg2 ...request.Option) (*synthetics.UpdateCanaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCanaryWithContext", varargs...)
	ret0, _ := ret[0].(*synthetics.UpdateCanaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCanaryWithContext indicates an expected call of UpdateCanaryWithContext
func (mr *MockSyntheticsAPIMockRecorder) UpdateCanaryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCanaryWithContext", reflect.TypeOf((*MockSyntheticsAPI)(nil).UpdateCanaryWithContext), varargs...)
}