// Code generated by mockery v2.36.0. DO NOT EDIT.

package product

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/teq-quocbang/store/model"

	uuid "github.com/google/uuid"
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
func (_m *MockRepository) Create(_a0 context.Context, _a1 *model.Product) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Product) error); ok {
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
//   - _a1 *model.Product
func (_e *MockRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockRepository_Create_Call) Run(run func(_a0 context.Context, _a1 *model.Product)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Product))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, *model.Product) error) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateList provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) CreateList(_a0 context.Context, _a1 []model.Product) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []model.Product) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_CreateList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateList'
type MockRepository_CreateList_Call struct {
	*mock.Call
}

// CreateList is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []model.Product
func (_e *MockRepository_Expecter) CreateList(_a0 interface{}, _a1 interface{}) *MockRepository_CreateList_Call {
	return &MockRepository_CreateList_Call{Call: _e.mock.On("CreateList", _a0, _a1)}
}

func (_c *MockRepository_CreateList_Call) Run(run func(_a0 context.Context, _a1 []model.Product)) *MockRepository_CreateList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]model.Product))
	})
	return _c
}

func (_c *MockRepository_CreateList_Call) Return(_a0 error) *MockRepository_CreateList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_CreateList_Call) RunAndReturn(run func(context.Context, []model.Product) error) *MockRepository_CreateList_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Delete(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uuid.UUID
func (_e *MockRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}) *MockRepository_Delete_Call {
	return &MockRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *MockRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 uuid.UUID)) *MockRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockRepository_Delete_Call) Return(_a0 error) *MockRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Delete_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *MockRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (model.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) model.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockRepository_Expecter) GetByID(ctx interface{}, id interface{}) *MockRepository_GetByID_Call {
	return &MockRepository_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *MockRepository_GetByID_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockRepository_GetByID_Call) Return(_a0 model.Product, _a1 error) *MockRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (model.Product, error)) *MockRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: _a0
func (_m *MockRepository) GetList(_a0 context.Context) ([]model.Product, error) {
	ret := _m.Called(_a0)

	var r0 []model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockRepository_Expecter) GetList(_a0 interface{}) *MockRepository_GetList_Call {
	return &MockRepository_GetList_Call{Call: _e.mock.On("GetList", _a0)}
}

func (_c *MockRepository_GetList_Call) Run(run func(_a0 context.Context)) *MockRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepository_GetList_Call) Return(_a0 []model.Product, _a1 error) *MockRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetList_Call) RunAndReturn(run func(context.Context) ([]model.Product, error)) *MockRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Update(_a0 context.Context, _a1 *model.Product) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Product) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Product
func (_e *MockRepository_Expecter) Update(_a0 interface{}, _a1 interface{}) *MockRepository_Update_Call {
	return &MockRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *MockRepository_Update_Call) Run(run func(_a0 context.Context, _a1 *model.Product)) *MockRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Product))
	})
	return _c
}

func (_c *MockRepository_Update_Call) Return(_a0 error) *MockRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Update_Call) RunAndReturn(run func(context.Context, *model.Product) error) *MockRepository_Update_Call {
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
