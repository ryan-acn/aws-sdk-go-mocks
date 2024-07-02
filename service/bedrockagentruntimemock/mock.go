// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/bedrockagentruntime/bedrockagentruntimeiface (interfaces: BedrockAgentRuntimeAPI)

// Package bedrockagentruntimemock is a generated GoMock package.
package bedrockagentruntimemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	bedrockagentruntime "github.com/aws/aws-sdk-go/service/bedrockagentruntime"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBedrockAgentRuntimeAPI is a mock of BedrockAgentRuntimeAPI interface
type MockBedrockAgentRuntimeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBedrockAgentRuntimeAPIMockRecorder
}

// MockBedrockAgentRuntimeAPIMockRecorder is the mock recorder for MockBedrockAgentRuntimeAPI
type MockBedrockAgentRuntimeAPIMockRecorder struct {
	mock *MockBedrockAgentRuntimeAPI
}

// NewMockBedrockAgentRuntimeAPI creates a new mock instance
func NewMockBedrockAgentRuntimeAPI(ctrl *gomock.Controller) *MockBedrockAgentRuntimeAPI {
	mock := &MockBedrockAgentRuntimeAPI{ctrl: ctrl}
	mock.recorder = &MockBedrockAgentRuntimeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBedrockAgentRuntimeAPI) EXPECT() *MockBedrockAgentRuntimeAPIMockRecorder {
	return m.recorder
}

// InvokeAgent mocks base method
func (m *MockBedrockAgentRuntimeAPI) InvokeAgent(arg0 *bedrockagentruntime.InvokeAgentInput) (*bedrockagentruntime.InvokeAgentOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeAgent", arg0)
	ret0, _ := ret[0].(*bedrockagentruntime.InvokeAgentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeAgent indicates an expected call of InvokeAgent
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) InvokeAgent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeAgent", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).InvokeAgent), arg0)
}

// InvokeAgentRequest mocks base method
func (m *MockBedrockAgentRuntimeAPI) InvokeAgentRequest(arg0 *bedrockagentruntime.InvokeAgentInput) (*request.Request, *bedrockagentruntime.InvokeAgentOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeAgentRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*bedrockagentruntime.InvokeAgentOutput)
	return ret0, ret1
}

// InvokeAgentRequest indicates an expected call of InvokeAgentRequest
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) InvokeAgentRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeAgentRequest", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).InvokeAgentRequest), arg0)
}

// InvokeAgentWithContext mocks base method
func (m *MockBedrockAgentRuntimeAPI) InvokeAgentWithContext(arg0 context.Context, arg1 *bedrockagentruntime.InvokeAgentInput, arg2 ...request.Option) (*bedrockagentruntime.InvokeAgentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeAgentWithContext", varargs...)
	ret0, _ := ret[0].(*bedrockagentruntime.InvokeAgentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeAgentWithContext indicates an expected call of InvokeAgentWithContext
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) InvokeAgentWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeAgentWithContext", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).InvokeAgentWithContext), varargs...)
}

// Retrieve mocks base method
func (m *MockBedrockAgentRuntimeAPI) Retrieve(arg0 *bedrockagentruntime.RetrieveInput) (*bedrockagentruntime.RetrieveOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve", arg0)
	ret0, _ := ret[0].(*bedrockagentruntime.RetrieveOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) Retrieve(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).Retrieve), arg0)
}

// RetrieveAndGenerate mocks base method
func (m *MockBedrockAgentRuntimeAPI) RetrieveAndGenerate(arg0 *bedrockagentruntime.RetrieveAndGenerateInput) (*bedrockagentruntime.RetrieveAndGenerateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveAndGenerate", arg0)
	ret0, _ := ret[0].(*bedrockagentruntime.RetrieveAndGenerateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveAndGenerate indicates an expected call of RetrieveAndGenerate
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) RetrieveAndGenerate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAndGenerate", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).RetrieveAndGenerate), arg0)
}

// RetrieveAndGenerateRequest mocks base method
func (m *MockBedrockAgentRuntimeAPI) RetrieveAndGenerateRequest(arg0 *bedrockagentruntime.RetrieveAndGenerateInput) (*request.Request, *bedrockagentruntime.RetrieveAndGenerateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveAndGenerateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*bedrockagentruntime.RetrieveAndGenerateOutput)
	return ret0, ret1
}

// RetrieveAndGenerateRequest indicates an expected call of RetrieveAndGenerateRequest
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) RetrieveAndGenerateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAndGenerateRequest", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).RetrieveAndGenerateRequest), arg0)
}

// RetrieveAndGenerateWithContext mocks base method
func (m *MockBedrockAgentRuntimeAPI) RetrieveAndGenerateWithContext(arg0 context.Context, arg1 *bedrockagentruntime.RetrieveAndGenerateInput, arg2 ...request.Option) (*bedrockagentruntime.RetrieveAndGenerateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveAndGenerateWithContext", varargs...)
	ret0, _ := ret[0].(*bedrockagentruntime.RetrieveAndGenerateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveAndGenerateWithContext indicates an expected call of RetrieveAndGenerateWithContext
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) RetrieveAndGenerateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAndGenerateWithContext", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).RetrieveAndGenerateWithContext), varargs...)
}

// RetrievePages mocks base method
func (m *MockBedrockAgentRuntimeAPI) RetrievePages(arg0 *bedrockagentruntime.RetrieveInput, arg1 func(*bedrockagentruntime.RetrieveOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrievePages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RetrievePages indicates an expected call of RetrievePages
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) RetrievePages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrievePages", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).RetrievePages), arg0, arg1)
}

// RetrievePagesWithContext mocks base method
func (m *MockBedrockAgentRuntimeAPI) RetrievePagesWithContext(arg0 context.Context, arg1 *bedrockagentruntime.RetrieveInput, arg2 func(*bedrockagentruntime.RetrieveOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrievePagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RetrievePagesWithContext indicates an expected call of RetrievePagesWithContext
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) RetrievePagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrievePagesWithContext", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).RetrievePagesWithContext), varargs...)
}

// RetrieveRequest mocks base method
func (m *MockBedrockAgentRuntimeAPI) RetrieveRequest(arg0 *bedrockagentruntime.RetrieveInput) (*request.Request, *bedrockagentruntime.RetrieveOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*bedrockagentruntime.RetrieveOutput)
	return ret0, ret1
}

// RetrieveRequest indicates an expected call of RetrieveRequest
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) RetrieveRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveRequest", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).RetrieveRequest), arg0)
}

// RetrieveWithContext mocks base method
func (m *MockBedrockAgentRuntimeAPI) RetrieveWithContext(arg0 context.Context, arg1 *bedrockagentruntime.RetrieveInput, arg2 ...request.Option) (*bedrockagentruntime.RetrieveOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveWithContext", varargs...)
	ret0, _ := ret[0].(*bedrockagentruntime.RetrieveOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveWithContext indicates an expected call of RetrieveWithContext
func (mr *MockBedrockAgentRuntimeAPIMockRecorder) RetrieveWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveWithContext", reflect.TypeOf((*MockBedrockAgentRuntimeAPI)(nil).RetrieveWithContext), varargs...)
}