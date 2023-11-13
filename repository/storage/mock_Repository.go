// Code generated by mockery v2.36.0. DO NOT EDIT.

package storage

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/teq-quocbang/store/model"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// GetListStorageByLocat provides a mock function with given fields: ctx, locat
func (_m *MockRepository) GetListStorageByLocat(ctx context.Context, locat string) ([]model.Storage, error) {
	ret := _m.Called(ctx, locat)

	var r0 []model.Storage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]model.Storage, error)); ok {
		return rf(ctx, locat)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Storage); ok {
		r0 = rf(ctx, locat)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Storage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, locat)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetListStorageByLocat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListStorageByLocat'
type MockRepository_GetListStorageByLocat_Call struct {
	*mock.Call
}

// GetListStorageByLocat is a helper method to define mock.On call
//   - ctx context.Context
//   - locat string
func (_e *MockRepository_Expecter) GetListStorageByLocat(ctx interface{}, locat interface{}) *MockRepository_GetListStorageByLocat_Call {
	return &MockRepository_GetListStorageByLocat_Call{Call: _e.mock.On("GetListStorageByLocat", ctx, locat)}
}

func (_c *MockRepository_GetListStorageByLocat_Call) Run(run func(ctx context.Context, locat string)) *MockRepository_GetListStorageByLocat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetListStorageByLocat_Call) Return(_a0 []model.Storage, _a1 error) *MockRepository_GetListStorageByLocat_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetListStorageByLocat_Call) RunAndReturn(run func(context.Context, string) ([]model.Storage, error)) *MockRepository_GetListStorageByLocat_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertStorage provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) UpsertStorage(_a0 context.Context, _a1 *model.Storage) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Storage) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_UpsertStorage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertStorage'
type MockRepository_UpsertStorage_Call struct {
	*mock.Call
}

// UpsertStorage is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Storage
func (_e *MockRepository_Expecter) UpsertStorage(_a0 interface{}, _a1 interface{}) *MockRepository_UpsertStorage_Call {
	return &MockRepository_UpsertStorage_Call{Call: _e.mock.On("UpsertStorage", _a0, _a1)}
}

func (_c *MockRepository_UpsertStorage_Call) Run(run func(_a0 context.Context, _a1 *model.Storage)) *MockRepository_UpsertStorage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Storage))
	})
	return _c
}

func (_c *MockRepository_UpsertStorage_Call) Return(_a0 error) *MockRepository_UpsertStorage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_UpsertStorage_Call) RunAndReturn(run func(context.Context, *model.Storage) error) *MockRepository_UpsertStorage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}