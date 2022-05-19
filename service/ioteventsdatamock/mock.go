// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/ioteventsdata/ioteventsdataiface (interfaces: IoTEventsDataAPI)

// Package ioteventsdatamock is a generated GoMock package.
package ioteventsdatamock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	ioteventsdata "github.com/aws/aws-sdk-go/service/ioteventsdata"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIoTEventsDataAPI is a mock of IoTEventsDataAPI interface
type MockIoTEventsDataAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIoTEventsDataAPIMockRecorder
}

// MockIoTEventsDataAPIMockRecorder is the mock recorder for MockIoTEventsDataAPI
type MockIoTEventsDataAPIMockRecorder struct {
	mock *MockIoTEventsDataAPI
}

// NewMockIoTEventsDataAPI creates a new mock instance
func NewMockIoTEventsDataAPI(ctrl *gomock.Controller) *MockIoTEventsDataAPI {
	mock := &MockIoTEventsDataAPI{ctrl: ctrl}
	mock.recorder = &MockIoTEventsDataAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIoTEventsDataAPI) EXPECT() *MockIoTEventsDataAPIMockRecorder {
	return m.recorder
}

// BatchAcknowledgeAlarm mocks base method
func (m *MockIoTEventsDataAPI) BatchAcknowledgeAlarm(arg0 *ioteventsdata.BatchAcknowledgeAlarmInput) (*ioteventsdata.BatchAcknowledgeAlarmOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchAcknowledgeAlarm", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchAcknowledgeAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchAcknowledgeAlarm indicates an expected call of BatchAcknowledgeAlarm
func (mr *MockIoTEventsDataAPIMockRecorder) BatchAcknowledgeAlarm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchAcknowledgeAlarm", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchAcknowledgeAlarm), arg0)
}

// BatchAcknowledgeAlarmRequest mocks base method
func (m *MockIoTEventsDataAPI) BatchAcknowledgeAlarmRequest(arg0 *ioteventsdata.BatchAcknowledgeAlarmInput) (*request.Request, *ioteventsdata.BatchAcknowledgeAlarmOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchAcknowledgeAlarmRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchAcknowledgeAlarmOutput)
	return ret0, ret1
}

// BatchAcknowledgeAlarmRequest indicates an expected call of BatchAcknowledgeAlarmRequest
func (mr *MockIoTEventsDataAPIMockRecorder) BatchAcknowledgeAlarmRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchAcknowledgeAlarmRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchAcknowledgeAlarmRequest), arg0)
}

// BatchAcknowledgeAlarmWithContext mocks base method
func (m *MockIoTEventsDataAPI) BatchAcknowledgeAlarmWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchAcknowledgeAlarmInput, arg2 ...request.Option) (*ioteventsdata.BatchAcknowledgeAlarmOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchAcknowledgeAlarmWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchAcknowledgeAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchAcknowledgeAlarmWithContext indicates an expected call of BatchAcknowledgeAlarmWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) BatchAcknowledgeAlarmWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchAcknowledgeAlarmWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchAcknowledgeAlarmWithContext), varargs...)
}

// BatchDeleteDetector mocks base method
func (m *MockIoTEventsDataAPI) BatchDeleteDetector(arg0 *ioteventsdata.BatchDeleteDetectorInput) (*ioteventsdata.BatchDeleteDetectorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDeleteDetector", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchDeleteDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchDeleteDetector indicates an expected call of BatchDeleteDetector
func (mr *MockIoTEventsDataAPIMockRecorder) BatchDeleteDetector(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteDetector", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchDeleteDetector), arg0)
}

// BatchDeleteDetectorRequest mocks base method
func (m *MockIoTEventsDataAPI) BatchDeleteDetectorRequest(arg0 *ioteventsdata.BatchDeleteDetectorInput) (*request.Request, *ioteventsdata.BatchDeleteDetectorOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDeleteDetectorRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchDeleteDetectorOutput)
	return ret0, ret1
}

// BatchDeleteDetectorRequest indicates an expected call of BatchDeleteDetectorRequest
func (mr *MockIoTEventsDataAPIMockRecorder) BatchDeleteDetectorRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteDetectorRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchDeleteDetectorRequest), arg0)
}

// BatchDeleteDetectorWithContext mocks base method
func (m *MockIoTEventsDataAPI) BatchDeleteDetectorWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchDeleteDetectorInput, arg2 ...request.Option) (*ioteventsdata.BatchDeleteDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchDeleteDetectorWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchDeleteDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchDeleteDetectorWithContext indicates an expected call of BatchDeleteDetectorWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) BatchDeleteDetectorWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteDetectorWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchDeleteDetectorWithContext), varargs...)
}

// BatchDisableAlarm mocks base method
func (m *MockIoTEventsDataAPI) BatchDisableAlarm(arg0 *ioteventsdata.BatchDisableAlarmInput) (*ioteventsdata.BatchDisableAlarmOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDisableAlarm", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchDisableAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchDisableAlarm indicates an expected call of BatchDisableAlarm
func (mr *MockIoTEventsDataAPIMockRecorder) BatchDisableAlarm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDisableAlarm", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchDisableAlarm), arg0)
}

// BatchDisableAlarmRequest mocks base method
func (m *MockIoTEventsDataAPI) BatchDisableAlarmRequest(arg0 *ioteventsdata.BatchDisableAlarmInput) (*request.Request, *ioteventsdata.BatchDisableAlarmOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDisableAlarmRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchDisableAlarmOutput)
	return ret0, ret1
}

// BatchDisableAlarmRequest indicates an expected call of BatchDisableAlarmRequest
func (mr *MockIoTEventsDataAPIMockRecorder) BatchDisableAlarmRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDisableAlarmRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchDisableAlarmRequest), arg0)
}

// BatchDisableAlarmWithContext mocks base method
func (m *MockIoTEventsDataAPI) BatchDisableAlarmWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchDisableAlarmInput, arg2 ...request.Option) (*ioteventsdata.BatchDisableAlarmOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchDisableAlarmWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchDisableAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchDisableAlarmWithContext indicates an expected call of BatchDisableAlarmWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) BatchDisableAlarmWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDisableAlarmWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchDisableAlarmWithContext), varargs...)
}

// BatchEnableAlarm mocks base method
func (m *MockIoTEventsDataAPI) BatchEnableAlarm(arg0 *ioteventsdata.BatchEnableAlarmInput) (*ioteventsdata.BatchEnableAlarmOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchEnableAlarm", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchEnableAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchEnableAlarm indicates an expected call of BatchEnableAlarm
func (mr *MockIoTEventsDataAPIMockRecorder) BatchEnableAlarm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchEnableAlarm", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchEnableAlarm), arg0)
}

// BatchEnableAlarmRequest mocks base method
func (m *MockIoTEventsDataAPI) BatchEnableAlarmRequest(arg0 *ioteventsdata.BatchEnableAlarmInput) (*request.Request, *ioteventsdata.BatchEnableAlarmOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchEnableAlarmRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchEnableAlarmOutput)
	return ret0, ret1
}

// BatchEnableAlarmRequest indicates an expected call of BatchEnableAlarmRequest
func (mr *MockIoTEventsDataAPIMockRecorder) BatchEnableAlarmRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchEnableAlarmRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchEnableAlarmRequest), arg0)
}

// BatchEnableAlarmWithContext mocks base method
func (m *MockIoTEventsDataAPI) BatchEnableAlarmWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchEnableAlarmInput, arg2 ...request.Option) (*ioteventsdata.BatchEnableAlarmOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchEnableAlarmWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchEnableAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchEnableAlarmWithContext indicates an expected call of BatchEnableAlarmWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) BatchEnableAlarmWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchEnableAlarmWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchEnableAlarmWithContext), varargs...)
}

// BatchPutMessage mocks base method
func (m *MockIoTEventsDataAPI) BatchPutMessage(arg0 *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchPutMessage", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchPutMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchPutMessage indicates an expected call of BatchPutMessage
func (mr *MockIoTEventsDataAPIMockRecorder) BatchPutMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchPutMessage", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchPutMessage), arg0)
}

// BatchPutMessageRequest mocks base method
func (m *MockIoTEventsDataAPI) BatchPutMessageRequest(arg0 *ioteventsdata.BatchPutMessageInput) (*request.Request, *ioteventsdata.BatchPutMessageOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchPutMessageRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchPutMessageOutput)
	return ret0, ret1
}

// BatchPutMessageRequest indicates an expected call of BatchPutMessageRequest
func (mr *MockIoTEventsDataAPIMockRecorder) BatchPutMessageRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchPutMessageRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchPutMessageRequest), arg0)
}

// BatchPutMessageWithContext mocks base method
func (m *MockIoTEventsDataAPI) BatchPutMessageWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchPutMessageInput, arg2 ...request.Option) (*ioteventsdata.BatchPutMessageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchPutMessageWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchPutMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchPutMessageWithContext indicates an expected call of BatchPutMessageWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) BatchPutMessageWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchPutMessageWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchPutMessageWithContext), varargs...)
}

// BatchResetAlarm mocks base method
func (m *MockIoTEventsDataAPI) BatchResetAlarm(arg0 *ioteventsdata.BatchResetAlarmInput) (*ioteventsdata.BatchResetAlarmOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchResetAlarm", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchResetAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchResetAlarm indicates an expected call of BatchResetAlarm
func (mr *MockIoTEventsDataAPIMockRecorder) BatchResetAlarm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchResetAlarm", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchResetAlarm), arg0)
}

// BatchResetAlarmRequest mocks base method
func (m *MockIoTEventsDataAPI) BatchResetAlarmRequest(arg0 *ioteventsdata.BatchResetAlarmInput) (*request.Request, *ioteventsdata.BatchResetAlarmOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchResetAlarmRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchResetAlarmOutput)
	return ret0, ret1
}

// BatchResetAlarmRequest indicates an expected call of BatchResetAlarmRequest
func (mr *MockIoTEventsDataAPIMockRecorder) BatchResetAlarmRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchResetAlarmRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchResetAlarmRequest), arg0)
}

// BatchResetAlarmWithContext mocks base method
func (m *MockIoTEventsDataAPI) BatchResetAlarmWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchResetAlarmInput, arg2 ...request.Option) (*ioteventsdata.BatchResetAlarmOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchResetAlarmWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchResetAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchResetAlarmWithContext indicates an expected call of BatchResetAlarmWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) BatchResetAlarmWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchResetAlarmWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchResetAlarmWithContext), varargs...)
}

// BatchSnoozeAlarm mocks base method
func (m *MockIoTEventsDataAPI) BatchSnoozeAlarm(arg0 *ioteventsdata.BatchSnoozeAlarmInput) (*ioteventsdata.BatchSnoozeAlarmOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchSnoozeAlarm", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchSnoozeAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchSnoozeAlarm indicates an expected call of BatchSnoozeAlarm
func (mr *MockIoTEventsDataAPIMockRecorder) BatchSnoozeAlarm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchSnoozeAlarm", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchSnoozeAlarm), arg0)
}

// BatchSnoozeAlarmRequest mocks base method
func (m *MockIoTEventsDataAPI) BatchSnoozeAlarmRequest(arg0 *ioteventsdata.BatchSnoozeAlarmInput) (*request.Request, *ioteventsdata.BatchSnoozeAlarmOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchSnoozeAlarmRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchSnoozeAlarmOutput)
	return ret0, ret1
}

// BatchSnoozeAlarmRequest indicates an expected call of BatchSnoozeAlarmRequest
func (mr *MockIoTEventsDataAPIMockRecorder) BatchSnoozeAlarmRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchSnoozeAlarmRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchSnoozeAlarmRequest), arg0)
}

// BatchSnoozeAlarmWithContext mocks base method
func (m *MockIoTEventsDataAPI) BatchSnoozeAlarmWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchSnoozeAlarmInput, arg2 ...request.Option) (*ioteventsdata.BatchSnoozeAlarmOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchSnoozeAlarmWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchSnoozeAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchSnoozeAlarmWithContext indicates an expected call of BatchSnoozeAlarmWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) BatchSnoozeAlarmWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchSnoozeAlarmWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchSnoozeAlarmWithContext), varargs...)
}

// BatchUpdateDetector mocks base method
func (m *MockIoTEventsDataAPI) BatchUpdateDetector(arg0 *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateDetector", arg0)
	ret0, _ := ret[0].(*ioteventsdata.BatchUpdateDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdateDetector indicates an expected call of BatchUpdateDetector
func (mr *MockIoTEventsDataAPIMockRecorder) BatchUpdateDetector(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateDetector", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchUpdateDetector), arg0)
}

// BatchUpdateDetectorRequest mocks base method
func (m *MockIoTEventsDataAPI) BatchUpdateDetectorRequest(arg0 *ioteventsdata.BatchUpdateDetectorInput) (*request.Request, *ioteventsdata.BatchUpdateDetectorOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateDetectorRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.BatchUpdateDetectorOutput)
	return ret0, ret1
}

// BatchUpdateDetectorRequest indicates an expected call of BatchUpdateDetectorRequest
func (mr *MockIoTEventsDataAPIMockRecorder) BatchUpdateDetectorRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateDetectorRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchUpdateDetectorRequest), arg0)
}

// BatchUpdateDetectorWithContext mocks base method
func (m *MockIoTEventsDataAPI) BatchUpdateDetectorWithContext(arg0 context.Context, arg1 *ioteventsdata.BatchUpdateDetectorInput, arg2 ...request.Option) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchUpdateDetectorWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.BatchUpdateDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdateDetectorWithContext indicates an expected call of BatchUpdateDetectorWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) BatchUpdateDetectorWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateDetectorWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).BatchUpdateDetectorWithContext), varargs...)
}

// DescribeAlarm mocks base method
func (m *MockIoTEventsDataAPI) DescribeAlarm(arg0 *ioteventsdata.DescribeAlarmInput) (*ioteventsdata.DescribeAlarmOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAlarm", arg0)
	ret0, _ := ret[0].(*ioteventsdata.DescribeAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAlarm indicates an expected call of DescribeAlarm
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeAlarm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAlarm", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeAlarm), arg0)
}

// DescribeAlarmRequest mocks base method
func (m *MockIoTEventsDataAPI) DescribeAlarmRequest(arg0 *ioteventsdata.DescribeAlarmInput) (*request.Request, *ioteventsdata.DescribeAlarmOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAlarmRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.DescribeAlarmOutput)
	return ret0, ret1
}

// DescribeAlarmRequest indicates an expected call of DescribeAlarmRequest
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeAlarmRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAlarmRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeAlarmRequest), arg0)
}

// DescribeAlarmWithContext mocks base method
func (m *MockIoTEventsDataAPI) DescribeAlarmWithContext(arg0 context.Context, arg1 *ioteventsdata.DescribeAlarmInput, arg2 ...request.Option) (*ioteventsdata.DescribeAlarmOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAlarmWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.DescribeAlarmOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAlarmWithContext indicates an expected call of DescribeAlarmWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeAlarmWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAlarmWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeAlarmWithContext), varargs...)
}

// DescribeDetector mocks base method
func (m *MockIoTEventsDataAPI) DescribeDetector(arg0 *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDetector", arg0)
	ret0, _ := ret[0].(*ioteventsdata.DescribeDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDetector indicates an expected call of DescribeDetector
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeDetector(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDetector", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeDetector), arg0)
}

// DescribeDetectorRequest mocks base method
func (m *MockIoTEventsDataAPI) DescribeDetectorRequest(arg0 *ioteventsdata.DescribeDetectorInput) (*request.Request, *ioteventsdata.DescribeDetectorOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDetectorRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.DescribeDetectorOutput)
	return ret0, ret1
}

// DescribeDetectorRequest indicates an expected call of DescribeDetectorRequest
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeDetectorRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDetectorRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeDetectorRequest), arg0)
}

// DescribeDetectorWithContext mocks base method
func (m *MockIoTEventsDataAPI) DescribeDetectorWithContext(arg0 context.Context, arg1 *ioteventsdata.DescribeDetectorInput, arg2 ...request.Option) (*ioteventsdata.DescribeDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDetectorWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.DescribeDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDetectorWithContext indicates an expected call of DescribeDetectorWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) DescribeDetectorWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDetectorWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).DescribeDetectorWithContext), varargs...)
}

// ListAlarms mocks base method
func (m *MockIoTEventsDataAPI) ListAlarms(arg0 *ioteventsdata.ListAlarmsInput) (*ioteventsdata.ListAlarmsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlarms", arg0)
	ret0, _ := ret[0].(*ioteventsdata.ListAlarmsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlarms indicates an expected call of ListAlarms
func (mr *MockIoTEventsDataAPIMockRecorder) ListAlarms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlarms", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListAlarms), arg0)
}

// ListAlarmsRequest mocks base method
func (m *MockIoTEventsDataAPI) ListAlarmsRequest(arg0 *ioteventsdata.ListAlarmsInput) (*request.Request, *ioteventsdata.ListAlarmsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlarmsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.ListAlarmsOutput)
	return ret0, ret1
}

// ListAlarmsRequest indicates an expected call of ListAlarmsRequest
func (mr *MockIoTEventsDataAPIMockRecorder) ListAlarmsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlarmsRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListAlarmsRequest), arg0)
}

// ListAlarmsWithContext mocks base method
func (m *MockIoTEventsDataAPI) ListAlarmsWithContext(arg0 context.Context, arg1 *ioteventsdata.ListAlarmsInput, arg2 ...request.Option) (*ioteventsdata.ListAlarmsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAlarmsWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.ListAlarmsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlarmsWithContext indicates an expected call of ListAlarmsWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) ListAlarmsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlarmsWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListAlarmsWithContext), varargs...)
}

// ListDetectors mocks base method
func (m *MockIoTEventsDataAPI) ListDetectors(arg0 *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDetectors", arg0)
	ret0, _ := ret[0].(*ioteventsdata.ListDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDetectors indicates an expected call of ListDetectors
func (mr *MockIoTEventsDataAPIMockRecorder) ListDetectors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetectors", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListDetectors), arg0)
}

// ListDetectorsRequest mocks base method
func (m *MockIoTEventsDataAPI) ListDetectorsRequest(arg0 *ioteventsdata.ListDetectorsInput) (*request.Request, *ioteventsdata.ListDetectorsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDetectorsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ioteventsdata.ListDetectorsOutput)
	return ret0, ret1
}

// ListDetectorsRequest indicates an expected call of ListDetectorsRequest
func (mr *MockIoTEventsDataAPIMockRecorder) ListDetectorsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetectorsRequest", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListDetectorsRequest), arg0)
}

// ListDetectorsWithContext mocks base method
func (m *MockIoTEventsDataAPI) ListDetectorsWithContext(arg0 context.Context, arg1 *ioteventsdata.ListDetectorsInput, arg2 ...request.Option) (*ioteventsdata.ListDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDetectorsWithContext", varargs...)
	ret0, _ := ret[0].(*ioteventsdata.ListDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDetectorsWithContext indicates an expected call of ListDetectorsWithContext
func (mr *MockIoTEventsDataAPIMockRecorder) ListDetectorsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetectorsWithContext", reflect.TypeOf((*MockIoTEventsDataAPI)(nil).ListDetectorsWithContext), varargs...)
}
