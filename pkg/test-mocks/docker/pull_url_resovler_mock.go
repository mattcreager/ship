// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicatedhq/ship/pkg/lifecycle/render/docker (interfaces: PullURLResolver)

// Package docker is a generated GoMock package.
package docker

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/replicatedhq/ship/pkg/api"
)

// MockPullURLResolver is a mock of PullURLResolver interface
type MockPullURLResolver struct {
	ctrl     *gomock.Controller
	recorder *MockPullURLResolverMockRecorder
}

// MockPullURLResolverMockRecorder is the mock recorder for MockPullURLResolver
type MockPullURLResolverMockRecorder struct {
	mock *MockPullURLResolver
}

// NewMockPullURLResolver creates a new mock instance
func NewMockPullURLResolver(ctrl *gomock.Controller) *MockPullURLResolver {
	mock := &MockPullURLResolver{ctrl: ctrl}
	mock.recorder = &MockPullURLResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPullURLResolver) EXPECT() *MockPullURLResolverMockRecorder {
	return m.recorder
}

// ResolvePullURL mocks base method
func (m *MockPullURLResolver) ResolvePullURL(arg0 *api.DockerAsset, arg1 api.ReleaseMetadata) (string, error) {
	ret := m.ctrl.Call(m, "ResolvePullURL", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolvePullURL indicates an expected call of ResolvePullURL
func (mr *MockPullURLResolverMockRecorder) ResolvePullURL(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolvePullURL", reflect.TypeOf((*MockPullURLResolver)(nil).ResolvePullURL), arg0, arg1)
}
