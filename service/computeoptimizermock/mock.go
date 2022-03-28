// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/computeoptimizer/computeoptimizeriface (interfaces: ComputeOptimizerAPI)

// Package computeoptimizermock is a generated GoMock package.
package computeoptimizermock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	computeoptimizer "github.com/aws/aws-sdk-go/service/computeoptimizer"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockComputeOptimizerAPI is a mock of ComputeOptimizerAPI interface
type MockComputeOptimizerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockComputeOptimizerAPIMockRecorder
}

// MockComputeOptimizerAPIMockRecorder is the mock recorder for MockComputeOptimizerAPI
type MockComputeOptimizerAPIMockRecorder struct {
	mock *MockComputeOptimizerAPI
}

// NewMockComputeOptimizerAPI creates a new mock instance
func NewMockComputeOptimizerAPI(ctrl *gomock.Controller) *MockComputeOptimizerAPI {
	mock := &MockComputeOptimizerAPI{ctrl: ctrl}
	mock.recorder = &MockComputeOptimizerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockComputeOptimizerAPI) EXPECT() *MockComputeOptimizerAPIMockRecorder {
	return m.recorder
}

// GetAutoScalingGroupRecommendations mocks base method
func (m *MockComputeOptimizerAPI) GetAutoScalingGroupRecommendations(arg0 *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutoScalingGroupRecommendations", arg0)
	ret0, _ := ret[0].(*computeoptimizer.GetAutoScalingGroupRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutoScalingGroupRecommendations indicates an expected call of GetAutoScalingGroupRecommendations
func (mr *MockComputeOptimizerAPIMockRecorder) GetAutoScalingGroupRecommendations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoScalingGroupRecommendations", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetAutoScalingGroupRecommendations), arg0)
}

// GetAutoScalingGroupRecommendationsRequest mocks base method
func (m *MockComputeOptimizerAPI) GetAutoScalingGroupRecommendationsRequest(arg0 *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*request.Request, *computeoptimizer.GetAutoScalingGroupRecommendationsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutoScalingGroupRecommendationsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*computeoptimizer.GetAutoScalingGroupRecommendationsOutput)
	return ret0, ret1
}

// GetAutoScalingGroupRecommendationsRequest indicates an expected call of GetAutoScalingGroupRecommendationsRequest
func (mr *MockComputeOptimizerAPIMockRecorder) GetAutoScalingGroupRecommendationsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoScalingGroupRecommendationsRequest", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetAutoScalingGroupRecommendationsRequest), arg0)
}

// GetAutoScalingGroupRecommendationsWithContext mocks base method
func (m *MockComputeOptimizerAPI) GetAutoScalingGroupRecommendationsWithContext(arg0 context.Context, arg1 *computeoptimizer.GetAutoScalingGroupRecommendationsInput, arg2 ...request.Option) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAutoScalingGroupRecommendationsWithContext", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetAutoScalingGroupRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutoScalingGroupRecommendationsWithContext indicates an expected call of GetAutoScalingGroupRecommendationsWithContext
func (mr *MockComputeOptimizerAPIMockRecorder) GetAutoScalingGroupRecommendationsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoScalingGroupRecommendationsWithContext", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetAutoScalingGroupRecommendationsWithContext), varargs...)
}

// GetEC2InstanceRecommendations mocks base method
func (m *MockComputeOptimizerAPI) GetEC2InstanceRecommendations(arg0 *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEC2InstanceRecommendations", arg0)
	ret0, _ := ret[0].(*computeoptimizer.GetEC2InstanceRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEC2InstanceRecommendations indicates an expected call of GetEC2InstanceRecommendations
func (mr *MockComputeOptimizerAPIMockRecorder) GetEC2InstanceRecommendations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2InstanceRecommendations", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEC2InstanceRecommendations), arg0)
}

// GetEC2InstanceRecommendationsRequest mocks base method
func (m *MockComputeOptimizerAPI) GetEC2InstanceRecommendationsRequest(arg0 *computeoptimizer.GetEC2InstanceRecommendationsInput) (*request.Request, *computeoptimizer.GetEC2InstanceRecommendationsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEC2InstanceRecommendationsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*computeoptimizer.GetEC2InstanceRecommendationsOutput)
	return ret0, ret1
}

// GetEC2InstanceRecommendationsRequest indicates an expected call of GetEC2InstanceRecommendationsRequest
func (mr *MockComputeOptimizerAPIMockRecorder) GetEC2InstanceRecommendationsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2InstanceRecommendationsRequest", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEC2InstanceRecommendationsRequest), arg0)
}

// GetEC2InstanceRecommendationsWithContext mocks base method
func (m *MockComputeOptimizerAPI) GetEC2InstanceRecommendationsWithContext(arg0 context.Context, arg1 *computeoptimizer.GetEC2InstanceRecommendationsInput, arg2 ...request.Option) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEC2InstanceRecommendationsWithContext", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEC2InstanceRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEC2InstanceRecommendationsWithContext indicates an expected call of GetEC2InstanceRecommendationsWithContext
func (mr *MockComputeOptimizerAPIMockRecorder) GetEC2InstanceRecommendationsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2InstanceRecommendationsWithContext", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEC2InstanceRecommendationsWithContext), varargs...)
}

// GetEC2RecommendationProjectedMetrics mocks base method
func (m *MockComputeOptimizerAPI) GetEC2RecommendationProjectedMetrics(arg0 *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEC2RecommendationProjectedMetrics", arg0)
	ret0, _ := ret[0].(*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEC2RecommendationProjectedMetrics indicates an expected call of GetEC2RecommendationProjectedMetrics
func (mr *MockComputeOptimizerAPIMockRecorder) GetEC2RecommendationProjectedMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2RecommendationProjectedMetrics", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEC2RecommendationProjectedMetrics), arg0)
}

// GetEC2RecommendationProjectedMetricsRequest mocks base method
func (m *MockComputeOptimizerAPI) GetEC2RecommendationProjectedMetricsRequest(arg0 *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*request.Request, *computeoptimizer.GetEC2RecommendationProjectedMetricsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEC2RecommendationProjectedMetricsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput)
	return ret0, ret1
}

// GetEC2RecommendationProjectedMetricsRequest indicates an expected call of GetEC2RecommendationProjectedMetricsRequest
func (mr *MockComputeOptimizerAPIMockRecorder) GetEC2RecommendationProjectedMetricsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2RecommendationProjectedMetricsRequest", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEC2RecommendationProjectedMetricsRequest), arg0)
}

// GetEC2RecommendationProjectedMetricsWithContext mocks base method
func (m *MockComputeOptimizerAPI) GetEC2RecommendationProjectedMetricsWithContext(arg0 context.Context, arg1 *computeoptimizer.GetEC2RecommendationProjectedMetricsInput, arg2 ...request.Option) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEC2RecommendationProjectedMetricsWithContext", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEC2RecommendationProjectedMetricsWithContext indicates an expected call of GetEC2RecommendationProjectedMetricsWithContext
func (mr *MockComputeOptimizerAPIMockRecorder) GetEC2RecommendationProjectedMetricsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2RecommendationProjectedMetricsWithContext", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEC2RecommendationProjectedMetricsWithContext), varargs...)
}

// GetEnrollmentStatus mocks base method
func (m *MockComputeOptimizerAPI) GetEnrollmentStatus(arg0 *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnrollmentStatus", arg0)
	ret0, _ := ret[0].(*computeoptimizer.GetEnrollmentStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnrollmentStatus indicates an expected call of GetEnrollmentStatus
func (mr *MockComputeOptimizerAPIMockRecorder) GetEnrollmentStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnrollmentStatus", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEnrollmentStatus), arg0)
}

// GetEnrollmentStatusRequest mocks base method
func (m *MockComputeOptimizerAPI) GetEnrollmentStatusRequest(arg0 *computeoptimizer.GetEnrollmentStatusInput) (*request.Request, *computeoptimizer.GetEnrollmentStatusOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnrollmentStatusRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*computeoptimizer.GetEnrollmentStatusOutput)
	return ret0, ret1
}

// GetEnrollmentStatusRequest indicates an expected call of GetEnrollmentStatusRequest
func (mr *MockComputeOptimizerAPIMockRecorder) GetEnrollmentStatusRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnrollmentStatusRequest", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEnrollmentStatusRequest), arg0)
}

// GetEnrollmentStatusWithContext mocks base method
func (m *MockComputeOptimizerAPI) GetEnrollmentStatusWithContext(arg0 context.Context, arg1 *computeoptimizer.GetEnrollmentStatusInput, arg2 ...request.Option) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnrollmentStatusWithContext", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEnrollmentStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnrollmentStatusWithContext indicates an expected call of GetEnrollmentStatusWithContext
func (mr *MockComputeOptimizerAPIMockRecorder) GetEnrollmentStatusWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnrollmentStatusWithContext", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetEnrollmentStatusWithContext), varargs...)
}

// GetRecommendationSummaries mocks base method
func (m *MockComputeOptimizerAPI) GetRecommendationSummaries(arg0 *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendationSummaries", arg0)
	ret0, _ := ret[0].(*computeoptimizer.GetRecommendationSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendationSummaries indicates an expected call of GetRecommendationSummaries
func (mr *MockComputeOptimizerAPIMockRecorder) GetRecommendationSummaries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationSummaries", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetRecommendationSummaries), arg0)
}

// GetRecommendationSummariesRequest mocks base method
func (m *MockComputeOptimizerAPI) GetRecommendationSummariesRequest(arg0 *computeoptimizer.GetRecommendationSummariesInput) (*request.Request, *computeoptimizer.GetRecommendationSummariesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendationSummariesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*computeoptimizer.GetRecommendationSummariesOutput)
	return ret0, ret1
}

// GetRecommendationSummariesRequest indicates an expected call of GetRecommendationSummariesRequest
func (mr *MockComputeOptimizerAPIMockRecorder) GetRecommendationSummariesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationSummariesRequest", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetRecommendationSummariesRequest), arg0)
}

// GetRecommendationSummariesWithContext mocks base method
func (m *MockComputeOptimizerAPI) GetRecommendationSummariesWithContext(arg0 context.Context, arg1 *computeoptimizer.GetRecommendationSummariesInput, arg2 ...request.Option) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecommendationSummariesWithContext", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetRecommendationSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendationSummariesWithContext indicates an expected call of GetRecommendationSummariesWithContext
func (mr *MockComputeOptimizerAPIMockRecorder) GetRecommendationSummariesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationSummariesWithContext", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).GetRecommendationSummariesWithContext), varargs...)
}

// UpdateEnrollmentStatus mocks base method
func (m *MockComputeOptimizerAPI) UpdateEnrollmentStatus(arg0 *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEnrollmentStatus", arg0)
	ret0, _ := ret[0].(*computeoptimizer.UpdateEnrollmentStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEnrollmentStatus indicates an expected call of UpdateEnrollmentStatus
func (mr *MockComputeOptimizerAPIMockRecorder) UpdateEnrollmentStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnrollmentStatus", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).UpdateEnrollmentStatus), arg0)
}

// UpdateEnrollmentStatusRequest mocks base method
func (m *MockComputeOptimizerAPI) UpdateEnrollmentStatusRequest(arg0 *computeoptimizer.UpdateEnrollmentStatusInput) (*request.Request, *computeoptimizer.UpdateEnrollmentStatusOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEnrollmentStatusRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*computeoptimizer.UpdateEnrollmentStatusOutput)
	return ret0, ret1
}

// UpdateEnrollmentStatusRequest indicates an expected call of UpdateEnrollmentStatusRequest
func (mr *MockComputeOptimizerAPIMockRecorder) UpdateEnrollmentStatusRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnrollmentStatusRequest", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).UpdateEnrollmentStatusRequest), arg0)
}

// UpdateEnrollmentStatusWithContext mocks base method
func (m *MockComputeOptimizerAPI) UpdateEnrollmentStatusWithContext(arg0 context.Context, arg1 *computeoptimizer.UpdateEnrollmentStatusInput, arg2 ...request.Option) (*computeoptimizer.UpdateEnrollmentStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEnrollmentStatusWithContext", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.UpdateEnrollmentStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEnrollmentStatusWithContext indicates an expected call of UpdateEnrollmentStatusWithContext
func (mr *MockComputeOptimizerAPIMockRecorder) UpdateEnrollmentStatusWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnrollmentStatusWithContext", reflect.TypeOf((*MockComputeOptimizerAPI)(nil).UpdateEnrollmentStatusWithContext), varargs...)
}