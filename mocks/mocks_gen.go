// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/drone/autoscaler (interfaces: Engine,ServerStore,Provider)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	autoscaler "github.com/drone/autoscaler"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Pause mocks base method
func (m *MockEngine) Pause() {
	m.ctrl.Call(m, "Pause")
}

// Pause indicates an expected call of Pause
func (mr *MockEngineMockRecorder) Pause() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockEngine)(nil).Pause))
}

// Paused mocks base method
func (m *MockEngine) Paused() bool {
	ret := m.ctrl.Call(m, "Paused")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Paused indicates an expected call of Paused
func (mr *MockEngineMockRecorder) Paused() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paused", reflect.TypeOf((*MockEngine)(nil).Paused))
}

// Resume mocks base method
func (m *MockEngine) Resume() {
	m.ctrl.Call(m, "Resume")
}

// Resume indicates an expected call of Resume
func (mr *MockEngineMockRecorder) Resume() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resume", reflect.TypeOf((*MockEngine)(nil).Resume))
}

// Start mocks base method
func (m *MockEngine) Start(arg0 context.Context) {
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start
func (mr *MockEngineMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEngine)(nil).Start), arg0)
}

// MockServerStore is a mock of ServerStore interface
type MockServerStore struct {
	ctrl     *gomock.Controller
	recorder *MockServerStoreMockRecorder
}

// MockServerStoreMockRecorder is the mock recorder for MockServerStore
type MockServerStoreMockRecorder struct {
	mock *MockServerStore
}

// NewMockServerStore creates a new mock instance
func NewMockServerStore(ctrl *gomock.Controller) *MockServerStore {
	mock := &MockServerStore{ctrl: ctrl}
	mock.recorder = &MockServerStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerStore) EXPECT() *MockServerStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServerStore) Create(arg0 context.Context, arg1 *autoscaler.Server) error {
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockServerStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServerStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockServerStore) Delete(arg0 context.Context, arg1 *autoscaler.Server) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockServerStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServerStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockServerStore) Find(arg0 context.Context, arg1 string) (*autoscaler.Server, error) {
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*autoscaler.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockServerStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockServerStore)(nil).Find), arg0, arg1)
}

// List mocks base method
func (m *MockServerStore) List(arg0 context.Context) ([]*autoscaler.Server, error) {
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*autoscaler.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockServerStoreMockRecorder) List(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServerStore)(nil).List), arg0)
}

// ListState mocks base method
func (m *MockServerStore) ListState(arg0 context.Context, arg1 autoscaler.ServerState) ([]*autoscaler.Server, error) {
	ret := m.ctrl.Call(m, "ListState", arg0, arg1)
	ret0, _ := ret[0].([]*autoscaler.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListState indicates an expected call of ListState
func (mr *MockServerStoreMockRecorder) ListState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListState", reflect.TypeOf((*MockServerStore)(nil).ListState), arg0, arg1)
}

// Update mocks base method
func (m *MockServerStore) Update(arg0 context.Context, arg1 *autoscaler.Server) error {
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockServerStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServerStore)(nil).Update), arg0, arg1)
}

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockProvider) Create(arg0 context.Context, arg1 autoscaler.InstanceCreateOpts) (*autoscaler.Instance, error) {
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*autoscaler.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockProviderMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProvider)(nil).Create), arg0, arg1)
}

// Destroy mocks base method
func (m *MockProvider) Destroy(arg0 context.Context, arg1 *autoscaler.Instance) error {
	ret := m.ctrl.Call(m, "Destroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockProviderMockRecorder) Destroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockProvider)(nil).Destroy), arg0, arg1)
}
