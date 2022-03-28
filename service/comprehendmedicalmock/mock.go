// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface (interfaces: ComprehendMedicalAPI)

// Package comprehendmedicalmock is a generated GoMock package.
package comprehendmedicalmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	comprehendmedical "github.com/aws/aws-sdk-go/service/comprehendmedical"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockComprehendMedicalAPI is a mock of ComprehendMedicalAPI interface
type MockComprehendMedicalAPI struct {
	ctrl     *gomock.Controller
	recorder *MockComprehendMedicalAPIMockRecorder
}

// MockComprehendMedicalAPIMockRecorder is the mock recorder for MockComprehendMedicalAPI
type MockComprehendMedicalAPIMockRecorder struct {
	mock *MockComprehendMedicalAPI
}

// NewMockComprehendMedicalAPI creates a new mock instance
func NewMockComprehendMedicalAPI(ctrl *gomock.Controller) *MockComprehendMedicalAPI {
	mock := &MockComprehendMedicalAPI{ctrl: ctrl}
	mock.recorder = &MockComprehendMedicalAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockComprehendMedicalAPI) EXPECT() *MockComprehendMedicalAPIMockRecorder {
	return m.recorder
}

// DetectEntities mocks base method
func (m *MockComprehendMedicalAPI) DetectEntities(arg0 *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectEntities", arg0)
	ret0, _ := ret[0].(*comprehendmedical.DetectEntitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectEntities indicates an expected call of DetectEntities
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntities", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntities), arg0)
}

// DetectEntitiesRequest mocks base method
func (m *MockComprehendMedicalAPI) DetectEntitiesRequest(arg0 *comprehendmedical.DetectEntitiesInput) (*request.Request, *comprehendmedical.DetectEntitiesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectEntitiesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.DetectEntitiesOutput)
	return ret0, ret1
}

// DetectEntitiesRequest indicates an expected call of DetectEntitiesRequest
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntitiesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntitiesRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntitiesRequest), arg0)
}

// DetectEntitiesWithContext mocks base method
func (m *MockComprehendMedicalAPI) DetectEntitiesWithContext(arg0 context.Context, arg1 *comprehendmedical.DetectEntitiesInput, arg2 ...request.Option) (*comprehendmedical.DetectEntitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DetectEntitiesWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.DetectEntitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectEntitiesWithContext indicates an expected call of DetectEntitiesWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntitiesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntitiesWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntitiesWithContext), varargs...)
}

// DetectPHI mocks base method
func (m *MockComprehendMedicalAPI) DetectPHI(arg0 *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectPHI", arg0)
	ret0, _ := ret[0].(*comprehendmedical.DetectPHIOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectPHI indicates an expected call of DetectPHI
func (mr *MockComprehendMedicalAPIMockRecorder) DetectPHI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectPHI", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectPHI), arg0)
}

// DetectPHIRequest mocks base method
func (m *MockComprehendMedicalAPI) DetectPHIRequest(arg0 *comprehendmedical.DetectPHIInput) (*request.Request, *comprehendmedical.DetectPHIOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectPHIRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.DetectPHIOutput)
	return ret0, ret1
}

// DetectPHIRequest indicates an expected call of DetectPHIRequest
func (mr *MockComprehendMedicalAPIMockRecorder) DetectPHIRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectPHIRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectPHIRequest), arg0)
}

// DetectPHIWithContext mocks base method
func (m *MockComprehendMedicalAPI) DetectPHIWithContext(arg0 context.Context, arg1 *comprehendmedical.DetectPHIInput, arg2 ...request.Option) (*comprehendmedical.DetectPHIOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DetectPHIWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.DetectPHIOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectPHIWithContext indicates an expected call of DetectPHIWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) DetectPHIWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectPHIWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectPHIWithContext), varargs...)
}