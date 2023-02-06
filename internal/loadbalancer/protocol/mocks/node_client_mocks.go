// /*
// Copyright ApeCloud, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apecloud/kubeblocks/internal/loadbalancer/protocol (interfaces: NodeClient)

// Package mock_protocol is a generated GoMock package.
package mock_protocol

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"

	protocol "github.com/apecloud/kubeblocks/internal/loadbalancer/protocol"
)

// MockNodeClient is a mock of NodeClient interface.
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient.
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance.
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// CleanNetworkForENI mocks base method.
func (m *MockNodeClient) CleanNetworkForENI(arg0 context.Context, arg1 *protocol.CleanNetworkForENIRequest, arg2 ...grpc.CallOption) (*protocol.CleanNetworkForENIResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CleanNetworkForENI", varargs...)
	ret0, _ := ret[0].(*protocol.CleanNetworkForENIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CleanNetworkForENI indicates an expected call of CleanNetworkForENI.
func (mr *MockNodeClientMockRecorder) CleanNetworkForENI(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanNetworkForENI", reflect.TypeOf((*MockNodeClient)(nil).CleanNetworkForENI), varargs...)
}

// CleanNetworkForService mocks base method.
func (m *MockNodeClient) CleanNetworkForService(arg0 context.Context, arg1 *protocol.CleanNetworkForServiceRequest, arg2 ...grpc.CallOption) (*protocol.CleanNetworkForServiceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CleanNetworkForService", varargs...)
	ret0, _ := ret[0].(*protocol.CleanNetworkForServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CleanNetworkForService indicates an expected call of CleanNetworkForService.
func (mr *MockNodeClientMockRecorder) CleanNetworkForService(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanNetworkForService", reflect.TypeOf((*MockNodeClient)(nil).CleanNetworkForService), varargs...)
}

// DescribeAllENIs mocks base method.
func (m *MockNodeClient) DescribeAllENIs(arg0 context.Context, arg1 *protocol.DescribeAllENIsRequest, arg2 ...grpc.CallOption) (*protocol.DescribeAllENIsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAllENIs", varargs...)
	ret0, _ := ret[0].(*protocol.DescribeAllENIsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAllENIs indicates an expected call of DescribeAllENIs.
func (mr *MockNodeClientMockRecorder) DescribeAllENIs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAllENIs", reflect.TypeOf((*MockNodeClient)(nil).DescribeAllENIs), varargs...)
}

// DescribeNodeInfo mocks base method.
func (m *MockNodeClient) DescribeNodeInfo(arg0 context.Context, arg1 *protocol.DescribeNodeInfoRequest, arg2 ...grpc.CallOption) (*protocol.DescribeNodeInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNodeInfo", varargs...)
	ret0, _ := ret[0].(*protocol.DescribeNodeInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNodeInfo indicates an expected call of DescribeNodeInfo.
func (mr *MockNodeClientMockRecorder) DescribeNodeInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNodeInfo", reflect.TypeOf((*MockNodeClient)(nil).DescribeNodeInfo), varargs...)
}

// SetupNetworkForENI mocks base method.
func (m *MockNodeClient) SetupNetworkForENI(arg0 context.Context, arg1 *protocol.SetupNetworkForENIRequest, arg2 ...grpc.CallOption) (*protocol.SetupNetworkForENIResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetupNetworkForENI", varargs...)
	ret0, _ := ret[0].(*protocol.SetupNetworkForENIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupNetworkForENI indicates an expected call of SetupNetworkForENI.
func (mr *MockNodeClientMockRecorder) SetupNetworkForENI(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupNetworkForENI", reflect.TypeOf((*MockNodeClient)(nil).SetupNetworkForENI), varargs...)
}

// SetupNetworkForService mocks base method.
func (m *MockNodeClient) SetupNetworkForService(arg0 context.Context, arg1 *protocol.SetupNetworkForServiceRequest, arg2 ...grpc.CallOption) (*protocol.SetupNetworkForServiceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetupNetworkForService", varargs...)
	ret0, _ := ret[0].(*protocol.SetupNetworkForServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupNetworkForService indicates an expected call of SetupNetworkForService.
func (mr *MockNodeClientMockRecorder) SetupNetworkForService(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupNetworkForService", reflect.TypeOf((*MockNodeClient)(nil).SetupNetworkForService), varargs...)
}

// WaitForENIAttached mocks base method.
func (m *MockNodeClient) WaitForENIAttached(arg0 context.Context, arg1 *protocol.WaitForENIAttachedRequest, arg2 ...grpc.CallOption) (*protocol.WaitForENIAttachedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForENIAttached", varargs...)
	ret0, _ := ret[0].(*protocol.WaitForENIAttachedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForENIAttached indicates an expected call of WaitForENIAttached.
func (mr *MockNodeClientMockRecorder) WaitForENIAttached(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForENIAttached", reflect.TypeOf((*MockNodeClient)(nil).WaitForENIAttached), varargs...)
}
