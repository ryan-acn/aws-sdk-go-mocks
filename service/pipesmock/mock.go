// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/pipes/pipesiface (interfaces: PipesAPI)

// Package pipesmock is a generated GoMock package.
package pipesmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	pipes "github.com/aws/aws-sdk-go/service/pipes"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPipesAPI is a mock of PipesAPI interface
type MockPipesAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPipesAPIMockRecorder
}

// MockPipesAPIMockRecorder is the mock recorder for MockPipesAPI
type MockPipesAPIMockRecorder struct {
	mock *MockPipesAPI
}

// NewMockPipesAPI creates a new mock instance
func NewMockPipesAPI(ctrl *gomock.Controller) *MockPipesAPI {
	mock := &MockPipesAPI{ctrl: ctrl}
	mock.recorder = &MockPipesAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPipesAPI) EXPECT() *MockPipesAPIMockRecorder {
	return m.recorder
}

// CreatePipe mocks base method
func (m *MockPipesAPI) CreatePipe(arg0 *pipes.CreatePipeInput) (*pipes.CreatePipeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipe", arg0)
	ret0, _ := ret[0].(*pipes.CreatePipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePipe indicates an expected call of CreatePipe
func (mr *MockPipesAPIMockRecorder) CreatePipe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipe", reflect.TypeOf((*MockPipesAPI)(nil).CreatePipe), arg0)
}

// CreatePipeRequest mocks base method
func (m *MockPipesAPI) CreatePipeRequest(arg0 *pipes.CreatePipeInput) (*request.Request, *pipes.CreatePipeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.CreatePipeOutput)
	return ret0, ret1
}

// CreatePipeRequest indicates an expected call of CreatePipeRequest
func (mr *MockPipesAPIMockRecorder) CreatePipeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeRequest", reflect.TypeOf((*MockPipesAPI)(nil).CreatePipeRequest), arg0)
}

// CreatePipeWithContext mocks base method
func (m *MockPipesAPI) CreatePipeWithContext(arg0 context.Context, arg1 *pipes.CreatePipeInput, arg2 ...request.Option) (*pipes.CreatePipeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePipeWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.CreatePipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePipeWithContext indicates an expected call of CreatePipeWithContext
func (mr *MockPipesAPIMockRecorder) CreatePipeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeWithContext", reflect.TypeOf((*MockPipesAPI)(nil).CreatePipeWithContext), varargs...)
}

// DeletePipe mocks base method
func (m *MockPipesAPI) DeletePipe(arg0 *pipes.DeletePipeInput) (*pipes.DeletePipeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipe", arg0)
	ret0, _ := ret[0].(*pipes.DeletePipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePipe indicates an expected call of DeletePipe
func (mr *MockPipesAPIMockRecorder) DeletePipe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipe", reflect.TypeOf((*MockPipesAPI)(nil).DeletePipe), arg0)
}

// DeletePipeRequest mocks base method
func (m *MockPipesAPI) DeletePipeRequest(arg0 *pipes.DeletePipeInput) (*request.Request, *pipes.DeletePipeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.DeletePipeOutput)
	return ret0, ret1
}

// DeletePipeRequest indicates an expected call of DeletePipeRequest
func (mr *MockPipesAPIMockRecorder) DeletePipeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeRequest", reflect.TypeOf((*MockPipesAPI)(nil).DeletePipeRequest), arg0)
}

// DeletePipeWithContext mocks base method
func (m *MockPipesAPI) DeletePipeWithContext(arg0 context.Context, arg1 *pipes.DeletePipeInput, arg2 ...request.Option) (*pipes.DeletePipeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePipeWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.DeletePipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePipeWithContext indicates an expected call of DeletePipeWithContext
func (mr *MockPipesAPIMockRecorder) DeletePipeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeWithContext", reflect.TypeOf((*MockPipesAPI)(nil).DeletePipeWithContext), varargs...)
}

// DescribePipe mocks base method
func (m *MockPipesAPI) DescribePipe(arg0 *pipes.DescribePipeInput) (*pipes.DescribePipeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribePipe", arg0)
	ret0, _ := ret[0].(*pipes.DescribePipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePipe indicates an expected call of DescribePipe
func (mr *MockPipesAPIMockRecorder) DescribePipe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePipe", reflect.TypeOf((*MockPipesAPI)(nil).DescribePipe), arg0)
}

// DescribePipeRequest mocks base method
func (m *MockPipesAPI) DescribePipeRequest(arg0 *pipes.DescribePipeInput) (*request.Request, *pipes.DescribePipeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribePipeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.DescribePipeOutput)
	return ret0, ret1
}

// DescribePipeRequest indicates an expected call of DescribePipeRequest
func (mr *MockPipesAPIMockRecorder) DescribePipeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePipeRequest", reflect.TypeOf((*MockPipesAPI)(nil).DescribePipeRequest), arg0)
}

// DescribePipeWithContext mocks base method
func (m *MockPipesAPI) DescribePipeWithContext(arg0 context.Context, arg1 *pipes.DescribePipeInput, arg2 ...request.Option) (*pipes.DescribePipeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePipeWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.DescribePipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePipeWithContext indicates an expected call of DescribePipeWithContext
func (mr *MockPipesAPIMockRecorder) DescribePipeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePipeWithContext", reflect.TypeOf((*MockPipesAPI)(nil).DescribePipeWithContext), varargs...)
}

// ListPipes mocks base method
func (m *MockPipesAPI) ListPipes(arg0 *pipes.ListPipesInput) (*pipes.ListPipesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipes", arg0)
	ret0, _ := ret[0].(*pipes.ListPipesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPipes indicates an expected call of ListPipes
func (mr *MockPipesAPIMockRecorder) ListPipes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipes", reflect.TypeOf((*MockPipesAPI)(nil).ListPipes), arg0)
}

// ListPipesPages mocks base method
func (m *MockPipesAPI) ListPipesPages(arg0 *pipes.ListPipesInput, arg1 func(*pipes.ListPipesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListPipesPages indicates an expected call of ListPipesPages
func (mr *MockPipesAPIMockRecorder) ListPipesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipesPages", reflect.TypeOf((*MockPipesAPI)(nil).ListPipesPages), arg0, arg1)
}

// ListPipesPagesWithContext mocks base method
func (m *MockPipesAPI) ListPipesPagesWithContext(arg0 context.Context, arg1 *pipes.ListPipesInput, arg2 func(*pipes.ListPipesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPipesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListPipesPagesWithContext indicates an expected call of ListPipesPagesWithContext
func (mr *MockPipesAPIMockRecorder) ListPipesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipesPagesWithContext", reflect.TypeOf((*MockPipesAPI)(nil).ListPipesPagesWithContext), varargs...)
}

// ListPipesRequest mocks base method
func (m *MockPipesAPI) ListPipesRequest(arg0 *pipes.ListPipesInput) (*request.Request, *pipes.ListPipesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.ListPipesOutput)
	return ret0, ret1
}

// ListPipesRequest indicates an expected call of ListPipesRequest
func (mr *MockPipesAPIMockRecorder) ListPipesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipesRequest", reflect.TypeOf((*MockPipesAPI)(nil).ListPipesRequest), arg0)
}

// ListPipesWithContext mocks base method
func (m *MockPipesAPI) ListPipesWithContext(arg0 context.Context, arg1 *pipes.ListPipesInput, arg2 ...request.Option) (*pipes.ListPipesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPipesWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.ListPipesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPipesWithContext indicates an expected call of ListPipesWithContext
func (mr *MockPipesAPIMockRecorder) ListPipesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipesWithContext", reflect.TypeOf((*MockPipesAPI)(nil).ListPipesWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockPipesAPI) ListTagsForResource(arg0 *pipes.ListTagsForResourceInput) (*pipes.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*pipes.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockPipesAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockPipesAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockPipesAPI) ListTagsForResourceRequest(arg0 *pipes.ListTagsForResourceInput) (*request.Request, *pipes.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockPipesAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockPipesAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockPipesAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *pipes.ListTagsForResourceInput, arg2 ...request.Option) (*pipes.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockPipesAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockPipesAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// StartPipe mocks base method
func (m *MockPipesAPI) StartPipe(arg0 *pipes.StartPipeInput) (*pipes.StartPipeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartPipe", arg0)
	ret0, _ := ret[0].(*pipes.StartPipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartPipe indicates an expected call of StartPipe
func (mr *MockPipesAPIMockRecorder) StartPipe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPipe", reflect.TypeOf((*MockPipesAPI)(nil).StartPipe), arg0)
}

// StartPipeRequest mocks base method
func (m *MockPipesAPI) StartPipeRequest(arg0 *pipes.StartPipeInput) (*request.Request, *pipes.StartPipeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartPipeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.StartPipeOutput)
	return ret0, ret1
}

// StartPipeRequest indicates an expected call of StartPipeRequest
func (mr *MockPipesAPIMockRecorder) StartPipeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPipeRequest", reflect.TypeOf((*MockPipesAPI)(nil).StartPipeRequest), arg0)
}

// StartPipeWithContext mocks base method
func (m *MockPipesAPI) StartPipeWithContext(arg0 context.Context, arg1 *pipes.StartPipeInput, arg2 ...request.Option) (*pipes.StartPipeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartPipeWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.StartPipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartPipeWithContext indicates an expected call of StartPipeWithContext
func (mr *MockPipesAPIMockRecorder) StartPipeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPipeWithContext", reflect.TypeOf((*MockPipesAPI)(nil).StartPipeWithContext), varargs...)
}

// StopPipe mocks base method
func (m *MockPipesAPI) StopPipe(arg0 *pipes.StopPipeInput) (*pipes.StopPipeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopPipe", arg0)
	ret0, _ := ret[0].(*pipes.StopPipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopPipe indicates an expected call of StopPipe
func (mr *MockPipesAPIMockRecorder) StopPipe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPipe", reflect.TypeOf((*MockPipesAPI)(nil).StopPipe), arg0)
}

// StopPipeRequest mocks base method
func (m *MockPipesAPI) StopPipeRequest(arg0 *pipes.StopPipeInput) (*request.Request, *pipes.StopPipeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopPipeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.StopPipeOutput)
	return ret0, ret1
}

// StopPipeRequest indicates an expected call of StopPipeRequest
func (mr *MockPipesAPIMockRecorder) StopPipeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPipeRequest", reflect.TypeOf((*MockPipesAPI)(nil).StopPipeRequest), arg0)
}

// StopPipeWithContext mocks base method
func (m *MockPipesAPI) StopPipeWithContext(arg0 context.Context, arg1 *pipes.StopPipeInput, arg2 ...request.Option) (*pipes.StopPipeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopPipeWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.StopPipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopPipeWithContext indicates an expected call of StopPipeWithContext
func (mr *MockPipesAPIMockRecorder) StopPipeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPipeWithContext", reflect.TypeOf((*MockPipesAPI)(nil).StopPipeWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockPipesAPI) TagResource(arg0 *pipes.TagResourceInput) (*pipes.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*pipes.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockPipesAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockPipesAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockPipesAPI) TagResourceRequest(arg0 *pipes.TagResourceInput) (*request.Request, *pipes.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockPipesAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockPipesAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockPipesAPI) TagResourceWithContext(arg0 context.Context, arg1 *pipes.TagResourceInput, arg2 ...request.Option) (*pipes.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockPipesAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockPipesAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockPipesAPI) UntagResource(arg0 *pipes.UntagResourceInput) (*pipes.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*pipes.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockPipesAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockPipesAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockPipesAPI) UntagResourceRequest(arg0 *pipes.UntagResourceInput) (*request.Request, *pipes.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockPipesAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockPipesAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockPipesAPI) UntagResourceWithContext(arg0 context.Context, arg1 *pipes.UntagResourceInput, arg2 ...request.Option) (*pipes.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockPipesAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockPipesAPI)(nil).UntagResourceWithContext), varargs...)
}

// UpdatePipe mocks base method
func (m *MockPipesAPI) UpdatePipe(arg0 *pipes.UpdatePipeInput) (*pipes.UpdatePipeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipe", arg0)
	ret0, _ := ret[0].(*pipes.UpdatePipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePipe indicates an expected call of UpdatePipe
func (mr *MockPipesAPIMockRecorder) UpdatePipe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipe", reflect.TypeOf((*MockPipesAPI)(nil).UpdatePipe), arg0)
}

// UpdatePipeRequest mocks base method
func (m *MockPipesAPI) UpdatePipeRequest(arg0 *pipes.UpdatePipeInput) (*request.Request, *pipes.UpdatePipeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pipes.UpdatePipeOutput)
	return ret0, ret1
}

// UpdatePipeRequest indicates an expected call of UpdatePipeRequest
func (mr *MockPipesAPIMockRecorder) UpdatePipeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeRequest", reflect.TypeOf((*MockPipesAPI)(nil).UpdatePipeRequest), arg0)
}

// UpdatePipeWithContext mocks base method
func (m *MockPipesAPI) UpdatePipeWithContext(arg0 context.Context, arg1 *pipes.UpdatePipeInput, arg2 ...request.Option) (*pipes.UpdatePipeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePipeWithContext", varargs...)
	ret0, _ := ret[0].(*pipes.UpdatePipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePipeWithContext indicates an expected call of UpdatePipeWithContext
func (mr *MockPipesAPIMockRecorder) UpdatePipeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeWithContext", reflect.TypeOf((*MockPipesAPI)(nil).UpdatePipeWithContext), varargs...)
}
