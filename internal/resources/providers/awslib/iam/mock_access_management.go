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

package iam

import (
	context "context"

	awslib "github.com/elastic/cloudbeat/internal/resources/providers/awslib"

	mock "github.com/stretchr/testify/mock"
)

// MockAccessManagement is an autogenerated mock type for the AccessManagement type
type MockAccessManagement struct {
	mock.Mock
}

type MockAccessManagement_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccessManagement) EXPECT() *MockAccessManagement_Expecter {
	return &MockAccessManagement_Expecter{mock: &_m.Mock}
}

// GetAccessAnalyzers provides a mock function with given fields: ctx
func (_m *MockAccessManagement) GetAccessAnalyzers(ctx context.Context) (awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAccessAnalyzers")
	}

	var r0 awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_GetAccessAnalyzers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccessAnalyzers'
type MockAccessManagement_GetAccessAnalyzers_Call struct {
	*mock.Call
}

// GetAccessAnalyzers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccessManagement_Expecter) GetAccessAnalyzers(ctx interface{}) *MockAccessManagement_GetAccessAnalyzers_Call {
	return &MockAccessManagement_GetAccessAnalyzers_Call{Call: _e.mock.On("GetAccessAnalyzers", ctx)}
}

func (_c *MockAccessManagement_GetAccessAnalyzers_Call) Run(run func(ctx context.Context)) *MockAccessManagement_GetAccessAnalyzers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccessManagement_GetAccessAnalyzers_Call) Return(_a0 awslib.AwsResource, _a1 error) *MockAccessManagement_GetAccessAnalyzers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccessManagement_GetAccessAnalyzers_Call) RunAndReturn(run func(context.Context) (awslib.AwsResource, error)) *MockAccessManagement_GetAccessAnalyzers_Call {
	_c.Call.Return(run)
	return _c
}

// GetIAMRolePermissions provides a mock function with given fields: ctx, roleName
func (_m *MockAccessManagement) GetIAMRolePermissions(ctx context.Context, roleName string) ([]RolePolicyInfo, error) {
	ret := _m.Called(ctx, roleName)

	if len(ret) == 0 {
		panic("no return value specified for GetIAMRolePermissions")
	}

	var r0 []RolePolicyInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]RolePolicyInfo, error)); ok {
		return rf(ctx, roleName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []RolePolicyInfo); ok {
		r0 = rf(ctx, roleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RolePolicyInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, roleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_GetIAMRolePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIAMRolePermissions'
type MockAccessManagement_GetIAMRolePermissions_Call struct {
	*mock.Call
}

// GetIAMRolePermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - roleName string
func (_e *MockAccessManagement_Expecter) GetIAMRolePermissions(ctx interface{}, roleName interface{}) *MockAccessManagement_GetIAMRolePermissions_Call {
	return &MockAccessManagement_GetIAMRolePermissions_Call{Call: _e.mock.On("GetIAMRolePermissions", ctx, roleName)}
}

func (_c *MockAccessManagement_GetIAMRolePermissions_Call) Run(run func(ctx context.Context, roleName string)) *MockAccessManagement_GetIAMRolePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccessManagement_GetIAMRolePermissions_Call) Return(_a0 []RolePolicyInfo, _a1 error) *MockAccessManagement_GetIAMRolePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccessManagement_GetIAMRolePermissions_Call) RunAndReturn(run func(context.Context, string) ([]RolePolicyInfo, error)) *MockAccessManagement_GetIAMRolePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// GetPasswordPolicy provides a mock function with given fields: ctx
func (_m *MockAccessManagement) GetPasswordPolicy(ctx context.Context) (awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPasswordPolicy")
	}

	var r0 awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_GetPasswordPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPasswordPolicy'
type MockAccessManagement_GetPasswordPolicy_Call struct {
	*mock.Call
}

// GetPasswordPolicy is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccessManagement_Expecter) GetPasswordPolicy(ctx interface{}) *MockAccessManagement_GetPasswordPolicy_Call {
	return &MockAccessManagement_GetPasswordPolicy_Call{Call: _e.mock.On("GetPasswordPolicy", ctx)}
}

func (_c *MockAccessManagement_GetPasswordPolicy_Call) Run(run func(ctx context.Context)) *MockAccessManagement_GetPasswordPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccessManagement_GetPasswordPolicy_Call) Return(_a0 awslib.AwsResource, _a1 error) *MockAccessManagement_GetPasswordPolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccessManagement_GetPasswordPolicy_Call) RunAndReturn(run func(context.Context) (awslib.AwsResource, error)) *MockAccessManagement_GetPasswordPolicy_Call {
	_c.Call.Return(run)
	return _c
}

// GetPolicies provides a mock function with given fields: ctx
func (_m *MockAccessManagement) GetPolicies(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPolicies")
	}

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_GetPolicies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPolicies'
type MockAccessManagement_GetPolicies_Call struct {
	*mock.Call
}

// GetPolicies is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccessManagement_Expecter) GetPolicies(ctx interface{}) *MockAccessManagement_GetPolicies_Call {
	return &MockAccessManagement_GetPolicies_Call{Call: _e.mock.On("GetPolicies", ctx)}
}

func (_c *MockAccessManagement_GetPolicies_Call) Run(run func(ctx context.Context)) *MockAccessManagement_GetPolicies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccessManagement_GetPolicies_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockAccessManagement_GetPolicies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccessManagement_GetPolicies_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *MockAccessManagement_GetPolicies_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsers provides a mock function with given fields: ctx
func (_m *MockAccessManagement) GetUsers(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_GetUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsers'
type MockAccessManagement_GetUsers_Call struct {
	*mock.Call
}

// GetUsers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccessManagement_Expecter) GetUsers(ctx interface{}) *MockAccessManagement_GetUsers_Call {
	return &MockAccessManagement_GetUsers_Call{Call: _e.mock.On("GetUsers", ctx)}
}

func (_c *MockAccessManagement_GetUsers_Call) Run(run func(ctx context.Context)) *MockAccessManagement_GetUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccessManagement_GetUsers_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockAccessManagement_GetUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccessManagement_GetUsers_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *MockAccessManagement_GetUsers_Call {
	_c.Call.Return(run)
	return _c
}

// ListServerCertificates provides a mock function with given fields: ctx
func (_m *MockAccessManagement) ListServerCertificates(ctx context.Context) (awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListServerCertificates")
	}

	var r0 awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_ListServerCertificates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListServerCertificates'
type MockAccessManagement_ListServerCertificates_Call struct {
	*mock.Call
}

// ListServerCertificates is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccessManagement_Expecter) ListServerCertificates(ctx interface{}) *MockAccessManagement_ListServerCertificates_Call {
	return &MockAccessManagement_ListServerCertificates_Call{Call: _e.mock.On("ListServerCertificates", ctx)}
}

func (_c *MockAccessManagement_ListServerCertificates_Call) Run(run func(ctx context.Context)) *MockAccessManagement_ListServerCertificates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccessManagement_ListServerCertificates_Call) Return(_a0 awslib.AwsResource, _a1 error) *MockAccessManagement_ListServerCertificates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccessManagement_ListServerCertificates_Call) RunAndReturn(run func(context.Context) (awslib.AwsResource, error)) *MockAccessManagement_ListServerCertificates_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccessManagement creates a new instance of MockAccessManagement. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccessManagement(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccessManagement {
	mock := &MockAccessManagement{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
