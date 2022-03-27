// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/iot1clickprojects/iot1clickprojectsiface (interfaces: IoT1ClickProjectsAPI)

// Package iot1clickprojectsmock is a generated GoMock package.
package iot1clickprojectsmock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	iot1clickprojects "github.com/aws/aws-sdk-go/service/iot1clickprojects"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIoT1ClickProjectsAPI is a mock of IoT1ClickProjectsAPI interface
type MockIoT1ClickProjectsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIoT1ClickProjectsAPIMockRecorder
}

// MockIoT1ClickProjectsAPIMockRecorder is the mock recorder for MockIoT1ClickProjectsAPI
type MockIoT1ClickProjectsAPIMockRecorder struct {
	mock *MockIoT1ClickProjectsAPI
}

// NewMockIoT1ClickProjectsAPI creates a new mock instance
func NewMockIoT1ClickProjectsAPI(ctrl *gomock.Controller) *MockIoT1ClickProjectsAPI {
	mock := &MockIoT1ClickProjectsAPI{ctrl: ctrl}
	mock.recorder = &MockIoT1ClickProjectsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIoT1ClickProjectsAPI) EXPECT() *MockIoT1ClickProjectsAPIMockRecorder {
	return m.recorder
}

// AssociateDeviceWithPlacement mocks base method
func (m *MockIoT1ClickProjectsAPI) AssociateDeviceWithPlacement(arg0 *iot1clickprojects.AssociateDeviceWithPlacementInput) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateDeviceWithPlacement", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.AssociateDeviceWithPlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssociateDeviceWithPlacement indicates an expected call of AssociateDeviceWithPlacement
func (mr *MockIoT1ClickProjectsAPIMockRecorder) AssociateDeviceWithPlacement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateDeviceWithPlacement", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).AssociateDeviceWithPlacement), arg0)
}

// AssociateDeviceWithPlacementRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) AssociateDeviceWithPlacementRequest(arg0 *iot1clickprojects.AssociateDeviceWithPlacementInput) (*request.Request, *iot1clickprojects.AssociateDeviceWithPlacementOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateDeviceWithPlacementRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.AssociateDeviceWithPlacementOutput)
	return ret0, ret1
}

// AssociateDeviceWithPlacementRequest indicates an expected call of AssociateDeviceWithPlacementRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) AssociateDeviceWithPlacementRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateDeviceWithPlacementRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).AssociateDeviceWithPlacementRequest), arg0)
}

// AssociateDeviceWithPlacementWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) AssociateDeviceWithPlacementWithContext(arg0 aws.Context, arg1 *iot1clickprojects.AssociateDeviceWithPlacementInput, arg2 ...request.Option) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssociateDeviceWithPlacementWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.AssociateDeviceWithPlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssociateDeviceWithPlacementWithContext indicates an expected call of AssociateDeviceWithPlacementWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) AssociateDeviceWithPlacementWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateDeviceWithPlacementWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).AssociateDeviceWithPlacementWithContext), varargs...)
}

// CreatePlacement mocks base method
func (m *MockIoT1ClickProjectsAPI) CreatePlacement(arg0 *iot1clickprojects.CreatePlacementInput) (*iot1clickprojects.CreatePlacementOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlacement", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.CreatePlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlacement indicates an expected call of CreatePlacement
func (mr *MockIoT1ClickProjectsAPIMockRecorder) CreatePlacement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlacement", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).CreatePlacement), arg0)
}

// CreatePlacementRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) CreatePlacementRequest(arg0 *iot1clickprojects.CreatePlacementInput) (*request.Request, *iot1clickprojects.CreatePlacementOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlacementRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.CreatePlacementOutput)
	return ret0, ret1
}

// CreatePlacementRequest indicates an expected call of CreatePlacementRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) CreatePlacementRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlacementRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).CreatePlacementRequest), arg0)
}

// CreatePlacementWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) CreatePlacementWithContext(arg0 aws.Context, arg1 *iot1clickprojects.CreatePlacementInput, arg2 ...request.Option) (*iot1clickprojects.CreatePlacementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePlacementWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.CreatePlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlacementWithContext indicates an expected call of CreatePlacementWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) CreatePlacementWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlacementWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).CreatePlacementWithContext), varargs...)
}

// CreateProject mocks base method
func (m *MockIoT1ClickProjectsAPI) CreateProject(arg0 *iot1clickprojects.CreateProjectInput) (*iot1clickprojects.CreateProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.CreateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject
func (mr *MockIoT1ClickProjectsAPIMockRecorder) CreateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).CreateProject), arg0)
}

// CreateProjectRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) CreateProjectRequest(arg0 *iot1clickprojects.CreateProjectInput) (*request.Request, *iot1clickprojects.CreateProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.CreateProjectOutput)
	return ret0, ret1
}

// CreateProjectRequest indicates an expected call of CreateProjectRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) CreateProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).CreateProjectRequest), arg0)
}

// CreateProjectWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) CreateProjectWithContext(arg0 aws.Context, arg1 *iot1clickprojects.CreateProjectInput, arg2 ...request.Option) (*iot1clickprojects.CreateProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProjectWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.CreateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectWithContext indicates an expected call of CreateProjectWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) CreateProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).CreateProjectWithContext), varargs...)
}

// DeletePlacement mocks base method
func (m *MockIoT1ClickProjectsAPI) DeletePlacement(arg0 *iot1clickprojects.DeletePlacementInput) (*iot1clickprojects.DeletePlacementOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlacement", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.DeletePlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePlacement indicates an expected call of DeletePlacement
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DeletePlacement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlacement", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DeletePlacement), arg0)
}

// DeletePlacementRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) DeletePlacementRequest(arg0 *iot1clickprojects.DeletePlacementInput) (*request.Request, *iot1clickprojects.DeletePlacementOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlacementRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.DeletePlacementOutput)
	return ret0, ret1
}

// DeletePlacementRequest indicates an expected call of DeletePlacementRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DeletePlacementRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlacementRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DeletePlacementRequest), arg0)
}

// DeletePlacementWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) DeletePlacementWithContext(arg0 aws.Context, arg1 *iot1clickprojects.DeletePlacementInput, arg2 ...request.Option) (*iot1clickprojects.DeletePlacementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePlacementWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.DeletePlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePlacementWithContext indicates an expected call of DeletePlacementWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DeletePlacementWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlacementWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DeletePlacementWithContext), varargs...)
}

// DeleteProject mocks base method
func (m *MockIoT1ClickProjectsAPI) DeleteProject(arg0 *iot1clickprojects.DeleteProjectInput) (*iot1clickprojects.DeleteProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.DeleteProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DeleteProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DeleteProject), arg0)
}

// DeleteProjectRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) DeleteProjectRequest(arg0 *iot1clickprojects.DeleteProjectInput) (*request.Request, *iot1clickprojects.DeleteProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.DeleteProjectOutput)
	return ret0, ret1
}

// DeleteProjectRequest indicates an expected call of DeleteProjectRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DeleteProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DeleteProjectRequest), arg0)
}

// DeleteProjectWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) DeleteProjectWithContext(arg0 aws.Context, arg1 *iot1clickprojects.DeleteProjectInput, arg2 ...request.Option) (*iot1clickprojects.DeleteProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProjectWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.DeleteProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectWithContext indicates an expected call of DeleteProjectWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DeleteProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DeleteProjectWithContext), varargs...)
}

// DescribePlacement mocks base method
func (m *MockIoT1ClickProjectsAPI) DescribePlacement(arg0 *iot1clickprojects.DescribePlacementInput) (*iot1clickprojects.DescribePlacementOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribePlacement", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.DescribePlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePlacement indicates an expected call of DescribePlacement
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DescribePlacement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePlacement", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DescribePlacement), arg0)
}

// DescribePlacementRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) DescribePlacementRequest(arg0 *iot1clickprojects.DescribePlacementInput) (*request.Request, *iot1clickprojects.DescribePlacementOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribePlacementRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.DescribePlacementOutput)
	return ret0, ret1
}

// DescribePlacementRequest indicates an expected call of DescribePlacementRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DescribePlacementRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePlacementRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DescribePlacementRequest), arg0)
}

// DescribePlacementWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) DescribePlacementWithContext(arg0 aws.Context, arg1 *iot1clickprojects.DescribePlacementInput, arg2 ...request.Option) (*iot1clickprojects.DescribePlacementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePlacementWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.DescribePlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePlacementWithContext indicates an expected call of DescribePlacementWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DescribePlacementWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePlacementWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DescribePlacementWithContext), varargs...)
}

// DescribeProject mocks base method
func (m *MockIoT1ClickProjectsAPI) DescribeProject(arg0 *iot1clickprojects.DescribeProjectInput) (*iot1clickprojects.DescribeProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeProject", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.DescribeProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProject indicates an expected call of DescribeProject
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DescribeProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProject", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DescribeProject), arg0)
}

// DescribeProjectRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) DescribeProjectRequest(arg0 *iot1clickprojects.DescribeProjectInput) (*request.Request, *iot1clickprojects.DescribeProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.DescribeProjectOutput)
	return ret0, ret1
}

// DescribeProjectRequest indicates an expected call of DescribeProjectRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DescribeProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProjectRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DescribeProjectRequest), arg0)
}

// DescribeProjectWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) DescribeProjectWithContext(arg0 aws.Context, arg1 *iot1clickprojects.DescribeProjectInput, arg2 ...request.Option) (*iot1clickprojects.DescribeProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeProjectWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.DescribeProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProjectWithContext indicates an expected call of DescribeProjectWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DescribeProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProjectWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DescribeProjectWithContext), varargs...)
}

// DisassociateDeviceFromPlacement mocks base method
func (m *MockIoT1ClickProjectsAPI) DisassociateDeviceFromPlacement(arg0 *iot1clickprojects.DisassociateDeviceFromPlacementInput) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisassociateDeviceFromPlacement", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.DisassociateDeviceFromPlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisassociateDeviceFromPlacement indicates an expected call of DisassociateDeviceFromPlacement
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DisassociateDeviceFromPlacement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisassociateDeviceFromPlacement", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DisassociateDeviceFromPlacement), arg0)
}

// DisassociateDeviceFromPlacementRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) DisassociateDeviceFromPlacementRequest(arg0 *iot1clickprojects.DisassociateDeviceFromPlacementInput) (*request.Request, *iot1clickprojects.DisassociateDeviceFromPlacementOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisassociateDeviceFromPlacementRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.DisassociateDeviceFromPlacementOutput)
	return ret0, ret1
}

// DisassociateDeviceFromPlacementRequest indicates an expected call of DisassociateDeviceFromPlacementRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DisassociateDeviceFromPlacementRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisassociateDeviceFromPlacementRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DisassociateDeviceFromPlacementRequest), arg0)
}

// DisassociateDeviceFromPlacementWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) DisassociateDeviceFromPlacementWithContext(arg0 aws.Context, arg1 *iot1clickprojects.DisassociateDeviceFromPlacementInput, arg2 ...request.Option) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisassociateDeviceFromPlacementWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.DisassociateDeviceFromPlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisassociateDeviceFromPlacementWithContext indicates an expected call of DisassociateDeviceFromPlacementWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) DisassociateDeviceFromPlacementWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisassociateDeviceFromPlacementWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).DisassociateDeviceFromPlacementWithContext), varargs...)
}

// GetDevicesInPlacement mocks base method
func (m *MockIoT1ClickProjectsAPI) GetDevicesInPlacement(arg0 *iot1clickprojects.GetDevicesInPlacementInput) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesInPlacement", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.GetDevicesInPlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesInPlacement indicates an expected call of GetDevicesInPlacement
func (mr *MockIoT1ClickProjectsAPIMockRecorder) GetDevicesInPlacement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesInPlacement", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).GetDevicesInPlacement), arg0)
}

// GetDevicesInPlacementRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) GetDevicesInPlacementRequest(arg0 *iot1clickprojects.GetDevicesInPlacementInput) (*request.Request, *iot1clickprojects.GetDevicesInPlacementOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesInPlacementRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.GetDevicesInPlacementOutput)
	return ret0, ret1
}

// GetDevicesInPlacementRequest indicates an expected call of GetDevicesInPlacementRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) GetDevicesInPlacementRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesInPlacementRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).GetDevicesInPlacementRequest), arg0)
}

// GetDevicesInPlacementWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) GetDevicesInPlacementWithContext(arg0 aws.Context, arg1 *iot1clickprojects.GetDevicesInPlacementInput, arg2 ...request.Option) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDevicesInPlacementWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.GetDevicesInPlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesInPlacementWithContext indicates an expected call of GetDevicesInPlacementWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) GetDevicesInPlacementWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesInPlacementWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).GetDevicesInPlacementWithContext), varargs...)
}

// ListPlacements mocks base method
func (m *MockIoT1ClickProjectsAPI) ListPlacements(arg0 *iot1clickprojects.ListPlacementsInput) (*iot1clickprojects.ListPlacementsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlacements", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.ListPlacementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlacements indicates an expected call of ListPlacements
func (mr *MockIoT1ClickProjectsAPIMockRecorder) ListPlacements(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlacements", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).ListPlacements), arg0)
}

// ListPlacementsRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) ListPlacementsRequest(arg0 *iot1clickprojects.ListPlacementsInput) (*request.Request, *iot1clickprojects.ListPlacementsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlacementsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.ListPlacementsOutput)
	return ret0, ret1
}

// ListPlacementsRequest indicates an expected call of ListPlacementsRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) ListPlacementsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlacementsRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).ListPlacementsRequest), arg0)
}

// ListPlacementsWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) ListPlacementsWithContext(arg0 aws.Context, arg1 *iot1clickprojects.ListPlacementsInput, arg2 ...request.Option) (*iot1clickprojects.ListPlacementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPlacementsWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.ListPlacementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlacementsWithContext indicates an expected call of ListPlacementsWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) ListPlacementsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlacementsWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).ListPlacementsWithContext), varargs...)
}

// ListProjects mocks base method
func (m *MockIoT1ClickProjectsAPI) ListProjects(arg0 *iot1clickprojects.ListProjectsInput) (*iot1clickprojects.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects
func (mr *MockIoT1ClickProjectsAPIMockRecorder) ListProjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).ListProjects), arg0)
}

// ListProjectsRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) ListProjectsRequest(arg0 *iot1clickprojects.ListProjectsInput) (*request.Request, *iot1clickprojects.ListProjectsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.ListProjectsOutput)
	return ret0, ret1
}

// ListProjectsRequest indicates an expected call of ListProjectsRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) ListProjectsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).ListProjectsRequest), arg0)
}

// ListProjectsWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) ListProjectsWithContext(arg0 aws.Context, arg1 *iot1clickprojects.ListProjectsInput, arg2 ...request.Option) (*iot1clickprojects.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectsWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectsWithContext indicates an expected call of ListProjectsWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) ListProjectsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).ListProjectsWithContext), varargs...)
}

// UpdatePlacement mocks base method
func (m *MockIoT1ClickProjectsAPI) UpdatePlacement(arg0 *iot1clickprojects.UpdatePlacementInput) (*iot1clickprojects.UpdatePlacementOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlacement", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.UpdatePlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePlacement indicates an expected call of UpdatePlacement
func (mr *MockIoT1ClickProjectsAPIMockRecorder) UpdatePlacement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlacement", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).UpdatePlacement), arg0)
}

// UpdatePlacementRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) UpdatePlacementRequest(arg0 *iot1clickprojects.UpdatePlacementInput) (*request.Request, *iot1clickprojects.UpdatePlacementOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlacementRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.UpdatePlacementOutput)
	return ret0, ret1
}

// UpdatePlacementRequest indicates an expected call of UpdatePlacementRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) UpdatePlacementRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlacementRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).UpdatePlacementRequest), arg0)
}

// UpdatePlacementWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) UpdatePlacementWithContext(arg0 aws.Context, arg1 *iot1clickprojects.UpdatePlacementInput, arg2 ...request.Option) (*iot1clickprojects.UpdatePlacementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePlacementWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.UpdatePlacementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePlacementWithContext indicates an expected call of UpdatePlacementWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) UpdatePlacementWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlacementWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).UpdatePlacementWithContext), varargs...)
}

// UpdateProject mocks base method
func (m *MockIoT1ClickProjectsAPI) UpdateProject(arg0 *iot1clickprojects.UpdateProjectInput) (*iot1clickprojects.UpdateProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0)
	ret0, _ := ret[0].(*iot1clickprojects.UpdateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject
func (mr *MockIoT1ClickProjectsAPIMockRecorder) UpdateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).UpdateProject), arg0)
}

// UpdateProjectRequest mocks base method
func (m *MockIoT1ClickProjectsAPI) UpdateProjectRequest(arg0 *iot1clickprojects.UpdateProjectInput) (*request.Request, *iot1clickprojects.UpdateProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iot1clickprojects.UpdateProjectOutput)
	return ret0, ret1
}

// UpdateProjectRequest indicates an expected call of UpdateProjectRequest
func (mr *MockIoT1ClickProjectsAPIMockRecorder) UpdateProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectRequest", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).UpdateProjectRequest), arg0)
}

// UpdateProjectWithContext mocks base method
func (m *MockIoT1ClickProjectsAPI) UpdateProjectWithContext(arg0 aws.Context, arg1 *iot1clickprojects.UpdateProjectInput, arg2 ...request.Option) (*iot1clickprojects.UpdateProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProjectWithContext", varargs...)
	ret0, _ := ret[0].(*iot1clickprojects.UpdateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectWithContext indicates an expected call of UpdateProjectWithContext
func (mr *MockIoT1ClickProjectsAPIMockRecorder) UpdateProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectWithContext", reflect.TypeOf((*MockIoT1ClickProjectsAPI)(nil).UpdateProjectWithContext), varargs...)
}