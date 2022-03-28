// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface (interfaces: WorkMailMessageFlowAPI)

// Package workmailmessageflowmock is a generated GoMock package.
package workmailmessageflowmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	workmailmessageflow "github.com/aws/aws-sdk-go/service/workmailmessageflow"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWorkMailMessageFlowAPI is a mock of WorkMailMessageFlowAPI interface
type MockWorkMailMessageFlowAPI struct {
	ctrl     *gomock.Controller
	recorder *MockWorkMailMessageFlowAPIMockRecorder
}

// MockWorkMailMessageFlowAPIMockRecorder is the mock recorder for MockWorkMailMessageFlowAPI
type MockWorkMailMessageFlowAPIMockRecorder struct {
	mock *MockWorkMailMessageFlowAPI
}

// NewMockWorkMailMessageFlowAPI creates a new mock instance
func NewMockWorkMailMessageFlowAPI(ctrl *gomock.Controller) *MockWorkMailMessageFlowAPI {
	mock := &MockWorkMailMessageFlowAPI{ctrl: ctrl}
	mock.recorder = &MockWorkMailMessageFlowAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkMailMessageFlowAPI) EXPECT() *MockWorkMailMessageFlowAPIMockRecorder {
	return m.recorder
}

// GetRawMessageContent mocks base method
func (m *MockWorkMailMessageFlowAPI) GetRawMessageContent(arg0 *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawMessageContent", arg0)
	ret0, _ := ret[0].(*workmailmessageflow.GetRawMessageContentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawMessageContent indicates an expected call of GetRawMessageContent
func (mr *MockWorkMailMessageFlowAPIMockRecorder) GetRawMessageContent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawMessageContent", reflect.TypeOf((*MockWorkMailMessageFlowAPI)(nil).GetRawMessageContent), arg0)
}

// GetRawMessageContentRequest mocks base method
func (m *MockWorkMailMessageFlowAPI) GetRawMessageContentRequest(arg0 *workmailmessageflow.GetRawMessageContentInput) (*request.Request, *workmailmessageflow.GetRawMessageContentOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawMessageContentRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*workmailmessageflow.GetRawMessageContentOutput)
	return ret0, ret1
}

// GetRawMessageContentRequest indicates an expected call of GetRawMessageContentRequest
func (mr *MockWorkMailMessageFlowAPIMockRecorder) GetRawMessageContentRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawMessageContentRequest", reflect.TypeOf((*MockWorkMailMessageFlowAPI)(nil).GetRawMessageContentRequest), arg0)
}

// GetRawMessageContentWithContext mocks base method
func (m *MockWorkMailMessageFlowAPI) GetRawMessageContentWithContext(arg0 context.Context, arg1 *workmailmessageflow.GetRawMessageContentInput, arg2 ...request.Option) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRawMessageContentWithContext", varargs...)
	ret0, _ := ret[0].(*workmailmessageflow.GetRawMessageContentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawMessageContentWithContext indicates an expected call of GetRawMessageContentWithContext
func (mr *MockWorkMailMessageFlowAPIMockRecorder) GetRawMessageContentWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawMessageContentWithContext", reflect.TypeOf((*MockWorkMailMessageFlowAPI)(nil).GetRawMessageContentWithContext), varargs...)
}
