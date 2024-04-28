// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	reflect "reflect"

	apperrors "github.com/Angstreminus/exchanger/internal/apperrors"
	dto "github.com/Angstreminus/exchanger/internal/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockExchangerService is a mock of ExchangerService interface.
type MockExchangerService struct {
	ctrl     *gomock.Controller
	recorder *MockExchangerServiceMockRecorder
}

// MockExchangerServiceMockRecorder is the mock recorder for MockExchangerService.
type MockExchangerServiceMockRecorder struct {
	mock *MockExchangerService
}

// NewMockExchangerService creates a new mock instance.
func NewMockExchangerService(ctrl *gomock.Controller) *MockExchangerService {
	mock := &MockExchangerService{ctrl: ctrl}
	mock.recorder = &MockExchangerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangerService) EXPECT() *MockExchangerServiceMockRecorder {
	return m.recorder
}

// CreateExchange mocks base method.
func (m *MockExchangerService) CreateExchange(req *dto.Request) ([][]int, apperrors.Apperror) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExchange", req)
	ret0, _ := ret[0].([][]int)
	ret1, _ := ret[1].(apperrors.Apperror)
	return ret0, ret1
}

// CreateExchange indicates an expected call of CreateExchange.
func (mr *MockExchangerServiceMockRecorder) CreateExchange(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExchange", reflect.TypeOf((*MockExchangerService)(nil).CreateExchange), req)
}