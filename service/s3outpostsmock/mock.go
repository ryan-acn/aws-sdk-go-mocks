// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/s3outposts/s3outpostsiface (interfaces: S3OutpostsAPI)

// Package s3outpostsmock is a generated GoMock package.
package s3outpostsmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	s3outposts "github.com/aws/aws-sdk-go/service/s3outposts"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockS3OutpostsAPI is a mock of S3OutpostsAPI interface
type MockS3OutpostsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockS3OutpostsAPIMockRecorder
}

// MockS3OutpostsAPIMockRecorder is the mock recorder for MockS3OutpostsAPI
type MockS3OutpostsAPIMockRecorder struct {
	mock *MockS3OutpostsAPI
}

// NewMockS3OutpostsAPI creates a new mock instance
func NewMockS3OutpostsAPI(ctrl *gomock.Controller) *MockS3OutpostsAPI {
	mock := &MockS3OutpostsAPI{ctrl: ctrl}
	mock.recorder = &MockS3OutpostsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3OutpostsAPI) EXPECT() *MockS3OutpostsAPIMockRecorder {
	return m.recorder
}

// CreateEndpoint mocks base method
func (m *MockS3OutpostsAPI) CreateEndpoint(arg0 *s3outposts.CreateEndpointInput) (*s3outposts.CreateEndpointOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEndpoint", arg0)
	ret0, _ := ret[0].(*s3outposts.CreateEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEndpoint indicates an expected call of CreateEndpoint
func (mr *MockS3OutpostsAPIMockRecorder) CreateEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEndpoint", reflect.TypeOf((*MockS3OutpostsAPI)(nil).CreateEndpoint), arg0)
}

// CreateEndpointRequest mocks base method
func (m *MockS3OutpostsAPI) CreateEndpointRequest(arg0 *s3outposts.CreateEndpointInput) (*request.Request, *s3outposts.CreateEndpointOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEndpointRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3outposts.CreateEndpointOutput)
	return ret0, ret1
}

// CreateEndpointRequest indicates an expected call of CreateEndpointRequest
func (mr *MockS3OutpostsAPIMockRecorder) CreateEndpointRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEndpointRequest", reflect.TypeOf((*MockS3OutpostsAPI)(nil).CreateEndpointRequest), arg0)
}

// CreateEndpointWithContext mocks base method
func (m *MockS3OutpostsAPI) CreateEndpointWithContext(arg0 context.Context, arg1 *s3outposts.CreateEndpointInput, arg2 ...request.Option) (*s3outposts.CreateEndpointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEndpointWithContext", varargs...)
	ret0, _ := ret[0].(*s3outposts.CreateEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEndpointWithContext indicates an expected call of CreateEndpointWithContext
func (mr *MockS3OutpostsAPIMockRecorder) CreateEndpointWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEndpointWithContext", reflect.TypeOf((*MockS3OutpostsAPI)(nil).CreateEndpointWithContext), varargs...)
}

// DeleteEndpoint mocks base method
func (m *MockS3OutpostsAPI) DeleteEndpoint(arg0 *s3outposts.DeleteEndpointInput) (*s3outposts.DeleteEndpointOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEndpoint", arg0)
	ret0, _ := ret[0].(*s3outposts.DeleteEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEndpoint indicates an expected call of DeleteEndpoint
func (mr *MockS3OutpostsAPIMockRecorder) DeleteEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEndpoint", reflect.TypeOf((*MockS3OutpostsAPI)(nil).DeleteEndpoint), arg0)
}

// DeleteEndpointRequest mocks base method
func (m *MockS3OutpostsAPI) DeleteEndpointRequest(arg0 *s3outposts.DeleteEndpointInput) (*request.Request, *s3outposts.DeleteEndpointOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEndpointRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3outposts.DeleteEndpointOutput)
	return ret0, ret1
}

// DeleteEndpointRequest indicates an expected call of DeleteEndpointRequest
func (mr *MockS3OutpostsAPIMockRecorder) DeleteEndpointRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEndpointRequest", reflect.TypeOf((*MockS3OutpostsAPI)(nil).DeleteEndpointRequest), arg0)
}

// DeleteEndpointWithContext mocks base method
func (m *MockS3OutpostsAPI) DeleteEndpointWithContext(arg0 context.Context, arg1 *s3outposts.DeleteEndpointInput, arg2 ...request.Option) (*s3outposts.DeleteEndpointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEndpointWithContext", varargs...)
	ret0, _ := ret[0].(*s3outposts.DeleteEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEndpointWithContext indicates an expected call of DeleteEndpointWithContext
func (mr *MockS3OutpostsAPIMockRecorder) DeleteEndpointWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEndpointWithContext", reflect.TypeOf((*MockS3OutpostsAPI)(nil).DeleteEndpointWithContext), varargs...)
}

// ListEndpoints mocks base method
func (m *MockS3OutpostsAPI) ListEndpoints(arg0 *s3outposts.ListEndpointsInput) (*s3outposts.ListEndpointsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEndpoints", arg0)
	ret0, _ := ret[0].(*s3outposts.ListEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEndpoints indicates an expected call of ListEndpoints
func (mr *MockS3OutpostsAPIMockRecorder) ListEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpoints", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListEndpoints), arg0)
}

// ListEndpointsPages mocks base method
func (m *MockS3OutpostsAPI) ListEndpointsPages(arg0 *s3outposts.ListEndpointsInput, arg1 func(*s3outposts.ListEndpointsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEndpointsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListEndpointsPages indicates an expected call of ListEndpointsPages
func (mr *MockS3OutpostsAPIMockRecorder) ListEndpointsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsPages", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListEndpointsPages), arg0, arg1)
}

// ListEndpointsPagesWithContext mocks base method
func (m *MockS3OutpostsAPI) ListEndpointsPagesWithContext(arg0 context.Context, arg1 *s3outposts.ListEndpointsInput, arg2 func(*s3outposts.ListEndpointsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEndpointsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListEndpointsPagesWithContext indicates an expected call of ListEndpointsPagesWithContext
func (mr *MockS3OutpostsAPIMockRecorder) ListEndpointsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsPagesWithContext", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListEndpointsPagesWithContext), varargs...)
}

// ListEndpointsRequest mocks base method
func (m *MockS3OutpostsAPI) ListEndpointsRequest(arg0 *s3outposts.ListEndpointsInput) (*request.Request, *s3outposts.ListEndpointsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEndpointsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3outposts.ListEndpointsOutput)
	return ret0, ret1
}

// ListEndpointsRequest indicates an expected call of ListEndpointsRequest
func (mr *MockS3OutpostsAPIMockRecorder) ListEndpointsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsRequest", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListEndpointsRequest), arg0)
}

// ListEndpointsWithContext mocks base method
func (m *MockS3OutpostsAPI) ListEndpointsWithContext(arg0 context.Context, arg1 *s3outposts.ListEndpointsInput, arg2 ...request.Option) (*s3outposts.ListEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEndpointsWithContext", varargs...)
	ret0, _ := ret[0].(*s3outposts.ListEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEndpointsWithContext indicates an expected call of ListEndpointsWithContext
func (mr *MockS3OutpostsAPIMockRecorder) ListEndpointsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsWithContext", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListEndpointsWithContext), varargs...)
}

// ListOutpostsWithS3 mocks base method
func (m *MockS3OutpostsAPI) ListOutpostsWithS3(arg0 *s3outposts.ListOutpostsWithS3Input) (*s3outposts.ListOutpostsWithS3Output, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutpostsWithS3", arg0)
	ret0, _ := ret[0].(*s3outposts.ListOutpostsWithS3Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOutpostsWithS3 indicates an expected call of ListOutpostsWithS3
func (mr *MockS3OutpostsAPIMockRecorder) ListOutpostsWithS3(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutpostsWithS3", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListOutpostsWithS3), arg0)
}

// ListOutpostsWithS3Pages mocks base method
func (m *MockS3OutpostsAPI) ListOutpostsWithS3Pages(arg0 *s3outposts.ListOutpostsWithS3Input, arg1 func(*s3outposts.ListOutpostsWithS3Output, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutpostsWithS3Pages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListOutpostsWithS3Pages indicates an expected call of ListOutpostsWithS3Pages
func (mr *MockS3OutpostsAPIMockRecorder) ListOutpostsWithS3Pages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutpostsWithS3Pages", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListOutpostsWithS3Pages), arg0, arg1)
}

// ListOutpostsWithS3PagesWithContext mocks base method
func (m *MockS3OutpostsAPI) ListOutpostsWithS3PagesWithContext(arg0 context.Context, arg1 *s3outposts.ListOutpostsWithS3Input, arg2 func(*s3outposts.ListOutpostsWithS3Output, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOutpostsWithS3PagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListOutpostsWithS3PagesWithContext indicates an expected call of ListOutpostsWithS3PagesWithContext
func (mr *MockS3OutpostsAPIMockRecorder) ListOutpostsWithS3PagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutpostsWithS3PagesWithContext", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListOutpostsWithS3PagesWithContext), varargs...)
}

// ListOutpostsWithS3Request mocks base method
func (m *MockS3OutpostsAPI) ListOutpostsWithS3Request(arg0 *s3outposts.ListOutpostsWithS3Input) (*request.Request, *s3outposts.ListOutpostsWithS3Output) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutpostsWithS3Request", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3outposts.ListOutpostsWithS3Output)
	return ret0, ret1
}

// ListOutpostsWithS3Request indicates an expected call of ListOutpostsWithS3Request
func (mr *MockS3OutpostsAPIMockRecorder) ListOutpostsWithS3Request(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutpostsWithS3Request", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListOutpostsWithS3Request), arg0)
}

// ListOutpostsWithS3WithContext mocks base method
func (m *MockS3OutpostsAPI) ListOutpostsWithS3WithContext(arg0 context.Context, arg1 *s3outposts.ListOutpostsWithS3Input, arg2 ...request.Option) (*s3outposts.ListOutpostsWithS3Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOutpostsWithS3WithContext", varargs...)
	ret0, _ := ret[0].(*s3outposts.ListOutpostsWithS3Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOutpostsWithS3WithContext indicates an expected call of ListOutpostsWithS3WithContext
func (mr *MockS3OutpostsAPIMockRecorder) ListOutpostsWithS3WithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutpostsWithS3WithContext", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListOutpostsWithS3WithContext), varargs...)
}

// ListSharedEndpoints mocks base method
func (m *MockS3OutpostsAPI) ListSharedEndpoints(arg0 *s3outposts.ListSharedEndpointsInput) (*s3outposts.ListSharedEndpointsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSharedEndpoints", arg0)
	ret0, _ := ret[0].(*s3outposts.ListSharedEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSharedEndpoints indicates an expected call of ListSharedEndpoints
func (mr *MockS3OutpostsAPIMockRecorder) ListSharedEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSharedEndpoints", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListSharedEndpoints), arg0)
}

// ListSharedEndpointsPages mocks base method
func (m *MockS3OutpostsAPI) ListSharedEndpointsPages(arg0 *s3outposts.ListSharedEndpointsInput, arg1 func(*s3outposts.ListSharedEndpointsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSharedEndpointsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListSharedEndpointsPages indicates an expected call of ListSharedEndpointsPages
func (mr *MockS3OutpostsAPIMockRecorder) ListSharedEndpointsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSharedEndpointsPages", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListSharedEndpointsPages), arg0, arg1)
}

// ListSharedEndpointsPagesWithContext mocks base method
func (m *MockS3OutpostsAPI) ListSharedEndpointsPagesWithContext(arg0 context.Context, arg1 *s3outposts.ListSharedEndpointsInput, arg2 func(*s3outposts.ListSharedEndpointsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSharedEndpointsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListSharedEndpointsPagesWithContext indicates an expected call of ListSharedEndpointsPagesWithContext
func (mr *MockS3OutpostsAPIMockRecorder) ListSharedEndpointsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSharedEndpointsPagesWithContext", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListSharedEndpointsPagesWithContext), varargs...)
}

// ListSharedEndpointsRequest mocks base method
func (m *MockS3OutpostsAPI) ListSharedEndpointsRequest(arg0 *s3outposts.ListSharedEndpointsInput) (*request.Request, *s3outposts.ListSharedEndpointsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSharedEndpointsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3outposts.ListSharedEndpointsOutput)
	return ret0, ret1
}

// ListSharedEndpointsRequest indicates an expected call of ListSharedEndpointsRequest
func (mr *MockS3OutpostsAPIMockRecorder) ListSharedEndpointsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSharedEndpointsRequest", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListSharedEndpointsRequest), arg0)
}

// ListSharedEndpointsWithContext mocks base method
func (m *MockS3OutpostsAPI) ListSharedEndpointsWithContext(arg0 context.Context, arg1 *s3outposts.ListSharedEndpointsInput, arg2 ...request.Option) (*s3outposts.ListSharedEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSharedEndpointsWithContext", varargs...)
	ret0, _ := ret[0].(*s3outposts.ListSharedEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSharedEndpointsWithContext indicates an expected call of ListSharedEndpointsWithContext
func (mr *MockS3OutpostsAPIMockRecorder) ListSharedEndpointsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSharedEndpointsWithContext", reflect.TypeOf((*MockS3OutpostsAPI)(nil).ListSharedEndpointsWithContext), varargs...)
}
