// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// BeaconBlock is an autogenerated mock type for the BeaconBlock type
type BeaconBlock struct {
	mock.Mock
}

type BeaconBlock_Expecter struct {
	mock *mock.Mock
}

func (_m *BeaconBlock) EXPECT() *BeaconBlock_Expecter {
	return &BeaconBlock_Expecter{mock: &_m.Mock}
}

// GetExecutionNumber provides a mock function with given fields:
func (_m *BeaconBlock) GetExecutionNumber() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionNumber")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlock_GetExecutionNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExecutionNumber'
type BeaconBlock_GetExecutionNumber_Call struct {
	*mock.Call
}

// GetExecutionNumber is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetExecutionNumber() *BeaconBlock_GetExecutionNumber_Call {
	return &BeaconBlock_GetExecutionNumber_Call{Call: _e.mock.On("GetExecutionNumber")}
}

func (_c *BeaconBlock_GetExecutionNumber_Call) Run(run func()) *BeaconBlock_GetExecutionNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetExecutionNumber_Call) Return(_a0 math.U64) *BeaconBlock_GetExecutionNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetExecutionNumber_Call) RunAndReturn(run func() math.U64) *BeaconBlock_GetExecutionNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetParentBlockRoot provides a mock function with given fields:
func (_m *BeaconBlock) GetParentBlockRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetParentBlockRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlock_GetParentBlockRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParentBlockRoot'
type BeaconBlock_GetParentBlockRoot_Call struct {
	*mock.Call
}

// GetParentBlockRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetParentBlockRoot() *BeaconBlock_GetParentBlockRoot_Call {
	return &BeaconBlock_GetParentBlockRoot_Call{Call: _e.mock.On("GetParentBlockRoot")}
}

func (_c *BeaconBlock_GetParentBlockRoot_Call) Run(run func()) *BeaconBlock_GetParentBlockRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetParentBlockRoot_Call) Return(_a0 common.Root) *BeaconBlock_GetParentBlockRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetParentBlockRoot_Call) RunAndReturn(run func() common.Root) *BeaconBlock_GetParentBlockRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetSlot provides a mock function with given fields:
func (_m *BeaconBlock) GetSlot() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlock_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type BeaconBlock_GetSlot_Call struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetSlot() *BeaconBlock_GetSlot_Call {
	return &BeaconBlock_GetSlot_Call{Call: _e.mock.On("GetSlot")}
}

func (_c *BeaconBlock_GetSlot_Call) Run(run func()) *BeaconBlock_GetSlot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetSlot_Call) Return(_a0 math.U64) *BeaconBlock_GetSlot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetSlot_Call) RunAndReturn(run func() math.U64) *BeaconBlock_GetSlot_Call {
	_c.Call.Return(run)
	return _c
}

// GetStateRoot provides a mock function with given fields:
func (_m *BeaconBlock) GetStateRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStateRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlock_GetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStateRoot'
type BeaconBlock_GetStateRoot_Call struct {
	*mock.Call
}

// GetStateRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetStateRoot() *BeaconBlock_GetStateRoot_Call {
	return &BeaconBlock_GetStateRoot_Call{Call: _e.mock.On("GetStateRoot")}
}

func (_c *BeaconBlock_GetStateRoot_Call) Run(run func()) *BeaconBlock_GetStateRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetStateRoot_Call) Return(_a0 common.Root) *BeaconBlock_GetStateRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetStateRoot_Call) RunAndReturn(run func() common.Root) *BeaconBlock_GetStateRoot_Call {
	_c.Call.Return(run)
	return _c
}

// HashTreeRoot provides a mock function with given fields:
func (_m *BeaconBlock) HashTreeRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HashTreeRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlock_HashTreeRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashTreeRoot'
type BeaconBlock_HashTreeRoot_Call struct {
	*mock.Call
}

// HashTreeRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) HashTreeRoot() *BeaconBlock_HashTreeRoot_Call {
	return &BeaconBlock_HashTreeRoot_Call{Call: _e.mock.On("HashTreeRoot")}
}

func (_c *BeaconBlock_HashTreeRoot_Call) Run(run func()) *BeaconBlock_HashTreeRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_HashTreeRoot_Call) Return(_a0 common.Root) *BeaconBlock_HashTreeRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_HashTreeRoot_Call) RunAndReturn(run func() common.Root) *BeaconBlock_HashTreeRoot_Call {
	_c.Call.Return(run)
	return _c
}

// NewBeaconBlock creates a new instance of BeaconBlock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconBlock(t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconBlock {
	mock := &BeaconBlock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
