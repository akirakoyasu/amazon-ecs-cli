// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/aws/clients/ecr (interfaces: Client)

package mock_ecr

import (
	ecr "github.com/aws/amazon-ecs-cli/ecs-cli/modules/aws/clients/ecr"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) CreateRepository(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "CreateRepository", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) CreateRepository(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRepository", arg0)
}

func (_m *MockClient) GetAuthorizationToken(_param0 string) (*ecr.Auth, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorizationToken", _param0)
	ret0, _ := ret[0].(*ecr.Auth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetAuthorizationToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizationToken", arg0)
}

func (_m *MockClient) GetAuthorizationTokenByID(_param0 string) (*ecr.Auth, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorizationTokenByID", _param0)
	ret0, _ := ret[0].(*ecr.Auth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) GetAuthorizationTokenByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizationTokenByID", arg0)
}

func (_m *MockClient) GetImages(_param0 []*string, _param1 string, _param2 string, _param3 ecr.ProcessImageDetails) error {
	ret := _m.ctrl.Call(_m, "GetImages", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) GetImages(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetImages", arg0, arg1, arg2, arg3)
}

func (_m *MockClient) RepositoryExists(_param0 string) bool {
	ret := _m.ctrl.Call(_m, "RepositoryExists", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockClientRecorder) RepositoryExists(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RepositoryExists", arg0)
}
