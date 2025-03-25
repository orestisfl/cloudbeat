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

package inventory

import (
	context "context"

	cycle "github.com/elastic/cloudbeat/internal/resources/fetching/cycle"
	mock "github.com/stretchr/testify/mock"
)

// MockStorageAccountProviderAPI is an autogenerated mock type for the StorageAccountProviderAPI type
type MockStorageAccountProviderAPI struct {
	mock.Mock
}

type MockStorageAccountProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorageAccountProviderAPI) EXPECT() *MockStorageAccountProviderAPI_Expecter {
	return &MockStorageAccountProviderAPI_Expecter{mock: &_m.Mock}
}

// ListDiagnosticSettingsAssetTypes provides a mock function with given fields: ctx, cycleMetadata, subscriptionIDs
func (_m *MockStorageAccountProviderAPI) ListDiagnosticSettingsAssetTypes(ctx context.Context, cycleMetadata cycle.Metadata, subscriptionIDs []string) ([]AzureAsset, error) {
	ret := _m.Called(ctx, cycleMetadata, subscriptionIDs)

	if len(ret) == 0 {
		panic("no return value specified for ListDiagnosticSettingsAssetTypes")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cycle.Metadata, []string) ([]AzureAsset, error)); ok {
		return rf(ctx, cycleMetadata, subscriptionIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cycle.Metadata, []string) []AzureAsset); ok {
		r0 = rf(ctx, cycleMetadata, subscriptionIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cycle.Metadata, []string) error); ok {
		r1 = rf(ctx, cycleMetadata, subscriptionIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDiagnosticSettingsAssetTypes'
type MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call struct {
	*mock.Call
}

// ListDiagnosticSettingsAssetTypes is a helper method to define mock.On call
//   - ctx context.Context
//   - cycleMetadata cycle.Metadata
//   - subscriptionIDs []string
func (_e *MockStorageAccountProviderAPI_Expecter) ListDiagnosticSettingsAssetTypes(ctx interface{}, cycleMetadata interface{}, subscriptionIDs interface{}) *MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call {
	return &MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call{Call: _e.mock.On("ListDiagnosticSettingsAssetTypes", ctx, cycleMetadata, subscriptionIDs)}
}

func (_c *MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call) Run(run func(ctx context.Context, cycleMetadata cycle.Metadata, subscriptionIDs []string)) *MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cycle.Metadata), args[2].([]string))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call) RunAndReturn(run func(context.Context, cycle.Metadata, []string) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListDiagnosticSettingsAssetTypes_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountBlobContainers provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountBlobContainers(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountBlobContainers")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountBlobContainers'
type MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call struct {
	*mock.Call
}

// ListStorageAccountBlobContainers is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountBlobContainers(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call{Call: _e.mock.On("ListStorageAccountBlobContainers", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountBlobContainers_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountBlobServices provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountBlobServices(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountBlobServices")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountBlobServices'
type MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call struct {
	*mock.Call
}

// ListStorageAccountBlobServices is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountBlobServices(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call{Call: _e.mock.On("ListStorageAccountBlobServices", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountBlobServices_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountFileServices provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountFileServices(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountFileServices")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountFileServices'
type MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call struct {
	*mock.Call
}

// ListStorageAccountFileServices is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountFileServices(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call{Call: _e.mock.On("ListStorageAccountFileServices", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountFileServices_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountFileShares provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountFileShares(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountFileShares")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountFileShares'
type MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call struct {
	*mock.Call
}

// ListStorageAccountFileShares is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountFileShares(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call{Call: _e.mock.On("ListStorageAccountFileShares", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountFileShares_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountQueueServices provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountQueueServices(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountQueueServices")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountQueueServices'
type MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call struct {
	*mock.Call
}

// ListStorageAccountQueueServices is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountQueueServices(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call{Call: _e.mock.On("ListStorageAccountQueueServices", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountQueueServices_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountQueues provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountQueues(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountQueues")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountQueues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountQueues'
type MockStorageAccountProviderAPI_ListStorageAccountQueues_Call struct {
	*mock.Call
}

// ListStorageAccountQueues is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountQueues(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountQueues_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountQueues_Call{Call: _e.mock.On("ListStorageAccountQueues", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountQueues_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountQueues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountQueues_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountQueues_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountQueues_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountQueues_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountTableServices provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountTableServices(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountTableServices")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountTableServices'
type MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call struct {
	*mock.Call
}

// ListStorageAccountTableServices is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountTableServices(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call{Call: _e.mock.On("ListStorageAccountTableServices", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountTableServices_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountTables provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountTables(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountTables")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountTables'
type MockStorageAccountProviderAPI_ListStorageAccountTables_Call struct {
	*mock.Call
}

// ListStorageAccountTables is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountTables(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountTables_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountTables_Call{Call: _e.mock.On("ListStorageAccountTables", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountTables_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountTables_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountTables_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountTables_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccounts provides a mock function with given fields: ctx, storageAccountsSubscriptionsIds
func (_m *MockStorageAccountProviderAPI) ListStorageAccounts(ctx context.Context, storageAccountsSubscriptionsIds []string) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccountsSubscriptionsIds)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccounts")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccountsSubscriptionsIds)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []AzureAsset); ok {
		r0 = rf(ctx, storageAccountsSubscriptionsIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, storageAccountsSubscriptionsIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccounts'
type MockStorageAccountProviderAPI_ListStorageAccounts_Call struct {
	*mock.Call
}

// ListStorageAccounts is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccountsSubscriptionsIds []string
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccounts(ctx interface{}, storageAccountsSubscriptionsIds interface{}) *MockStorageAccountProviderAPI_ListStorageAccounts_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccounts_Call{Call: _e.mock.On("ListStorageAccounts", ctx, storageAccountsSubscriptionsIds)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccounts_Call) Run(run func(ctx context.Context, storageAccountsSubscriptionsIds []string)) *MockStorageAccountProviderAPI_ListStorageAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccounts_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccounts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccounts_Call) RunAndReturn(run func(context.Context, []string) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountsBlobDiagnosticSettings provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountsBlobDiagnosticSettings(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountsBlobDiagnosticSettings")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountsBlobDiagnosticSettings'
type MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call struct {
	*mock.Call
}

// ListStorageAccountsBlobDiagnosticSettings is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountsBlobDiagnosticSettings(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call{Call: _e.mock.On("ListStorageAccountsBlobDiagnosticSettings", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountsBlobDiagnosticSettings_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountsQueueDiagnosticSettings provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountsQueueDiagnosticSettings(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountsQueueDiagnosticSettings")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountsQueueDiagnosticSettings'
type MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call struct {
	*mock.Call
}

// ListStorageAccountsQueueDiagnosticSettings is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountsQueueDiagnosticSettings(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call{Call: _e.mock.On("ListStorageAccountsQueueDiagnosticSettings", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountsQueueDiagnosticSettings_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountsTableDiagnosticSettings provides a mock function with given fields: ctx, storageAccounts
func (_m *MockStorageAccountProviderAPI) ListStorageAccountsTableDiagnosticSettings(ctx context.Context, storageAccounts []AzureAsset) ([]AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	if len(ret) == 0 {
		panic("no return value specified for ListStorageAccountsTableDiagnosticSettings")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) ([]AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []AzureAsset) []AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountsTableDiagnosticSettings'
type MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call struct {
	*mock.Call
}

// ListStorageAccountsTableDiagnosticSettings is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []AzureAsset
func (_e *MockStorageAccountProviderAPI_Expecter) ListStorageAccountsTableDiagnosticSettings(ctx interface{}, storageAccounts interface{}) *MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call {
	return &MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call{Call: _e.mock.On("ListStorageAccountsTableDiagnosticSettings", ctx, storageAccounts)}
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call) Run(run func(ctx context.Context, storageAccounts []AzureAsset)) *MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]AzureAsset))
	})
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call) Return(_a0 []AzureAsset, _a1 error) *MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call) RunAndReturn(run func(context.Context, []AzureAsset) ([]AzureAsset, error)) *MockStorageAccountProviderAPI_ListStorageAccountsTableDiagnosticSettings_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorageAccountProviderAPI creates a new instance of MockStorageAccountProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorageAccountProviderAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorageAccountProviderAPI {
	mock := &MockStorageAccountProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
