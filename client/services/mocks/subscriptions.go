// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: SubscriptionsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	subscription "github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	gomock "github.com/golang/mock/gomock"
)

// MockSubscriptionsClient is a mock of SubscriptionsClient interface.
type MockSubscriptionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionsClientMockRecorder
}

// MockSubscriptionsClientMockRecorder is the mock recorder for MockSubscriptionsClient.
type MockSubscriptionsClientMockRecorder struct {
	mock *MockSubscriptionsClient
}

// NewMockSubscriptionsClient creates a new mock instance.
func NewMockSubscriptionsClient(ctrl *gomock.Controller) *MockSubscriptionsClient {
	mock := &MockSubscriptionsClient{ctrl: ctrl}
	mock.recorder = &MockSubscriptionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionsClient) EXPECT() *MockSubscriptionsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSubscriptionsClient) Get(arg0 context.Context, arg1 string) (subscription.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(subscription.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSubscriptionsClientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubscriptionsClient)(nil).Get), arg0, arg1)
}

// ListLocations mocks base method.
func (m *MockSubscriptionsClient) ListLocations(arg0 context.Context, arg1 string) (subscription.LocationListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocations", arg0, arg1)
	ret0, _ := ret[0].(subscription.LocationListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocations indicates an expected call of ListLocations.
func (mr *MockSubscriptionsClientMockRecorder) ListLocations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocations", reflect.TypeOf((*MockSubscriptionsClient)(nil).ListLocations), arg0, arg1)
}
