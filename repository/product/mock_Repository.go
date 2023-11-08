// Code generated by mockery v2.36.0. DO NOT EDIT.

package product

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

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Create(_a0 context.Context, _a1 model.Product) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Product) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.Product
func (_e *MockRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockRepository_Create_Call) Run(run func(_a0 context.Context, _a1 model.Product)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Product))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, model.Product) error) *MockRepository_Create_Call {
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
