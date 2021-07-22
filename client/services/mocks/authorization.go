// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: RoleAssignmentsClient,RoleDefinitionsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	authorization "github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	gomock "github.com/golang/mock/gomock"
)

// MockRoleAssignmentsClient is a mock of RoleAssignmentsClient interface.
type MockRoleAssignmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoleAssignmentsClientMockRecorder
}

// MockRoleAssignmentsClientMockRecorder is the mock recorder for MockRoleAssignmentsClient.
type MockRoleAssignmentsClientMockRecorder struct {
	mock *MockRoleAssignmentsClient
}

// NewMockRoleAssignmentsClient creates a new mock instance.
func NewMockRoleAssignmentsClient(ctrl *gomock.Controller) *MockRoleAssignmentsClient {
	mock := &MockRoleAssignmentsClient{ctrl: ctrl}
	mock.recorder = &MockRoleAssignmentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleAssignmentsClient) EXPECT() *MockRoleAssignmentsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockRoleAssignmentsClient) List(arg0 context.Context, arg1 string) (authorization.RoleAssignmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(authorization.RoleAssignmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoleAssignmentsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleAssignmentsClient)(nil).List), arg0, arg1)
}

// MockRoleDefinitionsClient is a mock of RoleDefinitionsClient interface.
type MockRoleDefinitionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoleDefinitionsClientMockRecorder
}

// MockRoleDefinitionsClientMockRecorder is the mock recorder for MockRoleDefinitionsClient.
type MockRoleDefinitionsClientMockRecorder struct {
	mock *MockRoleDefinitionsClient
}

// NewMockRoleDefinitionsClient creates a new mock instance.
func NewMockRoleDefinitionsClient(ctrl *gomock.Controller) *MockRoleDefinitionsClient {
	mock := &MockRoleDefinitionsClient{ctrl: ctrl}
	mock.recorder = &MockRoleDefinitionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleDefinitionsClient) EXPECT() *MockRoleDefinitionsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockRoleDefinitionsClient) List(arg0 context.Context, arg1, arg2 string) (authorization.RoleDefinitionListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(authorization.RoleDefinitionListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoleDefinitionsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleDefinitionsClient)(nil).List), arg0, arg1, arg2)
}