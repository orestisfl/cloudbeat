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

package configservice

import (
	context "context"

	awslib "github.com/elastic/cloudbeat/internal/resources/providers/awslib"

	mock "github.com/stretchr/testify/mock"
)

// MockConfigService is an autogenerated mock type for the ConfigService type
type MockConfigService struct {
	mock.Mock
}

type MockConfigService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConfigService) EXPECT() *MockConfigService_Expecter {
	return &MockConfigService_Expecter{mock: &_m.Mock}
}

// DescribeConfigRecorders provides a mock function with given fields: ctx
func (_m *MockConfigService) DescribeConfigRecorders(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DescribeConfigRecorders")
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

// MockConfigService_DescribeConfigRecorders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeConfigRecorders'
type MockConfigService_DescribeConfigRecorders_Call struct {
	*mock.Call
}

// DescribeConfigRecorders is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockConfigService_Expecter) DescribeConfigRecorders(ctx interface{}) *MockConfigService_DescribeConfigRecorders_Call {
	return &MockConfigService_DescribeConfigRecorders_Call{Call: _e.mock.On("DescribeConfigRecorders", ctx)}
}

func (_c *MockConfigService_DescribeConfigRecorders_Call) Run(run func(ctx context.Context)) *MockConfigService_DescribeConfigRecorders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockConfigService_DescribeConfigRecorders_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockConfigService_DescribeConfigRecorders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConfigService_DescribeConfigRecorders_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *MockConfigService_DescribeConfigRecorders_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConfigService creates a new instance of MockConfigService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConfigService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConfigService {
	mock := &MockConfigService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
