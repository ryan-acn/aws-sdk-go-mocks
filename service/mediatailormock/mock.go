// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/mediatailor/mediatailoriface (interfaces: MediaTailorAPI)

// Package mediatailormock is a generated GoMock package.
package mediatailormock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	mediatailor "github.com/aws/aws-sdk-go/service/mediatailor"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMediaTailorAPI is a mock of MediaTailorAPI interface
type MockMediaTailorAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMediaTailorAPIMockRecorder
}

// MockMediaTailorAPIMockRecorder is the mock recorder for MockMediaTailorAPI
type MockMediaTailorAPIMockRecorder struct {
	mock *MockMediaTailorAPI
}

// NewMockMediaTailorAPI creates a new mock instance
func NewMockMediaTailorAPI(ctrl *gomock.Controller) *MockMediaTailorAPI {
	mock := &MockMediaTailorAPI{ctrl: ctrl}
	mock.recorder = &MockMediaTailorAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMediaTailorAPI) EXPECT() *MockMediaTailorAPIMockRecorder {
	return m.recorder
}

// DeletePlaybackConfiguration mocks base method
func (m *MockMediaTailorAPI) DeletePlaybackConfiguration(arg0 *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaybackConfiguration", arg0)
	ret0, _ := ret[0].(*mediatailor.DeletePlaybackConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePlaybackConfiguration indicates an expected call of DeletePlaybackConfiguration
func (mr *MockMediaTailorAPIMockRecorder) DeletePlaybackConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaybackConfiguration", reflect.TypeOf((*MockMediaTailorAPI)(nil).DeletePlaybackConfiguration), arg0)
}

// DeletePlaybackConfigurationRequest mocks base method
func (m *MockMediaTailorAPI) DeletePlaybackConfigurationRequest(arg0 *mediatailor.DeletePlaybackConfigurationInput) (*request.Request, *mediatailor.DeletePlaybackConfigurationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaybackConfigurationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediatailor.DeletePlaybackConfigurationOutput)
	return ret0, ret1
}

// DeletePlaybackConfigurationRequest indicates an expected call of DeletePlaybackConfigurationRequest
func (mr *MockMediaTailorAPIMockRecorder) DeletePlaybackConfigurationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaybackConfigurationRequest", reflect.TypeOf((*MockMediaTailorAPI)(nil).DeletePlaybackConfigurationRequest), arg0)
}

// DeletePlaybackConfigurationWithContext mocks base method
func (m *MockMediaTailorAPI) DeletePlaybackConfigurationWithContext(arg0 context.Context, arg1 *mediatailor.DeletePlaybackConfigurationInput, arg2 ...request.Option) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePlaybackConfigurationWithContext", varargs...)
	ret0, _ := ret[0].(*mediatailor.DeletePlaybackConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePlaybackConfigurationWithContext indicates an expected call of DeletePlaybackConfigurationWithContext
func (mr *MockMediaTailorAPIMockRecorder) DeletePlaybackConfigurationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaybackConfigurationWithContext", reflect.TypeOf((*MockMediaTailorAPI)(nil).DeletePlaybackConfigurationWithContext), varargs...)
}

// GetPlaybackConfiguration mocks base method
func (m *MockMediaTailorAPI) GetPlaybackConfiguration(arg0 *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaybackConfiguration", arg0)
	ret0, _ := ret[0].(*mediatailor.GetPlaybackConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaybackConfiguration indicates an expected call of GetPlaybackConfiguration
func (mr *MockMediaTailorAPIMockRecorder) GetPlaybackConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaybackConfiguration", reflect.TypeOf((*MockMediaTailorAPI)(nil).GetPlaybackConfiguration), arg0)
}

// GetPlaybackConfigurationRequest mocks base method
func (m *MockMediaTailorAPI) GetPlaybackConfigurationRequest(arg0 *mediatailor.GetPlaybackConfigurationInput) (*request.Request, *mediatailor.GetPlaybackConfigurationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaybackConfigurationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediatailor.GetPlaybackConfigurationOutput)
	return ret0, ret1
}

// GetPlaybackConfigurationRequest indicates an expected call of GetPlaybackConfigurationRequest
func (mr *MockMediaTailorAPIMockRecorder) GetPlaybackConfigurationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaybackConfigurationRequest", reflect.TypeOf((*MockMediaTailorAPI)(nil).GetPlaybackConfigurationRequest), arg0)
}

// GetPlaybackConfigurationWithContext mocks base method
func (m *MockMediaTailorAPI) GetPlaybackConfigurationWithContext(arg0 context.Context, arg1 *mediatailor.GetPlaybackConfigurationInput, arg2 ...request.Option) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPlaybackConfigurationWithContext", varargs...)
	ret0, _ := ret[0].(*mediatailor.GetPlaybackConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaybackConfigurationWithContext indicates an expected call of GetPlaybackConfigurationWithContext
func (mr *MockMediaTailorAPIMockRecorder) GetPlaybackConfigurationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaybackConfigurationWithContext", reflect.TypeOf((*MockMediaTailorAPI)(nil).GetPlaybackConfigurationWithContext), varargs...)
}

// ListPlaybackConfigurations mocks base method
func (m *MockMediaTailorAPI) ListPlaybackConfigurations(arg0 *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaybackConfigurations", arg0)
	ret0, _ := ret[0].(*mediatailor.ListPlaybackConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaybackConfigurations indicates an expected call of ListPlaybackConfigurations
func (mr *MockMediaTailorAPIMockRecorder) ListPlaybackConfigurations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaybackConfigurations", reflect.TypeOf((*MockMediaTailorAPI)(nil).ListPlaybackConfigurations), arg0)
}

// ListPlaybackConfigurationsRequest mocks base method
func (m *MockMediaTailorAPI) ListPlaybackConfigurationsRequest(arg0 *mediatailor.ListPlaybackConfigurationsInput) (*request.Request, *mediatailor.ListPlaybackConfigurationsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaybackConfigurationsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediatailor.ListPlaybackConfigurationsOutput)
	return ret0, ret1
}

// ListPlaybackConfigurationsRequest indicates an expected call of ListPlaybackConfigurationsRequest
func (mr *MockMediaTailorAPIMockRecorder) ListPlaybackConfigurationsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaybackConfigurationsRequest", reflect.TypeOf((*MockMediaTailorAPI)(nil).ListPlaybackConfigurationsRequest), arg0)
}

// ListPlaybackConfigurationsWithContext mocks base method
func (m *MockMediaTailorAPI) ListPlaybackConfigurationsWithContext(arg0 context.Context, arg1 *mediatailor.ListPlaybackConfigurationsInput, arg2 ...request.Option) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPlaybackConfigurationsWithContext", varargs...)
	ret0, _ := ret[0].(*mediatailor.ListPlaybackConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaybackConfigurationsWithContext indicates an expected call of ListPlaybackConfigurationsWithContext
func (mr *MockMediaTailorAPIMockRecorder) ListPlaybackConfigurationsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaybackConfigurationsWithContext", reflect.TypeOf((*MockMediaTailorAPI)(nil).ListPlaybackConfigurationsWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockMediaTailorAPI) ListTagsForResource(arg0 *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*mediatailor.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockMediaTailorAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockMediaTailorAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockMediaTailorAPI) ListTagsForResourceRequest(arg0 *mediatailor.ListTagsForResourceInput) (*request.Request, *mediatailor.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediatailor.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockMediaTailorAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockMediaTailorAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockMediaTailorAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *mediatailor.ListTagsForResourceInput, arg2 ...request.Option) (*mediatailor.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*mediatailor.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockMediaTailorAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockMediaTailorAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// PutPlaybackConfiguration mocks base method
func (m *MockMediaTailorAPI) PutPlaybackConfiguration(arg0 *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPlaybackConfiguration", arg0)
	ret0, _ := ret[0].(*mediatailor.PutPlaybackConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutPlaybackConfiguration indicates an expected call of PutPlaybackConfiguration
func (mr *MockMediaTailorAPIMockRecorder) PutPlaybackConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPlaybackConfiguration", reflect.TypeOf((*MockMediaTailorAPI)(nil).PutPlaybackConfiguration), arg0)
}

// PutPlaybackConfigurationRequest mocks base method
func (m *MockMediaTailorAPI) PutPlaybackConfigurationRequest(arg0 *mediatailor.PutPlaybackConfigurationInput) (*request.Request, *mediatailor.PutPlaybackConfigurationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPlaybackConfigurationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediatailor.PutPlaybackConfigurationOutput)
	return ret0, ret1
}

// PutPlaybackConfigurationRequest indicates an expected call of PutPlaybackConfigurationRequest
func (mr *MockMediaTailorAPIMockRecorder) PutPlaybackConfigurationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPlaybackConfigurationRequest", reflect.TypeOf((*MockMediaTailorAPI)(nil).PutPlaybackConfigurationRequest), arg0)
}

// PutPlaybackConfigurationWithContext mocks base method
func (m *MockMediaTailorAPI) PutPlaybackConfigurationWithContext(arg0 context.Context, arg1 *mediatailor.PutPlaybackConfigurationInput, arg2 ...request.Option) (*mediatailor.PutPlaybackConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutPlaybackConfigurationWithContext", varargs...)
	ret0, _ := ret[0].(*mediatailor.PutPlaybackConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutPlaybackConfigurationWithContext indicates an expected call of PutPlaybackConfigurationWithContext
func (mr *MockMediaTailorAPIMockRecorder) PutPlaybackConfigurationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPlaybackConfigurationWithContext", reflect.TypeOf((*MockMediaTailorAPI)(nil).PutPlaybackConfigurationWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockMediaTailorAPI) TagResource(arg0 *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*mediatailor.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockMediaTailorAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockMediaTailorAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockMediaTailorAPI) TagResourceRequest(arg0 *mediatailor.TagResourceInput) (*request.Request, *mediatailor.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediatailor.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockMediaTailorAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockMediaTailorAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockMediaTailorAPI) TagResourceWithContext(arg0 context.Context, arg1 *mediatailor.TagResourceInput, arg2 ...request.Option) (*mediatailor.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*mediatailor.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockMediaTailorAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockMediaTailorAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockMediaTailorAPI) UntagResource(arg0 *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*mediatailor.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockMediaTailorAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockMediaTailorAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockMediaTailorAPI) UntagResourceRequest(arg0 *mediatailor.UntagResourceInput) (*request.Request, *mediatailor.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediatailor.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockMediaTailorAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockMediaTailorAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockMediaTailorAPI) UntagResourceWithContext(arg0 context.Context, arg1 *mediatailor.UntagResourceInput, arg2 ...request.Option) (*mediatailor.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*mediatailor.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockMediaTailorAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockMediaTailorAPI)(nil).UntagResourceWithContext), varargs...)
}