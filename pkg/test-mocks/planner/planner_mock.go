// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicatedhq/ship/pkg/lifecycle/render/planner (interfaces: Planner)

// Package planner is a generated GoMock package.
package planner

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	libyaml "github.com/replicatedhq/libyaml"
	api "github.com/replicatedhq/ship/pkg/api"
	config "github.com/replicatedhq/ship/pkg/lifecycle/render/config"
	planner "github.com/replicatedhq/ship/pkg/lifecycle/render/planner"
)

// MockPlanner is a mock of Planner interface
type MockPlanner struct {
	ctrl     *gomock.Controller
	recorder *MockPlannerMockRecorder
}

// MockPlannerMockRecorder is the mock recorder for MockPlanner
type MockPlannerMockRecorder struct {
	mock *MockPlanner
}

// NewMockPlanner creates a new mock instance
func NewMockPlanner(ctrl *gomock.Controller) *MockPlanner {
	mock := &MockPlanner{ctrl: ctrl}
	mock.recorder = &MockPlannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlanner) EXPECT() *MockPlannerMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockPlanner) Build(arg0 []api.Asset, arg1 []libyaml.ConfigGroup, arg2 api.ReleaseMetadata, arg3 map[string]interface{}) planner.Plan {
	ret := m.ctrl.Call(m, "Build", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(planner.Plan)
	return ret0
}

// Build indicates an expected call of Build
func (mr *MockPlannerMockRecorder) Build(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockPlanner)(nil).Build), arg0, arg1, arg2, arg3)
}

// Confirm mocks base method
func (m *MockPlanner) Confirm(arg0 planner.Plan) (bool, error) {
	ret := m.ctrl.Call(m, "Confirm", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Confirm indicates an expected call of Confirm
func (mr *MockPlannerMockRecorder) Confirm(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Confirm", reflect.TypeOf((*MockPlanner)(nil).Confirm), arg0)
}

// Execute mocks base method
func (m *MockPlanner) Execute(arg0 context.Context, arg1 planner.Plan) error {
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockPlannerMockRecorder) Execute(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockPlanner)(nil).Execute), arg0, arg1)
}

// WithDaemon mocks base method
func (m *MockPlanner) WithDaemon(arg0 config.Daemon) planner.Planner {
	ret := m.ctrl.Call(m, "WithDaemon", arg0)
	ret0, _ := ret[0].(planner.Planner)
	return ret0
}

// WithDaemon indicates an expected call of WithDaemon
func (mr *MockPlannerMockRecorder) WithDaemon(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDaemon", reflect.TypeOf((*MockPlanner)(nil).WithDaemon), arg0)
}
