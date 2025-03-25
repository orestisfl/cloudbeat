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

package awsfetcher

import (
	context "context"

	ec2 "github.com/elastic/cloudbeat/internal/resources/providers/awslib/ec2"
	mock "github.com/stretchr/testify/mock"
)

// mockEc2InstancesProvider is an autogenerated mock type for the ec2InstancesProvider type
type mockEc2InstancesProvider struct {
	mock.Mock
}

type mockEc2InstancesProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *mockEc2InstancesProvider) EXPECT() *mockEc2InstancesProvider_Expecter {
	return &mockEc2InstancesProvider_Expecter{mock: &_m.Mock}
}

// DescribeInstances provides a mock function with given fields: ctx
func (_m *mockEc2InstancesProvider) DescribeInstances(ctx context.Context) ([]*ec2.Ec2Instance, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DescribeInstances")
	}

	var r0 []*ec2.Ec2Instance
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*ec2.Ec2Instance, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*ec2.Ec2Instance); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Ec2Instance)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockEc2InstancesProvider_DescribeInstances_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeInstances'
type mockEc2InstancesProvider_DescribeInstances_Call struct {
	*mock.Call
}

// DescribeInstances is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockEc2InstancesProvider_Expecter) DescribeInstances(ctx interface{}) *mockEc2InstancesProvider_DescribeInstances_Call {
	return &mockEc2InstancesProvider_DescribeInstances_Call{Call: _e.mock.On("DescribeInstances", ctx)}
}

func (_c *mockEc2InstancesProvider_DescribeInstances_Call) Run(run func(ctx context.Context)) *mockEc2InstancesProvider_DescribeInstances_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockEc2InstancesProvider_DescribeInstances_Call) Return(_a0 []*ec2.Ec2Instance, _a1 error) *mockEc2InstancesProvider_DescribeInstances_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockEc2InstancesProvider_DescribeInstances_Call) RunAndReturn(run func(context.Context) ([]*ec2.Ec2Instance, error)) *mockEc2InstancesProvider_DescribeInstances_Call {
	_c.Call.Return(run)
	return _c
}

// newMockEc2InstancesProvider creates a new instance of mockEc2InstancesProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockEc2InstancesProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockEc2InstancesProvider {
	mock := &mockEc2InstancesProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
