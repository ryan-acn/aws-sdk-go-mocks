// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/elasticinference/elasticinferenceiface (interfaces: ElasticInferenceAPI)

// Package elasticinferencemock is a generated GoMock package.
package elasticinferencemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	elasticinference "github.com/aws/aws-sdk-go/service/elasticinference"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockElasticInferenceAPI is a mock of ElasticInferenceAPI interface
type MockElasticInferenceAPI struct {
	ctrl     *gomock.Controller
	recorder *MockElasticInferenceAPIMockRecorder
}

// MockElasticInferenceAPIMockRecorder is the mock recorder for MockElasticInferenceAPI
type MockElasticInferenceAPIMockRecorder struct {
	mock *MockElasticInferenceAPI
}

// NewMockElasticInferenceAPI creates a new mock instance
func NewMockElasticInferenceAPI(ctrl *gomock.Controller) *MockElasticInferenceAPI {
	mock := &MockElasticInferenceAPI{ctrl: ctrl}
	mock.recorder = &MockElasticInferenceAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockElasticInferenceAPI) EXPECT() *MockElasticInferenceAPIMockRecorder {
	return m.recorder
}

// ListTagsForResource mocks base method
func (m *MockElasticInferenceAPI) ListTagsForResource(arg0 *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*elasticinference.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockElasticInferenceAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockElasticInferenceAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockElasticInferenceAPI) ListTagsForResourceRequest(arg0 *elasticinference.ListTagsForResourceInput) (*request.Request, *elasticinference.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*elasticinference.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockElasticInferenceAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockElasticInferenceAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockElasticInferenceAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *elasticinference.ListTagsForResourceInput, arg2 ...request.Option) (*elasticinference.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*elasticinference.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockElasticInferenceAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockElasticInferenceAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockElasticInferenceAPI) TagResource(arg0 *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*elasticinference.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockElasticInferenceAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockElasticInferenceAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockElasticInferenceAPI) TagResourceRequest(arg0 *elasticinference.TagResourceInput) (*request.Request, *elasticinference.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*elasticinference.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockElasticInferenceAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockElasticInferenceAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockElasticInferenceAPI) TagResourceWithContext(arg0 context.Context, arg1 *elasticinference.TagResourceInput, arg2 ...request.Option) (*elasticinference.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*elasticinference.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockElasticInferenceAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockElasticInferenceAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockElasticInferenceAPI) UntagResource(arg0 *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*elasticinference.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockElasticInferenceAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockElasticInferenceAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockElasticInferenceAPI) UntagResourceRequest(arg0 *elasticinference.UntagResourceInput) (*request.Request, *elasticinference.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*elasticinference.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockElasticInferenceAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockElasticInferenceAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockElasticInferenceAPI) UntagResourceWithContext(arg0 context.Context, arg1 *elasticinference.UntagResourceInput, arg2 ...request.Option) (*elasticinference.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*elasticinference.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockElasticInferenceAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockElasticInferenceAPI)(nil).UntagResourceWithContext), varargs...)
}