// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/scheduler/scheduleriface (interfaces: SchedulerAPI)

// Package schedulermock is a generated GoMock package.
package schedulermock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	scheduler "github.com/aws/aws-sdk-go/service/scheduler"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSchedulerAPI is a mock of SchedulerAPI interface
type MockSchedulerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerAPIMockRecorder
}

// MockSchedulerAPIMockRecorder is the mock recorder for MockSchedulerAPI
type MockSchedulerAPIMockRecorder struct {
	mock *MockSchedulerAPI
}

// NewMockSchedulerAPI creates a new mock instance
func NewMockSchedulerAPI(ctrl *gomock.Controller) *MockSchedulerAPI {
	mock := &MockSchedulerAPI{ctrl: ctrl}
	mock.recorder = &MockSchedulerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSchedulerAPI) EXPECT() *MockSchedulerAPIMockRecorder {
	return m.recorder
}

// CreateSchedule mocks base method
func (m *MockSchedulerAPI) CreateSchedule(arg0 *scheduler.CreateScheduleInput) (*scheduler.CreateScheduleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchedule", arg0)
	ret0, _ := ret[0].(*scheduler.CreateScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchedule indicates an expected call of CreateSchedule
func (mr *MockSchedulerAPIMockRecorder) CreateSchedule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchedule", reflect.TypeOf((*MockSchedulerAPI)(nil).CreateSchedule), arg0)
}

// CreateScheduleGroup mocks base method
func (m *MockSchedulerAPI) CreateScheduleGroup(arg0 *scheduler.CreateScheduleGroupInput) (*scheduler.CreateScheduleGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScheduleGroup", arg0)
	ret0, _ := ret[0].(*scheduler.CreateScheduleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScheduleGroup indicates an expected call of CreateScheduleGroup
func (mr *MockSchedulerAPIMockRecorder) CreateScheduleGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScheduleGroup", reflect.TypeOf((*MockSchedulerAPI)(nil).CreateScheduleGroup), arg0)
}

// CreateScheduleGroupRequest mocks base method
func (m *MockSchedulerAPI) CreateScheduleGroupRequest(arg0 *scheduler.CreateScheduleGroupInput) (*request.Request, *scheduler.CreateScheduleGroupOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScheduleGroupRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.CreateScheduleGroupOutput)
	return ret0, ret1
}

// CreateScheduleGroupRequest indicates an expected call of CreateScheduleGroupRequest
func (mr *MockSchedulerAPIMockRecorder) CreateScheduleGroupRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScheduleGroupRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).CreateScheduleGroupRequest), arg0)
}

// CreateScheduleGroupWithContext mocks base method
func (m *MockSchedulerAPI) CreateScheduleGroupWithContext(arg0 context.Context, arg1 *scheduler.CreateScheduleGroupInput, arg2 ...request.Option) (*scheduler.CreateScheduleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateScheduleGroupWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.CreateScheduleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScheduleGroupWithContext indicates an expected call of CreateScheduleGroupWithContext
func (mr *MockSchedulerAPIMockRecorder) CreateScheduleGroupWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScheduleGroupWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).CreateScheduleGroupWithContext), varargs...)
}

// CreateScheduleRequest mocks base method
func (m *MockSchedulerAPI) CreateScheduleRequest(arg0 *scheduler.CreateScheduleInput) (*request.Request, *scheduler.CreateScheduleOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScheduleRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.CreateScheduleOutput)
	return ret0, ret1
}

// CreateScheduleRequest indicates an expected call of CreateScheduleRequest
func (mr *MockSchedulerAPIMockRecorder) CreateScheduleRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScheduleRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).CreateScheduleRequest), arg0)
}

// CreateScheduleWithContext mocks base method
func (m *MockSchedulerAPI) CreateScheduleWithContext(arg0 context.Context, arg1 *scheduler.CreateScheduleInput, arg2 ...request.Option) (*scheduler.CreateScheduleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateScheduleWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.CreateScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScheduleWithContext indicates an expected call of CreateScheduleWithContext
func (mr *MockSchedulerAPIMockRecorder) CreateScheduleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScheduleWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).CreateScheduleWithContext), varargs...)
}

// DeleteSchedule mocks base method
func (m *MockSchedulerAPI) DeleteSchedule(arg0 *scheduler.DeleteScheduleInput) (*scheduler.DeleteScheduleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSchedule", arg0)
	ret0, _ := ret[0].(*scheduler.DeleteScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSchedule indicates an expected call of DeleteSchedule
func (mr *MockSchedulerAPIMockRecorder) DeleteSchedule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSchedule", reflect.TypeOf((*MockSchedulerAPI)(nil).DeleteSchedule), arg0)
}

// DeleteScheduleGroup mocks base method
func (m *MockSchedulerAPI) DeleteScheduleGroup(arg0 *scheduler.DeleteScheduleGroupInput) (*scheduler.DeleteScheduleGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScheduleGroup", arg0)
	ret0, _ := ret[0].(*scheduler.DeleteScheduleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScheduleGroup indicates an expected call of DeleteScheduleGroup
func (mr *MockSchedulerAPIMockRecorder) DeleteScheduleGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduleGroup", reflect.TypeOf((*MockSchedulerAPI)(nil).DeleteScheduleGroup), arg0)
}

// DeleteScheduleGroupRequest mocks base method
func (m *MockSchedulerAPI) DeleteScheduleGroupRequest(arg0 *scheduler.DeleteScheduleGroupInput) (*request.Request, *scheduler.DeleteScheduleGroupOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScheduleGroupRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.DeleteScheduleGroupOutput)
	return ret0, ret1
}

// DeleteScheduleGroupRequest indicates an expected call of DeleteScheduleGroupRequest
func (mr *MockSchedulerAPIMockRecorder) DeleteScheduleGroupRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduleGroupRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).DeleteScheduleGroupRequest), arg0)
}

// DeleteScheduleGroupWithContext mocks base method
func (m *MockSchedulerAPI) DeleteScheduleGroupWithContext(arg0 context.Context, arg1 *scheduler.DeleteScheduleGroupInput, arg2 ...request.Option) (*scheduler.DeleteScheduleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteScheduleGroupWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.DeleteScheduleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScheduleGroupWithContext indicates an expected call of DeleteScheduleGroupWithContext
func (mr *MockSchedulerAPIMockRecorder) DeleteScheduleGroupWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduleGroupWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).DeleteScheduleGroupWithContext), varargs...)
}

// DeleteScheduleRequest mocks base method
func (m *MockSchedulerAPI) DeleteScheduleRequest(arg0 *scheduler.DeleteScheduleInput) (*request.Request, *scheduler.DeleteScheduleOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScheduleRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.DeleteScheduleOutput)
	return ret0, ret1
}

// DeleteScheduleRequest indicates an expected call of DeleteScheduleRequest
func (mr *MockSchedulerAPIMockRecorder) DeleteScheduleRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduleRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).DeleteScheduleRequest), arg0)
}

// DeleteScheduleWithContext mocks base method
func (m *MockSchedulerAPI) DeleteScheduleWithContext(arg0 context.Context, arg1 *scheduler.DeleteScheduleInput, arg2 ...request.Option) (*scheduler.DeleteScheduleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteScheduleWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.DeleteScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScheduleWithContext indicates an expected call of DeleteScheduleWithContext
func (mr *MockSchedulerAPIMockRecorder) DeleteScheduleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduleWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).DeleteScheduleWithContext), varargs...)
}

// GetSchedule mocks base method
func (m *MockSchedulerAPI) GetSchedule(arg0 *scheduler.GetScheduleInput) (*scheduler.GetScheduleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", arg0)
	ret0, _ := ret[0].(*scheduler.GetScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule
func (mr *MockSchedulerAPIMockRecorder) GetSchedule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockSchedulerAPI)(nil).GetSchedule), arg0)
}

// GetScheduleGroup mocks base method
func (m *MockSchedulerAPI) GetScheduleGroup(arg0 *scheduler.GetScheduleGroupInput) (*scheduler.GetScheduleGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduleGroup", arg0)
	ret0, _ := ret[0].(*scheduler.GetScheduleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduleGroup indicates an expected call of GetScheduleGroup
func (mr *MockSchedulerAPIMockRecorder) GetScheduleGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduleGroup", reflect.TypeOf((*MockSchedulerAPI)(nil).GetScheduleGroup), arg0)
}

// GetScheduleGroupRequest mocks base method
func (m *MockSchedulerAPI) GetScheduleGroupRequest(arg0 *scheduler.GetScheduleGroupInput) (*request.Request, *scheduler.GetScheduleGroupOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduleGroupRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.GetScheduleGroupOutput)
	return ret0, ret1
}

// GetScheduleGroupRequest indicates an expected call of GetScheduleGroupRequest
func (mr *MockSchedulerAPIMockRecorder) GetScheduleGroupRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduleGroupRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).GetScheduleGroupRequest), arg0)
}

// GetScheduleGroupWithContext mocks base method
func (m *MockSchedulerAPI) GetScheduleGroupWithContext(arg0 context.Context, arg1 *scheduler.GetScheduleGroupInput, arg2 ...request.Option) (*scheduler.GetScheduleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScheduleGroupWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.GetScheduleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduleGroupWithContext indicates an expected call of GetScheduleGroupWithContext
func (mr *MockSchedulerAPIMockRecorder) GetScheduleGroupWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduleGroupWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).GetScheduleGroupWithContext), varargs...)
}

// GetScheduleRequest mocks base method
func (m *MockSchedulerAPI) GetScheduleRequest(arg0 *scheduler.GetScheduleInput) (*request.Request, *scheduler.GetScheduleOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduleRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.GetScheduleOutput)
	return ret0, ret1
}

// GetScheduleRequest indicates an expected call of GetScheduleRequest
func (mr *MockSchedulerAPIMockRecorder) GetScheduleRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduleRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).GetScheduleRequest), arg0)
}

// GetScheduleWithContext mocks base method
func (m *MockSchedulerAPI) GetScheduleWithContext(arg0 context.Context, arg1 *scheduler.GetScheduleInput, arg2 ...request.Option) (*scheduler.GetScheduleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScheduleWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.GetScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduleWithContext indicates an expected call of GetScheduleWithContext
func (mr *MockSchedulerAPIMockRecorder) GetScheduleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduleWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).GetScheduleWithContext), varargs...)
}

// ListScheduleGroups mocks base method
func (m *MockSchedulerAPI) ListScheduleGroups(arg0 *scheduler.ListScheduleGroupsInput) (*scheduler.ListScheduleGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListScheduleGroups", arg0)
	ret0, _ := ret[0].(*scheduler.ListScheduleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListScheduleGroups indicates an expected call of ListScheduleGroups
func (mr *MockSchedulerAPIMockRecorder) ListScheduleGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListScheduleGroups", reflect.TypeOf((*MockSchedulerAPI)(nil).ListScheduleGroups), arg0)
}

// ListScheduleGroupsPages mocks base method
func (m *MockSchedulerAPI) ListScheduleGroupsPages(arg0 *scheduler.ListScheduleGroupsInput, arg1 func(*scheduler.ListScheduleGroupsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListScheduleGroupsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListScheduleGroupsPages indicates an expected call of ListScheduleGroupsPages
func (mr *MockSchedulerAPIMockRecorder) ListScheduleGroupsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListScheduleGroupsPages", reflect.TypeOf((*MockSchedulerAPI)(nil).ListScheduleGroupsPages), arg0, arg1)
}

// ListScheduleGroupsPagesWithContext mocks base method
func (m *MockSchedulerAPI) ListScheduleGroupsPagesWithContext(arg0 context.Context, arg1 *scheduler.ListScheduleGroupsInput, arg2 func(*scheduler.ListScheduleGroupsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListScheduleGroupsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListScheduleGroupsPagesWithContext indicates an expected call of ListScheduleGroupsPagesWithContext
func (mr *MockSchedulerAPIMockRecorder) ListScheduleGroupsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListScheduleGroupsPagesWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).ListScheduleGroupsPagesWithContext), varargs...)
}

// ListScheduleGroupsRequest mocks base method
func (m *MockSchedulerAPI) ListScheduleGroupsRequest(arg0 *scheduler.ListScheduleGroupsInput) (*request.Request, *scheduler.ListScheduleGroupsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListScheduleGroupsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.ListScheduleGroupsOutput)
	return ret0, ret1
}

// ListScheduleGroupsRequest indicates an expected call of ListScheduleGroupsRequest
func (mr *MockSchedulerAPIMockRecorder) ListScheduleGroupsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListScheduleGroupsRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).ListScheduleGroupsRequest), arg0)
}

// ListScheduleGroupsWithContext mocks base method
func (m *MockSchedulerAPI) ListScheduleGroupsWithContext(arg0 context.Context, arg1 *scheduler.ListScheduleGroupsInput, arg2 ...request.Option) (*scheduler.ListScheduleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListScheduleGroupsWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.ListScheduleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListScheduleGroupsWithContext indicates an expected call of ListScheduleGroupsWithContext
func (mr *MockSchedulerAPIMockRecorder) ListScheduleGroupsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListScheduleGroupsWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).ListScheduleGroupsWithContext), varargs...)
}

// ListSchedules mocks base method
func (m *MockSchedulerAPI) ListSchedules(arg0 *scheduler.ListSchedulesInput) (*scheduler.ListSchedulesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedules", arg0)
	ret0, _ := ret[0].(*scheduler.ListSchedulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedules indicates an expected call of ListSchedules
func (mr *MockSchedulerAPIMockRecorder) ListSchedules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedules", reflect.TypeOf((*MockSchedulerAPI)(nil).ListSchedules), arg0)
}

// ListSchedulesPages mocks base method
func (m *MockSchedulerAPI) ListSchedulesPages(arg0 *scheduler.ListSchedulesInput, arg1 func(*scheduler.ListSchedulesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListSchedulesPages indicates an expected call of ListSchedulesPages
func (mr *MockSchedulerAPIMockRecorder) ListSchedulesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulesPages", reflect.TypeOf((*MockSchedulerAPI)(nil).ListSchedulesPages), arg0, arg1)
}

// ListSchedulesPagesWithContext mocks base method
func (m *MockSchedulerAPI) ListSchedulesPagesWithContext(arg0 context.Context, arg1 *scheduler.ListSchedulesInput, arg2 func(*scheduler.ListSchedulesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSchedulesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListSchedulesPagesWithContext indicates an expected call of ListSchedulesPagesWithContext
func (mr *MockSchedulerAPIMockRecorder) ListSchedulesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulesPagesWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).ListSchedulesPagesWithContext), varargs...)
}

// ListSchedulesRequest mocks base method
func (m *MockSchedulerAPI) ListSchedulesRequest(arg0 *scheduler.ListSchedulesInput) (*request.Request, *scheduler.ListSchedulesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.ListSchedulesOutput)
	return ret0, ret1
}

// ListSchedulesRequest indicates an expected call of ListSchedulesRequest
func (mr *MockSchedulerAPIMockRecorder) ListSchedulesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulesRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).ListSchedulesRequest), arg0)
}

// ListSchedulesWithContext mocks base method
func (m *MockSchedulerAPI) ListSchedulesWithContext(arg0 context.Context, arg1 *scheduler.ListSchedulesInput, arg2 ...request.Option) (*scheduler.ListSchedulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSchedulesWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.ListSchedulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulesWithContext indicates an expected call of ListSchedulesWithContext
func (mr *MockSchedulerAPIMockRecorder) ListSchedulesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulesWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).ListSchedulesWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockSchedulerAPI) ListTagsForResource(arg0 *scheduler.ListTagsForResourceInput) (*scheduler.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*scheduler.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockSchedulerAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockSchedulerAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockSchedulerAPI) ListTagsForResourceRequest(arg0 *scheduler.ListTagsForResourceInput) (*request.Request, *scheduler.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockSchedulerAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockSchedulerAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *scheduler.ListTagsForResourceInput, arg2 ...request.Option) (*scheduler.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockSchedulerAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockSchedulerAPI) TagResource(arg0 *scheduler.TagResourceInput) (*scheduler.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*scheduler.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockSchedulerAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockSchedulerAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockSchedulerAPI) TagResourceRequest(arg0 *scheduler.TagResourceInput) (*request.Request, *scheduler.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockSchedulerAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockSchedulerAPI) TagResourceWithContext(arg0 context.Context, arg1 *scheduler.TagResourceInput, arg2 ...request.Option) (*scheduler.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockSchedulerAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockSchedulerAPI) UntagResource(arg0 *scheduler.UntagResourceInput) (*scheduler.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*scheduler.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockSchedulerAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockSchedulerAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockSchedulerAPI) UntagResourceRequest(arg0 *scheduler.UntagResourceInput) (*request.Request, *scheduler.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockSchedulerAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockSchedulerAPI) UntagResourceWithContext(arg0 context.Context, arg1 *scheduler.UntagResourceInput, arg2 ...request.Option) (*scheduler.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockSchedulerAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).UntagResourceWithContext), varargs...)
}

// UpdateSchedule mocks base method
func (m *MockSchedulerAPI) UpdateSchedule(arg0 *scheduler.UpdateScheduleInput) (*scheduler.UpdateScheduleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchedule", arg0)
	ret0, _ := ret[0].(*scheduler.UpdateScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSchedule indicates an expected call of UpdateSchedule
func (mr *MockSchedulerAPIMockRecorder) UpdateSchedule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchedule", reflect.TypeOf((*MockSchedulerAPI)(nil).UpdateSchedule), arg0)
}

// UpdateScheduleRequest mocks base method
func (m *MockSchedulerAPI) UpdateScheduleRequest(arg0 *scheduler.UpdateScheduleInput) (*request.Request, *scheduler.UpdateScheduleOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduleRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*scheduler.UpdateScheduleOutput)
	return ret0, ret1
}

// UpdateScheduleRequest indicates an expected call of UpdateScheduleRequest
func (mr *MockSchedulerAPIMockRecorder) UpdateScheduleRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduleRequest", reflect.TypeOf((*MockSchedulerAPI)(nil).UpdateScheduleRequest), arg0)
}

// UpdateScheduleWithContext mocks base method
func (m *MockSchedulerAPI) UpdateScheduleWithContext(arg0 context.Context, arg1 *scheduler.UpdateScheduleInput, arg2 ...request.Option) (*scheduler.UpdateScheduleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateScheduleWithContext", varargs...)
	ret0, _ := ret[0].(*scheduler.UpdateScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduleWithContext indicates an expected call of UpdateScheduleWithContext
func (mr *MockSchedulerAPIMockRecorder) UpdateScheduleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduleWithContext", reflect.TypeOf((*MockSchedulerAPI)(nil).UpdateScheduleWithContext), varargs...)
}
