// Code generated by mockery v2.36.0. DO NOT EDIT.

package checkout

import (
	context "context"

	codetype "github.com/teq-quocbang/store/codetype"

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

// CreateCustomerOrder provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) CreateCustomerOrder(_a0 context.Context, _a1 *model.CustomerOrder) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CustomerOrder) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_CreateCustomerOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomerOrder'
type MockRepository_CreateCustomerOrder_Call struct {
	*mock.Call
}

// CreateCustomerOrder is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.CustomerOrder
func (_e *MockRepository_Expecter) CreateCustomerOrder(_a0 interface{}, _a1 interface{}) *MockRepository_CreateCustomerOrder_Call {
	return &MockRepository_CreateCustomerOrder_Call{Call: _e.mock.On("CreateCustomerOrder", _a0, _a1)}
}

func (_c *MockRepository_CreateCustomerOrder_Call) Run(run func(_a0 context.Context, _a1 *model.CustomerOrder)) *MockRepository_CreateCustomerOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.CustomerOrder))
	})
	return _c
}

func (_c *MockRepository_CreateCustomerOrder_Call) Return(_a0 error) *MockRepository_CreateCustomerOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_CreateCustomerOrder_Call) RunAndReturn(run func(context.Context, *model.CustomerOrder) error) *MockRepository_CreateCustomerOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetCartByConstraint provides a mock function with given fields: ctx, accountID, productID
func (_m *MockRepository) GetCartByConstraint(ctx context.Context, accountID uuid.UUID, productID uuid.UUID) (model.Cart, error) {
	ret := _m.Called(ctx, accountID, productID)

	var r0 model.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) (model.Cart, error)); ok {
		return rf(ctx, accountID, productID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) model.Cart); ok {
		r0 = rf(ctx, accountID, productID)
	} else {
		r0 = ret.Get(0).(model.Cart)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(ctx, accountID, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetCartByConstraint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCartByConstraint'
type MockRepository_GetCartByConstraint_Call struct {
	*mock.Call
}

// GetCartByConstraint is a helper method to define mock.On call
//   - ctx context.Context
//   - accountID uuid.UUID
//   - productID uuid.UUID
func (_e *MockRepository_Expecter) GetCartByConstraint(ctx interface{}, accountID interface{}, productID interface{}) *MockRepository_GetCartByConstraint_Call {
	return &MockRepository_GetCartByConstraint_Call{Call: _e.mock.On("GetCartByConstraint", ctx, accountID, productID)}
}

func (_c *MockRepository_GetCartByConstraint_Call) Run(run func(ctx context.Context, accountID uuid.UUID, productID uuid.UUID)) *MockRepository_GetCartByConstraint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *MockRepository_GetCartByConstraint_Call) Return(_a0 model.Cart, _a1 error) *MockRepository_GetCartByConstraint_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetCartByConstraint_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID) (model.Cart, error)) *MockRepository_GetCartByConstraint_Call {
	_c.Call.Return(run)
	return _c
}

// GetListCart provides a mock function with given fields: ctx, accountID
func (_m *MockRepository) GetListCart(ctx context.Context, accountID uuid.UUID) ([]model.Cart, error) {
	ret := _m.Called(ctx, accountID)

	var r0 []model.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]model.Cart, error)); ok {
		return rf(ctx, accountID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []model.Cart); ok {
		r0 = rf(ctx, accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetListCart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListCart'
type MockRepository_GetListCart_Call struct {
	*mock.Call
}

// GetListCart is a helper method to define mock.On call
//   - ctx context.Context
//   - accountID uuid.UUID
func (_e *MockRepository_Expecter) GetListCart(ctx interface{}, accountID interface{}) *MockRepository_GetListCart_Call {
	return &MockRepository_GetListCart_Call{Call: _e.mock.On("GetListCart", ctx, accountID)}
}

func (_c *MockRepository_GetListCart_Call) Run(run func(ctx context.Context, accountID uuid.UUID)) *MockRepository_GetListCart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockRepository_GetListCart_Call) Return(_a0 []model.Cart, _a1 error) *MockRepository_GetListCart_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetListCart_Call) RunAndReturn(run func(context.Context, uuid.UUID) ([]model.Cart, error)) *MockRepository_GetListCart_Call {
	_c.Call.Return(run)
	return _c
}

// GetListOrdered provides a mock function with given fields: ctx, accountID, order, paginator
func (_m *MockRepository) GetListOrdered(ctx context.Context, accountID uuid.UUID, order []string, paginator codetype.Paginator) ([]model.CustomerOrder, int64, error) {
	ret := _m.Called(ctx, accountID, order, paginator)

	var r0 []model.CustomerOrder
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []string, codetype.Paginator) ([]model.CustomerOrder, int64, error)); ok {
		return rf(ctx, accountID, order, paginator)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []string, codetype.Paginator) []model.CustomerOrder); ok {
		r0 = rf(ctx, accountID, order, paginator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.CustomerOrder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, []string, codetype.Paginator) int64); ok {
		r1 = rf(ctx, accountID, order, paginator)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uuid.UUID, []string, codetype.Paginator) error); ok {
		r2 = rf(ctx, accountID, order, paginator)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepository_GetListOrdered_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListOrdered'
type MockRepository_GetListOrdered_Call struct {
	*mock.Call
}

// GetListOrdered is a helper method to define mock.On call
//   - ctx context.Context
//   - accountID uuid.UUID
//   - order []string
//   - paginator codetype.Paginator
func (_e *MockRepository_Expecter) GetListOrdered(ctx interface{}, accountID interface{}, order interface{}, paginator interface{}) *MockRepository_GetListOrdered_Call {
	return &MockRepository_GetListOrdered_Call{Call: _e.mock.On("GetListOrdered", ctx, accountID, order, paginator)}
}

func (_c *MockRepository_GetListOrdered_Call) Run(run func(ctx context.Context, accountID uuid.UUID, order []string, paginator codetype.Paginator)) *MockRepository_GetListOrdered_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]string), args[3].(codetype.Paginator))
	})
	return _c
}

func (_c *MockRepository_GetListOrdered_Call) Return(_a0 []model.CustomerOrder, _a1 int64, _a2 error) *MockRepository_GetListOrdered_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepository_GetListOrdered_Call) RunAndReturn(run func(context.Context, uuid.UUID, []string, codetype.Paginator) ([]model.CustomerOrder, int64, error)) *MockRepository_GetListOrdered_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveFromCart provides a mock function with given fields: ctx, accountID, productID, qty
func (_m *MockRepository) RemoveFromCart(ctx context.Context, accountID uuid.UUID, productID uuid.UUID, qty int64) error {
	ret := _m.Called(ctx, accountID, productID, qty)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, int64) error); ok {
		r0 = rf(ctx, accountID, productID, qty)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_RemoveFromCart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveFromCart'
type MockRepository_RemoveFromCart_Call struct {
	*mock.Call
}

// RemoveFromCart is a helper method to define mock.On call
//   - ctx context.Context
//   - accountID uuid.UUID
//   - productID uuid.UUID
//   - qty int64
func (_e *MockRepository_Expecter) RemoveFromCart(ctx interface{}, accountID interface{}, productID interface{}, qty interface{}) *MockRepository_RemoveFromCart_Call {
	return &MockRepository_RemoveFromCart_Call{Call: _e.mock.On("RemoveFromCart", ctx, accountID, productID, qty)}
}

func (_c *MockRepository_RemoveFromCart_Call) Run(run func(ctx context.Context, accountID uuid.UUID, productID uuid.UUID, qty int64)) *MockRepository_RemoveFromCart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID), args[3].(int64))
	})
	return _c
}

func (_c *MockRepository_RemoveFromCart_Call) Return(_a0 error) *MockRepository_RemoveFromCart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_RemoveFromCart_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID, int64) error) *MockRepository_RemoveFromCart_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertCart provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) UpsertCart(_a0 context.Context, _a1 *model.Cart) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Cart) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_UpsertCart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertCart'
type MockRepository_UpsertCart_Call struct {
	*mock.Call
}

// UpsertCart is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Cart
func (_e *MockRepository_Expecter) UpsertCart(_a0 interface{}, _a1 interface{}) *MockRepository_UpsertCart_Call {
	return &MockRepository_UpsertCart_Call{Call: _e.mock.On("UpsertCart", _a0, _a1)}
}

func (_c *MockRepository_UpsertCart_Call) Run(run func(_a0 context.Context, _a1 *model.Cart)) *MockRepository_UpsertCart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Cart))
	})
	return _c
}

func (_c *MockRepository_UpsertCart_Call) Return(_a0 error) *MockRepository_UpsertCart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_UpsertCart_Call) RunAndReturn(run func(context.Context, *model.Cart) error) *MockRepository_UpsertCart_Call {
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