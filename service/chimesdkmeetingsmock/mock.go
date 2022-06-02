// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/chimesdkmeetings/chimesdkmeetingsiface (interfaces: ChimeSDKMeetingsAPI)

// Package chimesdkmeetingsmock is a generated GoMock package.
package chimesdkmeetingsmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	chimesdkmeetings "github.com/aws/aws-sdk-go/service/chimesdkmeetings"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChimeSDKMeetingsAPI is a mock of ChimeSDKMeetingsAPI interface
type MockChimeSDKMeetingsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockChimeSDKMeetingsAPIMockRecorder
}

// MockChimeSDKMeetingsAPIMockRecorder is the mock recorder for MockChimeSDKMeetingsAPI
type MockChimeSDKMeetingsAPIMockRecorder struct {
	mock *MockChimeSDKMeetingsAPI
}

// NewMockChimeSDKMeetingsAPI creates a new mock instance
func NewMockChimeSDKMeetingsAPI(ctrl *gomock.Controller) *MockChimeSDKMeetingsAPI {
	mock := &MockChimeSDKMeetingsAPI{ctrl: ctrl}
	mock.recorder = &MockChimeSDKMeetingsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChimeSDKMeetingsAPI) EXPECT() *MockChimeSDKMeetingsAPIMockRecorder {
	return m.recorder
}

// BatchCreateAttendee mocks base method
func (m *MockChimeSDKMeetingsAPI) BatchCreateAttendee(arg0 *chimesdkmeetings.BatchCreateAttendeeInput) (*chimesdkmeetings.BatchCreateAttendeeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateAttendee", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.BatchCreateAttendeeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCreateAttendee indicates an expected call of BatchCreateAttendee
func (mr *MockChimeSDKMeetingsAPIMockRecorder) BatchCreateAttendee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateAttendee", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).BatchCreateAttendee), arg0)
}

// BatchCreateAttendeeRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) BatchCreateAttendeeRequest(arg0 *chimesdkmeetings.BatchCreateAttendeeInput) (*request.Request, *chimesdkmeetings.BatchCreateAttendeeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateAttendeeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.BatchCreateAttendeeOutput)
	return ret0, ret1
}

// BatchCreateAttendeeRequest indicates an expected call of BatchCreateAttendeeRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) BatchCreateAttendeeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateAttendeeRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).BatchCreateAttendeeRequest), arg0)
}

// BatchCreateAttendeeWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) BatchCreateAttendeeWithContext(arg0 context.Context, arg1 *chimesdkmeetings.BatchCreateAttendeeInput, arg2 ...request.Option) (*chimesdkmeetings.BatchCreateAttendeeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchCreateAttendeeWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.BatchCreateAttendeeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCreateAttendeeWithContext indicates an expected call of BatchCreateAttendeeWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) BatchCreateAttendeeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateAttendeeWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).BatchCreateAttendeeWithContext), varargs...)
}

// BatchUpdateAttendeeCapabilitiesExcept mocks base method
func (m *MockChimeSDKMeetingsAPI) BatchUpdateAttendeeCapabilitiesExcept(arg0 *chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptInput) (*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateAttendeeCapabilitiesExcept", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdateAttendeeCapabilitiesExcept indicates an expected call of BatchUpdateAttendeeCapabilitiesExcept
func (mr *MockChimeSDKMeetingsAPIMockRecorder) BatchUpdateAttendeeCapabilitiesExcept(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateAttendeeCapabilitiesExcept", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).BatchUpdateAttendeeCapabilitiesExcept), arg0)
}

// BatchUpdateAttendeeCapabilitiesExceptRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) BatchUpdateAttendeeCapabilitiesExceptRequest(arg0 *chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptInput) (*request.Request, *chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateAttendeeCapabilitiesExceptRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput)
	return ret0, ret1
}

// BatchUpdateAttendeeCapabilitiesExceptRequest indicates an expected call of BatchUpdateAttendeeCapabilitiesExceptRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) BatchUpdateAttendeeCapabilitiesExceptRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateAttendeeCapabilitiesExceptRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).BatchUpdateAttendeeCapabilitiesExceptRequest), arg0)
}

// BatchUpdateAttendeeCapabilitiesExceptWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) BatchUpdateAttendeeCapabilitiesExceptWithContext(arg0 context.Context, arg1 *chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptInput, arg2 ...request.Option) (*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchUpdateAttendeeCapabilitiesExceptWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdateAttendeeCapabilitiesExceptWithContext indicates an expected call of BatchUpdateAttendeeCapabilitiesExceptWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) BatchUpdateAttendeeCapabilitiesExceptWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateAttendeeCapabilitiesExceptWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).BatchUpdateAttendeeCapabilitiesExceptWithContext), varargs...)
}

// CreateAttendee mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateAttendee(arg0 *chimesdkmeetings.CreateAttendeeInput) (*chimesdkmeetings.CreateAttendeeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttendee", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.CreateAttendeeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttendee indicates an expected call of CreateAttendee
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateAttendee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttendee", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateAttendee), arg0)
}

// CreateAttendeeRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateAttendeeRequest(arg0 *chimesdkmeetings.CreateAttendeeInput) (*request.Request, *chimesdkmeetings.CreateAttendeeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttendeeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.CreateAttendeeOutput)
	return ret0, ret1
}

// CreateAttendeeRequest indicates an expected call of CreateAttendeeRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateAttendeeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttendeeRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateAttendeeRequest), arg0)
}

// CreateAttendeeWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateAttendeeWithContext(arg0 context.Context, arg1 *chimesdkmeetings.CreateAttendeeInput, arg2 ...request.Option) (*chimesdkmeetings.CreateAttendeeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAttendeeWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.CreateAttendeeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttendeeWithContext indicates an expected call of CreateAttendeeWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateAttendeeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttendeeWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateAttendeeWithContext), varargs...)
}

// CreateMeeting mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateMeeting(arg0 *chimesdkmeetings.CreateMeetingInput) (*chimesdkmeetings.CreateMeetingOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeeting", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.CreateMeetingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMeeting indicates an expected call of CreateMeeting
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateMeeting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeeting", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateMeeting), arg0)
}

// CreateMeetingRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateMeetingRequest(arg0 *chimesdkmeetings.CreateMeetingInput) (*request.Request, *chimesdkmeetings.CreateMeetingOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeetingRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.CreateMeetingOutput)
	return ret0, ret1
}

// CreateMeetingRequest indicates an expected call of CreateMeetingRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateMeetingRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeetingRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateMeetingRequest), arg0)
}

// CreateMeetingWithAttendees mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateMeetingWithAttendees(arg0 *chimesdkmeetings.CreateMeetingWithAttendeesInput) (*chimesdkmeetings.CreateMeetingWithAttendeesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeetingWithAttendees", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.CreateMeetingWithAttendeesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMeetingWithAttendees indicates an expected call of CreateMeetingWithAttendees
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateMeetingWithAttendees(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeetingWithAttendees", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateMeetingWithAttendees), arg0)
}

// CreateMeetingWithAttendeesRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateMeetingWithAttendeesRequest(arg0 *chimesdkmeetings.CreateMeetingWithAttendeesInput) (*request.Request, *chimesdkmeetings.CreateMeetingWithAttendeesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeetingWithAttendeesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.CreateMeetingWithAttendeesOutput)
	return ret0, ret1
}

// CreateMeetingWithAttendeesRequest indicates an expected call of CreateMeetingWithAttendeesRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateMeetingWithAttendeesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeetingWithAttendeesRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateMeetingWithAttendeesRequest), arg0)
}

// CreateMeetingWithAttendeesWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateMeetingWithAttendeesWithContext(arg0 context.Context, arg1 *chimesdkmeetings.CreateMeetingWithAttendeesInput, arg2 ...request.Option) (*chimesdkmeetings.CreateMeetingWithAttendeesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMeetingWithAttendeesWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.CreateMeetingWithAttendeesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMeetingWithAttendeesWithContext indicates an expected call of CreateMeetingWithAttendeesWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateMeetingWithAttendeesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeetingWithAttendeesWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateMeetingWithAttendeesWithContext), varargs...)
}

// CreateMeetingWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) CreateMeetingWithContext(arg0 context.Context, arg1 *chimesdkmeetings.CreateMeetingInput, arg2 ...request.Option) (*chimesdkmeetings.CreateMeetingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMeetingWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.CreateMeetingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMeetingWithContext indicates an expected call of CreateMeetingWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) CreateMeetingWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeetingWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).CreateMeetingWithContext), varargs...)
}

// DeleteAttendee mocks base method
func (m *MockChimeSDKMeetingsAPI) DeleteAttendee(arg0 *chimesdkmeetings.DeleteAttendeeInput) (*chimesdkmeetings.DeleteAttendeeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAttendee", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.DeleteAttendeeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAttendee indicates an expected call of DeleteAttendee
func (mr *MockChimeSDKMeetingsAPIMockRecorder) DeleteAttendee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttendee", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).DeleteAttendee), arg0)
}

// DeleteAttendeeRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) DeleteAttendeeRequest(arg0 *chimesdkmeetings.DeleteAttendeeInput) (*request.Request, *chimesdkmeetings.DeleteAttendeeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAttendeeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.DeleteAttendeeOutput)
	return ret0, ret1
}

// DeleteAttendeeRequest indicates an expected call of DeleteAttendeeRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) DeleteAttendeeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttendeeRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).DeleteAttendeeRequest), arg0)
}

// DeleteAttendeeWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) DeleteAttendeeWithContext(arg0 context.Context, arg1 *chimesdkmeetings.DeleteAttendeeInput, arg2 ...request.Option) (*chimesdkmeetings.DeleteAttendeeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAttendeeWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.DeleteAttendeeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAttendeeWithContext indicates an expected call of DeleteAttendeeWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) DeleteAttendeeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttendeeWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).DeleteAttendeeWithContext), varargs...)
}

// DeleteMeeting mocks base method
func (m *MockChimeSDKMeetingsAPI) DeleteMeeting(arg0 *chimesdkmeetings.DeleteMeetingInput) (*chimesdkmeetings.DeleteMeetingOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMeeting", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.DeleteMeetingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMeeting indicates an expected call of DeleteMeeting
func (mr *MockChimeSDKMeetingsAPIMockRecorder) DeleteMeeting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeeting", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).DeleteMeeting), arg0)
}

// DeleteMeetingRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) DeleteMeetingRequest(arg0 *chimesdkmeetings.DeleteMeetingInput) (*request.Request, *chimesdkmeetings.DeleteMeetingOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMeetingRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.DeleteMeetingOutput)
	return ret0, ret1
}

// DeleteMeetingRequest indicates an expected call of DeleteMeetingRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) DeleteMeetingRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeetingRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).DeleteMeetingRequest), arg0)
}

// DeleteMeetingWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) DeleteMeetingWithContext(arg0 context.Context, arg1 *chimesdkmeetings.DeleteMeetingInput, arg2 ...request.Option) (*chimesdkmeetings.DeleteMeetingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMeetingWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.DeleteMeetingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMeetingWithContext indicates an expected call of DeleteMeetingWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) DeleteMeetingWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeetingWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).DeleteMeetingWithContext), varargs...)
}

// GetAttendee mocks base method
func (m *MockChimeSDKMeetingsAPI) GetAttendee(arg0 *chimesdkmeetings.GetAttendeeInput) (*chimesdkmeetings.GetAttendeeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendee", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.GetAttendeeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttendee indicates an expected call of GetAttendee
func (mr *MockChimeSDKMeetingsAPIMockRecorder) GetAttendee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendee", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).GetAttendee), arg0)
}

// GetAttendeeRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) GetAttendeeRequest(arg0 *chimesdkmeetings.GetAttendeeInput) (*request.Request, *chimesdkmeetings.GetAttendeeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendeeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.GetAttendeeOutput)
	return ret0, ret1
}

// GetAttendeeRequest indicates an expected call of GetAttendeeRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) GetAttendeeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendeeRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).GetAttendeeRequest), arg0)
}

// GetAttendeeWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) GetAttendeeWithContext(arg0 context.Context, arg1 *chimesdkmeetings.GetAttendeeInput, arg2 ...request.Option) (*chimesdkmeetings.GetAttendeeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAttendeeWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.GetAttendeeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttendeeWithContext indicates an expected call of GetAttendeeWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) GetAttendeeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendeeWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).GetAttendeeWithContext), varargs...)
}

// GetMeeting mocks base method
func (m *MockChimeSDKMeetingsAPI) GetMeeting(arg0 *chimesdkmeetings.GetMeetingInput) (*chimesdkmeetings.GetMeetingOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeeting", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.GetMeetingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeeting indicates an expected call of GetMeeting
func (mr *MockChimeSDKMeetingsAPIMockRecorder) GetMeeting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeeting", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).GetMeeting), arg0)
}

// GetMeetingRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) GetMeetingRequest(arg0 *chimesdkmeetings.GetMeetingInput) (*request.Request, *chimesdkmeetings.GetMeetingOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeetingRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.GetMeetingOutput)
	return ret0, ret1
}

// GetMeetingRequest indicates an expected call of GetMeetingRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) GetMeetingRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).GetMeetingRequest), arg0)
}

// GetMeetingWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) GetMeetingWithContext(arg0 context.Context, arg1 *chimesdkmeetings.GetMeetingInput, arg2 ...request.Option) (*chimesdkmeetings.GetMeetingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMeetingWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.GetMeetingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingWithContext indicates an expected call of GetMeetingWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) GetMeetingWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).GetMeetingWithContext), varargs...)
}

// ListAttendees mocks base method
func (m *MockChimeSDKMeetingsAPI) ListAttendees(arg0 *chimesdkmeetings.ListAttendeesInput) (*chimesdkmeetings.ListAttendeesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttendees", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.ListAttendeesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttendees indicates an expected call of ListAttendees
func (mr *MockChimeSDKMeetingsAPIMockRecorder) ListAttendees(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttendees", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).ListAttendees), arg0)
}

// ListAttendeesPages mocks base method
func (m *MockChimeSDKMeetingsAPI) ListAttendeesPages(arg0 *chimesdkmeetings.ListAttendeesInput, arg1 func(*chimesdkmeetings.ListAttendeesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttendeesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAttendeesPages indicates an expected call of ListAttendeesPages
func (mr *MockChimeSDKMeetingsAPIMockRecorder) ListAttendeesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttendeesPages", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).ListAttendeesPages), arg0, arg1)
}

// ListAttendeesPagesWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) ListAttendeesPagesWithContext(arg0 context.Context, arg1 *chimesdkmeetings.ListAttendeesInput, arg2 func(*chimesdkmeetings.ListAttendeesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttendeesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAttendeesPagesWithContext indicates an expected call of ListAttendeesPagesWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) ListAttendeesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttendeesPagesWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).ListAttendeesPagesWithContext), varargs...)
}

// ListAttendeesRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) ListAttendeesRequest(arg0 *chimesdkmeetings.ListAttendeesInput) (*request.Request, *chimesdkmeetings.ListAttendeesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttendeesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.ListAttendeesOutput)
	return ret0, ret1
}

// ListAttendeesRequest indicates an expected call of ListAttendeesRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) ListAttendeesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttendeesRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).ListAttendeesRequest), arg0)
}

// ListAttendeesWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) ListAttendeesWithContext(arg0 context.Context, arg1 *chimesdkmeetings.ListAttendeesInput, arg2 ...request.Option) (*chimesdkmeetings.ListAttendeesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttendeesWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.ListAttendeesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttendeesWithContext indicates an expected call of ListAttendeesWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) ListAttendeesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttendeesWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).ListAttendeesWithContext), varargs...)
}

// StartMeetingTranscription mocks base method
func (m *MockChimeSDKMeetingsAPI) StartMeetingTranscription(arg0 *chimesdkmeetings.StartMeetingTranscriptionInput) (*chimesdkmeetings.StartMeetingTranscriptionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartMeetingTranscription", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.StartMeetingTranscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartMeetingTranscription indicates an expected call of StartMeetingTranscription
func (mr *MockChimeSDKMeetingsAPIMockRecorder) StartMeetingTranscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMeetingTranscription", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).StartMeetingTranscription), arg0)
}

// StartMeetingTranscriptionRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) StartMeetingTranscriptionRequest(arg0 *chimesdkmeetings.StartMeetingTranscriptionInput) (*request.Request, *chimesdkmeetings.StartMeetingTranscriptionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartMeetingTranscriptionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.StartMeetingTranscriptionOutput)
	return ret0, ret1
}

// StartMeetingTranscriptionRequest indicates an expected call of StartMeetingTranscriptionRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) StartMeetingTranscriptionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMeetingTranscriptionRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).StartMeetingTranscriptionRequest), arg0)
}

// StartMeetingTranscriptionWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) StartMeetingTranscriptionWithContext(arg0 context.Context, arg1 *chimesdkmeetings.StartMeetingTranscriptionInput, arg2 ...request.Option) (*chimesdkmeetings.StartMeetingTranscriptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartMeetingTranscriptionWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.StartMeetingTranscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartMeetingTranscriptionWithContext indicates an expected call of StartMeetingTranscriptionWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) StartMeetingTranscriptionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMeetingTranscriptionWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).StartMeetingTranscriptionWithContext), varargs...)
}

// StopMeetingTranscription mocks base method
func (m *MockChimeSDKMeetingsAPI) StopMeetingTranscription(arg0 *chimesdkmeetings.StopMeetingTranscriptionInput) (*chimesdkmeetings.StopMeetingTranscriptionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopMeetingTranscription", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.StopMeetingTranscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopMeetingTranscription indicates an expected call of StopMeetingTranscription
func (mr *MockChimeSDKMeetingsAPIMockRecorder) StopMeetingTranscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopMeetingTranscription", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).StopMeetingTranscription), arg0)
}

// StopMeetingTranscriptionRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) StopMeetingTranscriptionRequest(arg0 *chimesdkmeetings.StopMeetingTranscriptionInput) (*request.Request, *chimesdkmeetings.StopMeetingTranscriptionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopMeetingTranscriptionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.StopMeetingTranscriptionOutput)
	return ret0, ret1
}

// StopMeetingTranscriptionRequest indicates an expected call of StopMeetingTranscriptionRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) StopMeetingTranscriptionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopMeetingTranscriptionRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).StopMeetingTranscriptionRequest), arg0)
}

// StopMeetingTranscriptionWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) StopMeetingTranscriptionWithContext(arg0 context.Context, arg1 *chimesdkmeetings.StopMeetingTranscriptionInput, arg2 ...request.Option) (*chimesdkmeetings.StopMeetingTranscriptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopMeetingTranscriptionWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.StopMeetingTranscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopMeetingTranscriptionWithContext indicates an expected call of StopMeetingTranscriptionWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) StopMeetingTranscriptionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopMeetingTranscriptionWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).StopMeetingTranscriptionWithContext), varargs...)
}

// UpdateAttendeeCapabilities mocks base method
func (m *MockChimeSDKMeetingsAPI) UpdateAttendeeCapabilities(arg0 *chimesdkmeetings.UpdateAttendeeCapabilitiesInput) (*chimesdkmeetings.UpdateAttendeeCapabilitiesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAttendeeCapabilities", arg0)
	ret0, _ := ret[0].(*chimesdkmeetings.UpdateAttendeeCapabilitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAttendeeCapabilities indicates an expected call of UpdateAttendeeCapabilities
func (mr *MockChimeSDKMeetingsAPIMockRecorder) UpdateAttendeeCapabilities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttendeeCapabilities", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).UpdateAttendeeCapabilities), arg0)
}

// UpdateAttendeeCapabilitiesRequest mocks base method
func (m *MockChimeSDKMeetingsAPI) UpdateAttendeeCapabilitiesRequest(arg0 *chimesdkmeetings.UpdateAttendeeCapabilitiesInput) (*request.Request, *chimesdkmeetings.UpdateAttendeeCapabilitiesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAttendeeCapabilitiesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmeetings.UpdateAttendeeCapabilitiesOutput)
	return ret0, ret1
}

// UpdateAttendeeCapabilitiesRequest indicates an expected call of UpdateAttendeeCapabilitiesRequest
func (mr *MockChimeSDKMeetingsAPIMockRecorder) UpdateAttendeeCapabilitiesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttendeeCapabilitiesRequest", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).UpdateAttendeeCapabilitiesRequest), arg0)
}

// UpdateAttendeeCapabilitiesWithContext mocks base method
func (m *MockChimeSDKMeetingsAPI) UpdateAttendeeCapabilitiesWithContext(arg0 context.Context, arg1 *chimesdkmeetings.UpdateAttendeeCapabilitiesInput, arg2 ...request.Option) (*chimesdkmeetings.UpdateAttendeeCapabilitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAttendeeCapabilitiesWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmeetings.UpdateAttendeeCapabilitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAttendeeCapabilitiesWithContext indicates an expected call of UpdateAttendeeCapabilitiesWithContext
func (mr *MockChimeSDKMeetingsAPIMockRecorder) UpdateAttendeeCapabilitiesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttendeeCapabilitiesWithContext", reflect.TypeOf((*MockChimeSDKMeetingsAPI)(nil).UpdateAttendeeCapabilitiesWithContext), varargs...)
}
