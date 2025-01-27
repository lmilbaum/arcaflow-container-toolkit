// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ce_service/ce_service.go

// Package mock_ce_service is a generated GoMock package.
package mock_ce_service

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	docker "go.arcalot.io/arcaflow-container-toolkit/internal/docker"
)

// MockContainerEngineService is a mock of ContainerEngineService interface.
type MockContainerEngineService struct {
	ctrl     *gomock.Controller
	recorder *MockContainerEngineServiceMockRecorder
}

// MockContainerEngineServiceMockRecorder is the mock recorder for MockContainerEngineService.
type MockContainerEngineServiceMockRecorder struct {
	mock *MockContainerEngineService
}

// NewMockContainerEngineService creates a new mock instance.
func NewMockContainerEngineService(ctrl *gomock.Controller) *MockContainerEngineService {
	mock := &MockContainerEngineService{ctrl: ctrl}
	mock.recorder = &MockContainerEngineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerEngineService) EXPECT() *MockContainerEngineServiceMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockContainerEngineService) Build(filepath, name string, tags []string, archetype string, build_options *docker.BuildOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", filepath, name, tags, archetype, build_options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockContainerEngineServiceMockRecorder) Build(filepath, name, tags, archetype, build_options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockContainerEngineService)(nil).Build), filepath, name, tags, archetype, build_options)
}

// Push mocks base method.
func (m *MockContainerEngineService) Push(destination, username, password, registry_address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", destination, username, password, registry_address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockContainerEngineServiceMockRecorder) Push(destination, username, password, registry_address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockContainerEngineService)(nil).Push), destination, username, password, registry_address)
}

// Tag mocks base method.
func (m *MockContainerEngineService) Tag(image_tag, destination string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag", image_tag, destination)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockContainerEngineServiceMockRecorder) Tag(image_tag, destination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockContainerEngineService)(nil).Tag), image_tag, destination)
}
