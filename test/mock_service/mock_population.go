// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/population.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	gomock "github.com/golang/mock/gomock"
)

// MockPopulationService is a mock of PopulationService interface.
type MockPopulationService struct {
	ctrl     *gomock.Controller
	recorder *MockPopulationServiceMockRecorder
}

// MockPopulationServiceMockRecorder is the mock recorder for MockPopulationService.
type MockPopulationServiceMockRecorder struct {
	mock *MockPopulationService
}

// NewMockPopulationService creates a new mock instance.
func NewMockPopulationService(ctrl *gomock.Controller) *MockPopulationService {
	mock := &MockPopulationService{ctrl: ctrl}
	mock.recorder = &MockPopulationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPopulationService) EXPECT() *MockPopulationServiceMockRecorder {
	return m.recorder
}

// GetPopulationInformation mocks base method.
func (m *MockPopulationService) GetPopulationInformation(citizenID string) (model.Population, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopulationInformation", citizenID)
	ret0, _ := ret[0].(model.Population)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopulationInformation indicates an expected call of GetPopulationInformation.
func (mr *MockPopulationServiceMockRecorder) GetPopulationInformation(citizenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopulationInformation", reflect.TypeOf((*MockPopulationService)(nil).GetPopulationInformation), citizenID)
}
