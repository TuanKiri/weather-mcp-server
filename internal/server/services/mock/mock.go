// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/server/services/services.go
//
// Generated by this command:
//
//	mockgen -source=./internal/server/services/services.go -destination=internal/server/services/mock/mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	services "github.com/TuanKiri/weather-mcp-server/internal/server/services"
	gomock "go.uber.org/mock/gomock"
)

// MockServices is a mock of Services interface.
type MockServices struct {
	ctrl     *gomock.Controller
	recorder *MockServicesMockRecorder
	isgomock struct{}
}

// MockServicesMockRecorder is the mock recorder for MockServices.
type MockServicesMockRecorder struct {
	mock *MockServices
}

// NewMockServices creates a new mock instance.
func NewMockServices(ctrl *gomock.Controller) *MockServices {
	mock := &MockServices{ctrl: ctrl}
	mock.recorder = &MockServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServices) EXPECT() *MockServicesMockRecorder {
	return m.recorder
}

// Weather mocks base method.
func (m *MockServices) Weather() services.WeatherService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Weather")
	ret0, _ := ret[0].(services.WeatherService)
	return ret0
}

// Weather indicates an expected call of Weather.
func (mr *MockServicesMockRecorder) Weather() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Weather", reflect.TypeOf((*MockServices)(nil).Weather))
}

// MockWeatherService is a mock of WeatherService interface.
type MockWeatherService struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherServiceMockRecorder
	isgomock struct{}
}

// MockWeatherServiceMockRecorder is the mock recorder for MockWeatherService.
type MockWeatherServiceMockRecorder struct {
	mock *MockWeatherService
}

// NewMockWeatherService creates a new mock instance.
func NewMockWeatherService(ctrl *gomock.Controller) *MockWeatherService {
	mock := &MockWeatherService{ctrl: ctrl}
	mock.recorder = &MockWeatherServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeatherService) EXPECT() *MockWeatherServiceMockRecorder {
	return m.recorder
}

// Current mocks base method.
func (m *MockWeatherService) Current(ctx context.Context, city string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current", ctx, city)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Current indicates an expected call of Current.
func (mr *MockWeatherServiceMockRecorder) Current(ctx, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockWeatherService)(nil).Current), ctx, city)
}
