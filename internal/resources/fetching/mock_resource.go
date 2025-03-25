// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.53.3. DO NOT EDIT.

package fetching

import mock "github.com/stretchr/testify/mock"

// MockResource is an autogenerated mock type for the Resource type
type MockResource struct {
	mock.Mock
}

type MockResource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockResource) EXPECT() *MockResource_Expecter {
	return &MockResource_Expecter{mock: &_m.Mock}
}

// GetData provides a mock function with no fields
func (_m *MockResource) GetData() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockResource_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type MockResource_GetData_Call struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
func (_e *MockResource_Expecter) GetData() *MockResource_GetData_Call {
	return &MockResource_GetData_Call{Call: _e.mock.On("GetData")}
}

func (_c *MockResource_GetData_Call) Run(run func()) *MockResource_GetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResource_GetData_Call) Return(_a0 interface{}) *MockResource_GetData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockResource_GetData_Call) RunAndReturn(run func() interface{}) *MockResource_GetData_Call {
	_c.Call.Return(run)
	return _c
}

// GetElasticCommonData provides a mock function with no fields
func (_m *MockResource) GetElasticCommonData() (map[string]interface{}, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetElasticCommonData")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]interface{}, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResource_GetElasticCommonData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetElasticCommonData'
type MockResource_GetElasticCommonData_Call struct {
	*mock.Call
}

// GetElasticCommonData is a helper method to define mock.On call
func (_e *MockResource_Expecter) GetElasticCommonData() *MockResource_GetElasticCommonData_Call {
	return &MockResource_GetElasticCommonData_Call{Call: _e.mock.On("GetElasticCommonData")}
}

func (_c *MockResource_GetElasticCommonData_Call) Run(run func()) *MockResource_GetElasticCommonData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResource_GetElasticCommonData_Call) Return(_a0 map[string]interface{}, _a1 error) *MockResource_GetElasticCommonData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResource_GetElasticCommonData_Call) RunAndReturn(run func() (map[string]interface{}, error)) *MockResource_GetElasticCommonData_Call {
	_c.Call.Return(run)
	return _c
}

// GetIds provides a mock function with no fields
func (_m *MockResource) GetIds() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetIds")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockResource_GetIds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIds'
type MockResource_GetIds_Call struct {
	*mock.Call
}

// GetIds is a helper method to define mock.On call
func (_e *MockResource_Expecter) GetIds() *MockResource_GetIds_Call {
	return &MockResource_GetIds_Call{Call: _e.mock.On("GetIds")}
}

func (_c *MockResource_GetIds_Call) Run(run func()) *MockResource_GetIds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResource_GetIds_Call) Return(_a0 []string) *MockResource_GetIds_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockResource_GetIds_Call) RunAndReturn(run func() []string) *MockResource_GetIds_Call {
	_c.Call.Return(run)
	return _c
}

// GetMetadata provides a mock function with no fields
func (_m *MockResource) GetMetadata() (ResourceMetadata, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 ResourceMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func() (ResourceMetadata, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ResourceMetadata); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ResourceMetadata)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResource_GetMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetadata'
type MockResource_GetMetadata_Call struct {
	*mock.Call
}

// GetMetadata is a helper method to define mock.On call
func (_e *MockResource_Expecter) GetMetadata() *MockResource_GetMetadata_Call {
	return &MockResource_GetMetadata_Call{Call: _e.mock.On("GetMetadata")}
}

func (_c *MockResource_GetMetadata_Call) Run(run func()) *MockResource_GetMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResource_GetMetadata_Call) Return(_a0 ResourceMetadata, _a1 error) *MockResource_GetMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResource_GetMetadata_Call) RunAndReturn(run func() (ResourceMetadata, error)) *MockResource_GetMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockResource creates a new instance of MockResource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockResource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockResource {
	mock := &MockResource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
