// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface (interfaces: CostExplorerAPI)

// Package costexplorermock is a generated GoMock package.
package costexplorermock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	costexplorer "github.com/aws/aws-sdk-go/service/costexplorer"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCostExplorerAPI is a mock of CostExplorerAPI interface
type MockCostExplorerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCostExplorerAPIMockRecorder
}

// MockCostExplorerAPIMockRecorder is the mock recorder for MockCostExplorerAPI
type MockCostExplorerAPIMockRecorder struct {
	mock *MockCostExplorerAPI
}

// NewMockCostExplorerAPI creates a new mock instance
func NewMockCostExplorerAPI(ctrl *gomock.Controller) *MockCostExplorerAPI {
	mock := &MockCostExplorerAPI{ctrl: ctrl}
	mock.recorder = &MockCostExplorerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCostExplorerAPI) EXPECT() *MockCostExplorerAPIMockRecorder {
	return m.recorder
}

// GetCostAndUsage mocks base method
func (m *MockCostExplorerAPI) GetCostAndUsage(arg0 *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCostAndUsage", arg0)
	ret0, _ := ret[0].(*costexplorer.GetCostAndUsageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCostAndUsage indicates an expected call of GetCostAndUsage
func (mr *MockCostExplorerAPIMockRecorder) GetCostAndUsage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostAndUsage", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetCostAndUsage), arg0)
}

// GetCostAndUsageRequest mocks base method
func (m *MockCostExplorerAPI) GetCostAndUsageRequest(arg0 *costexplorer.GetCostAndUsageInput) (*request.Request, *costexplorer.GetCostAndUsageOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCostAndUsageRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costexplorer.GetCostAndUsageOutput)
	return ret0, ret1
}

// GetCostAndUsageRequest indicates an expected call of GetCostAndUsageRequest
func (mr *MockCostExplorerAPIMockRecorder) GetCostAndUsageRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostAndUsageRequest", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetCostAndUsageRequest), arg0)
}

// GetCostAndUsageWithContext mocks base method
func (m *MockCostExplorerAPI) GetCostAndUsageWithContext(arg0 context.Context, arg1 *costexplorer.GetCostAndUsageInput, arg2 ...request.Option) (*costexplorer.GetCostAndUsageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCostAndUsageWithContext", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetCostAndUsageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCostAndUsageWithContext indicates an expected call of GetCostAndUsageWithContext
func (mr *MockCostExplorerAPIMockRecorder) GetCostAndUsageWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostAndUsageWithContext", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetCostAndUsageWithContext), varargs...)
}

// GetCostForecast mocks base method
func (m *MockCostExplorerAPI) GetCostForecast(arg0 *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCostForecast", arg0)
	ret0, _ := ret[0].(*costexplorer.GetCostForecastOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCostForecast indicates an expected call of GetCostForecast
func (mr *MockCostExplorerAPIMockRecorder) GetCostForecast(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostForecast", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetCostForecast), arg0)
}

// GetCostForecastRequest mocks base method
func (m *MockCostExplorerAPI) GetCostForecastRequest(arg0 *costexplorer.GetCostForecastInput) (*request.Request, *costexplorer.GetCostForecastOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCostForecastRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costexplorer.GetCostForecastOutput)
	return ret0, ret1
}

// GetCostForecastRequest indicates an expected call of GetCostForecastRequest
func (mr *MockCostExplorerAPIMockRecorder) GetCostForecastRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostForecastRequest", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetCostForecastRequest), arg0)
}

// GetCostForecastWithContext mocks base method
func (m *MockCostExplorerAPI) GetCostForecastWithContext(arg0 context.Context, arg1 *costexplorer.GetCostForecastInput, arg2 ...request.Option) (*costexplorer.GetCostForecastOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCostForecastWithContext", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetCostForecastOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCostForecastWithContext indicates an expected call of GetCostForecastWithContext
func (mr *MockCostExplorerAPIMockRecorder) GetCostForecastWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostForecastWithContext", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetCostForecastWithContext), varargs...)
}

// GetDimensionValues mocks base method
func (m *MockCostExplorerAPI) GetDimensionValues(arg0 *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensionValues", arg0)
	ret0, _ := ret[0].(*costexplorer.GetDimensionValuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensionValues indicates an expected call of GetDimensionValues
func (mr *MockCostExplorerAPIMockRecorder) GetDimensionValues(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionValues", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetDimensionValues), arg0)
}

// GetDimensionValuesRequest mocks base method
func (m *MockCostExplorerAPI) GetDimensionValuesRequest(arg0 *costexplorer.GetDimensionValuesInput) (*request.Request, *costexplorer.GetDimensionValuesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensionValuesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costexplorer.GetDimensionValuesOutput)
	return ret0, ret1
}

// GetDimensionValuesRequest indicates an expected call of GetDimensionValuesRequest
func (mr *MockCostExplorerAPIMockRecorder) GetDimensionValuesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionValuesRequest", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetDimensionValuesRequest), arg0)
}

// GetDimensionValuesWithContext mocks base method
func (m *MockCostExplorerAPI) GetDimensionValuesWithContext(arg0 context.Context, arg1 *costexplorer.GetDimensionValuesInput, arg2 ...request.Option) (*costexplorer.GetDimensionValuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDimensionValuesWithContext", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetDimensionValuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensionValuesWithContext indicates an expected call of GetDimensionValuesWithContext
func (mr *MockCostExplorerAPIMockRecorder) GetDimensionValuesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionValuesWithContext", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetDimensionValuesWithContext), varargs...)
}

// GetReservationCoverage mocks base method
func (m *MockCostExplorerAPI) GetReservationCoverage(arg0 *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservationCoverage", arg0)
	ret0, _ := ret[0].(*costexplorer.GetReservationCoverageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationCoverage indicates an expected call of GetReservationCoverage
func (mr *MockCostExplorerAPIMockRecorder) GetReservationCoverage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationCoverage", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationCoverage), arg0)
}

// GetReservationCoverageRequest mocks base method
func (m *MockCostExplorerAPI) GetReservationCoverageRequest(arg0 *costexplorer.GetReservationCoverageInput) (*request.Request, *costexplorer.GetReservationCoverageOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservationCoverageRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costexplorer.GetReservationCoverageOutput)
	return ret0, ret1
}

// GetReservationCoverageRequest indicates an expected call of GetReservationCoverageRequest
func (mr *MockCostExplorerAPIMockRecorder) GetReservationCoverageRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationCoverageRequest", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationCoverageRequest), arg0)
}

// GetReservationCoverageWithContext mocks base method
func (m *MockCostExplorerAPI) GetReservationCoverageWithContext(arg0 context.Context, arg1 *costexplorer.GetReservationCoverageInput, arg2 ...request.Option) (*costexplorer.GetReservationCoverageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReservationCoverageWithContext", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetReservationCoverageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationCoverageWithContext indicates an expected call of GetReservationCoverageWithContext
func (mr *MockCostExplorerAPIMockRecorder) GetReservationCoverageWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationCoverageWithContext", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationCoverageWithContext), varargs...)
}

// GetReservationPurchaseRecommendation mocks base method
func (m *MockCostExplorerAPI) GetReservationPurchaseRecommendation(arg0 *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservationPurchaseRecommendation", arg0)
	ret0, _ := ret[0].(*costexplorer.GetReservationPurchaseRecommendationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationPurchaseRecommendation indicates an expected call of GetReservationPurchaseRecommendation
func (mr *MockCostExplorerAPIMockRecorder) GetReservationPurchaseRecommendation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationPurchaseRecommendation", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationPurchaseRecommendation), arg0)
}

// GetReservationPurchaseRecommendationRequest mocks base method
func (m *MockCostExplorerAPI) GetReservationPurchaseRecommendationRequest(arg0 *costexplorer.GetReservationPurchaseRecommendationInput) (*request.Request, *costexplorer.GetReservationPurchaseRecommendationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservationPurchaseRecommendationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costexplorer.GetReservationPurchaseRecommendationOutput)
	return ret0, ret1
}

// GetReservationPurchaseRecommendationRequest indicates an expected call of GetReservationPurchaseRecommendationRequest
func (mr *MockCostExplorerAPIMockRecorder) GetReservationPurchaseRecommendationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationPurchaseRecommendationRequest", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationPurchaseRecommendationRequest), arg0)
}

// GetReservationPurchaseRecommendationWithContext mocks base method
func (m *MockCostExplorerAPI) GetReservationPurchaseRecommendationWithContext(arg0 context.Context, arg1 *costexplorer.GetReservationPurchaseRecommendationInput, arg2 ...request.Option) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReservationPurchaseRecommendationWithContext", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetReservationPurchaseRecommendationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationPurchaseRecommendationWithContext indicates an expected call of GetReservationPurchaseRecommendationWithContext
func (mr *MockCostExplorerAPIMockRecorder) GetReservationPurchaseRecommendationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationPurchaseRecommendationWithContext", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationPurchaseRecommendationWithContext), varargs...)
}

// GetReservationUtilization mocks base method
func (m *MockCostExplorerAPI) GetReservationUtilization(arg0 *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservationUtilization", arg0)
	ret0, _ := ret[0].(*costexplorer.GetReservationUtilizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationUtilization indicates an expected call of GetReservationUtilization
func (mr *MockCostExplorerAPIMockRecorder) GetReservationUtilization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationUtilization", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationUtilization), arg0)
}

// GetReservationUtilizationRequest mocks base method
func (m *MockCostExplorerAPI) GetReservationUtilizationRequest(arg0 *costexplorer.GetReservationUtilizationInput) (*request.Request, *costexplorer.GetReservationUtilizationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservationUtilizationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costexplorer.GetReservationUtilizationOutput)
	return ret0, ret1
}

// GetReservationUtilizationRequest indicates an expected call of GetReservationUtilizationRequest
func (mr *MockCostExplorerAPIMockRecorder) GetReservationUtilizationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationUtilizationRequest", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationUtilizationRequest), arg0)
}

// GetReservationUtilizationWithContext mocks base method
func (m *MockCostExplorerAPI) GetReservationUtilizationWithContext(arg0 context.Context, arg1 *costexplorer.GetReservationUtilizationInput, arg2 ...request.Option) (*costexplorer.GetReservationUtilizationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReservationUtilizationWithContext", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetReservationUtilizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationUtilizationWithContext indicates an expected call of GetReservationUtilizationWithContext
func (mr *MockCostExplorerAPIMockRecorder) GetReservationUtilizationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationUtilizationWithContext", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetReservationUtilizationWithContext), varargs...)
}

// GetTags mocks base method
func (m *MockCostExplorerAPI) GetTags(arg0 *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags", arg0)
	ret0, _ := ret[0].(*costexplorer.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags
func (mr *MockCostExplorerAPIMockRecorder) GetTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetTags), arg0)
}

// GetTagsRequest mocks base method
func (m *MockCostExplorerAPI) GetTagsRequest(arg0 *costexplorer.GetTagsInput) (*request.Request, *costexplorer.GetTagsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costexplorer.GetTagsOutput)
	return ret0, ret1
}

// GetTagsRequest indicates an expected call of GetTagsRequest
func (mr *MockCostExplorerAPIMockRecorder) GetTagsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsRequest", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetTagsRequest), arg0)
}

// GetTagsWithContext mocks base method
func (m *MockCostExplorerAPI) GetTagsWithContext(arg0 context.Context, arg1 *costexplorer.GetTagsInput, arg2 ...request.Option) (*costexplorer.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagsWithContext", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsWithContext indicates an expected call of GetTagsWithContext
func (mr *MockCostExplorerAPIMockRecorder) GetTagsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsWithContext", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetTagsWithContext), varargs...)
}

// GetUsageForecast mocks base method
func (m *MockCostExplorerAPI) GetUsageForecast(arg0 *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsageForecast", arg0)
	ret0, _ := ret[0].(*costexplorer.GetUsageForecastOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsageForecast indicates an expected call of GetUsageForecast
func (mr *MockCostExplorerAPIMockRecorder) GetUsageForecast(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsageForecast", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetUsageForecast), arg0)
}

// GetUsageForecastRequest mocks base method
func (m *MockCostExplorerAPI) GetUsageForecastRequest(arg0 *costexplorer.GetUsageForecastInput) (*request.Request, *costexplorer.GetUsageForecastOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsageForecastRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costexplorer.GetUsageForecastOutput)
	return ret0, ret1
}

// GetUsageForecastRequest indicates an expected call of GetUsageForecastRequest
func (mr *MockCostExplorerAPIMockRecorder) GetUsageForecastRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsageForecastRequest", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetUsageForecastRequest), arg0)
}

// GetUsageForecastWithContext mocks base method
func (m *MockCostExplorerAPI) GetUsageForecastWithContext(arg0 context.Context, arg1 *costexplorer.GetUsageForecastInput, arg2 ...request.Option) (*costexplorer.GetUsageForecastOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsageForecastWithContext", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetUsageForecastOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsageForecastWithContext indicates an expected call of GetUsageForecastWithContext
func (mr *MockCostExplorerAPIMockRecorder) GetUsageForecastWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsageForecastWithContext", reflect.TypeOf((*MockCostExplorerAPI)(nil).GetUsageForecastWithContext), varargs...)
}
