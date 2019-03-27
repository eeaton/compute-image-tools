//  Copyright 2019 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/common/domain (interfaces: TarGcsExtractorInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTarGcsExtractorInterface is a mock of TarGcsExtractorInterface interface
type MockTarGcsExtractorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTarGcsExtractorInterfaceMockRecorder
}

// MockTarGcsExtractorInterfaceMockRecorder is the mock recorder for MockTarGcsExtractorInterface
type MockTarGcsExtractorInterfaceMockRecorder struct {
	mock *MockTarGcsExtractorInterface
}

// NewMockTarGcsExtractorInterface creates a new mock instance
func NewMockTarGcsExtractorInterface(ctrl *gomock.Controller) *MockTarGcsExtractorInterface {
	mock := &MockTarGcsExtractorInterface{ctrl: ctrl}
	mock.recorder = &MockTarGcsExtractorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTarGcsExtractorInterface) EXPECT() *MockTarGcsExtractorInterfaceMockRecorder {
	return m.recorder
}

// ExtractTarToGcs mocks base method
func (m *MockTarGcsExtractorInterface) ExtractTarToGcs(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractTarToGcs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExtractTarToGcs indicates an expected call of ExtractTarToGcs
func (mr *MockTarGcsExtractorInterfaceMockRecorder) ExtractTarToGcs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractTarToGcs", reflect.TypeOf((*MockTarGcsExtractorInterface)(nil).ExtractTarToGcs), arg0, arg1)
}