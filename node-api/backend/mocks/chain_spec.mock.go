// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	math "github.com/berachain/beacon-kit/primitives/math"
	mock "github.com/stretchr/testify/mock"
)

// ChainSpec is an autogenerated mock type for the ChainSpec type
type ChainSpec struct {
	mock.Mock
}

type ChainSpec_Expecter struct {
	mock *mock.Mock
}

func (_m *ChainSpec) EXPECT() *ChainSpec_Expecter {
	return &ChainSpec_Expecter{mock: &_m.Mock}
}

// EpochsPerHistoricalVector provides a mock function with given fields:
func (_m *ChainSpec) EpochsPerHistoricalVector() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EpochsPerHistoricalVector")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ChainSpec_EpochsPerHistoricalVector_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EpochsPerHistoricalVector'
type ChainSpec_EpochsPerHistoricalVector_Call struct {
	*mock.Call
}

// EpochsPerHistoricalVector is a helper method to define mock.On call
func (_e *ChainSpec_Expecter) EpochsPerHistoricalVector() *ChainSpec_EpochsPerHistoricalVector_Call {
	return &ChainSpec_EpochsPerHistoricalVector_Call{Call: _e.mock.On("EpochsPerHistoricalVector")}
}

func (_c *ChainSpec_EpochsPerHistoricalVector_Call) Run(run func()) *ChainSpec_EpochsPerHistoricalVector_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ChainSpec_EpochsPerHistoricalVector_Call) Return(_a0 uint64) *ChainSpec_EpochsPerHistoricalVector_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChainSpec_EpochsPerHistoricalVector_Call) RunAndReturn(run func() uint64) *ChainSpec_EpochsPerHistoricalVector_Call {
	_c.Call.Return(run)
	return _c
}

// MaxBlobsPerBlock provides a mock function with given fields:
func (_m *ChainSpec) MaxBlobsPerBlock() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MaxBlobsPerBlock")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ChainSpec_MaxBlobsPerBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MaxBlobsPerBlock'
type ChainSpec_MaxBlobsPerBlock_Call struct {
	*mock.Call
}

// MaxBlobsPerBlock is a helper method to define mock.On call
func (_e *ChainSpec_Expecter) MaxBlobsPerBlock() *ChainSpec_MaxBlobsPerBlock_Call {
	return &ChainSpec_MaxBlobsPerBlock_Call{Call: _e.mock.On("MaxBlobsPerBlock")}
}

func (_c *ChainSpec_MaxBlobsPerBlock_Call) Run(run func()) *ChainSpec_MaxBlobsPerBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ChainSpec_MaxBlobsPerBlock_Call) Return(_a0 uint64) *ChainSpec_MaxBlobsPerBlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChainSpec_MaxBlobsPerBlock_Call) RunAndReturn(run func() uint64) *ChainSpec_MaxBlobsPerBlock_Call {
	_c.Call.Return(run)
	return _c
}

// MinEpochsForBlobsSidecarsRequest provides a mock function with given fields:
func (_m *ChainSpec) MinEpochsForBlobsSidecarsRequest() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MinEpochsForBlobsSidecarsRequest")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// ChainSpec_MinEpochsForBlobsSidecarsRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MinEpochsForBlobsSidecarsRequest'
type ChainSpec_MinEpochsForBlobsSidecarsRequest_Call struct {
	*mock.Call
}

// MinEpochsForBlobsSidecarsRequest is a helper method to define mock.On call
func (_e *ChainSpec_Expecter) MinEpochsForBlobsSidecarsRequest() *ChainSpec_MinEpochsForBlobsSidecarsRequest_Call {
	return &ChainSpec_MinEpochsForBlobsSidecarsRequest_Call{Call: _e.mock.On("MinEpochsForBlobsSidecarsRequest")}
}

func (_c *ChainSpec_MinEpochsForBlobsSidecarsRequest_Call) Run(run func()) *ChainSpec_MinEpochsForBlobsSidecarsRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ChainSpec_MinEpochsForBlobsSidecarsRequest_Call) Return(_a0 math.U64) *ChainSpec_MinEpochsForBlobsSidecarsRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChainSpec_MinEpochsForBlobsSidecarsRequest_Call) RunAndReturn(run func() math.U64) *ChainSpec_MinEpochsForBlobsSidecarsRequest_Call {
	_c.Call.Return(run)
	return _c
}

// SlotToEpoch provides a mock function with given fields: slot
func (_m *ChainSpec) SlotToEpoch(slot math.U64) math.U64 {
	ret := _m.Called(slot)

	if len(ret) == 0 {
		panic("no return value specified for SlotToEpoch")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func(math.U64) math.U64); ok {
		r0 = rf(slot)
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// ChainSpec_SlotToEpoch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SlotToEpoch'
type ChainSpec_SlotToEpoch_Call struct {
	*mock.Call
}

// SlotToEpoch is a helper method to define mock.On call
//   - slot math.U64
func (_e *ChainSpec_Expecter) SlotToEpoch(slot interface{}) *ChainSpec_SlotToEpoch_Call {
	return &ChainSpec_SlotToEpoch_Call{Call: _e.mock.On("SlotToEpoch", slot)}
}

func (_c *ChainSpec_SlotToEpoch_Call) Run(run func(slot math.U64)) *ChainSpec_SlotToEpoch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64))
	})
	return _c
}

func (_c *ChainSpec_SlotToEpoch_Call) Return(_a0 math.U64) *ChainSpec_SlotToEpoch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChainSpec_SlotToEpoch_Call) RunAndReturn(run func(math.U64) math.U64) *ChainSpec_SlotToEpoch_Call {
	_c.Call.Return(run)
	return _c
}

// SlotsPerHistoricalRoot provides a mock function with given fields:
func (_m *ChainSpec) SlotsPerHistoricalRoot() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SlotsPerHistoricalRoot")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ChainSpec_SlotsPerHistoricalRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SlotsPerHistoricalRoot'
type ChainSpec_SlotsPerHistoricalRoot_Call struct {
	*mock.Call
}

// SlotsPerHistoricalRoot is a helper method to define mock.On call
func (_e *ChainSpec_Expecter) SlotsPerHistoricalRoot() *ChainSpec_SlotsPerHistoricalRoot_Call {
	return &ChainSpec_SlotsPerHistoricalRoot_Call{Call: _e.mock.On("SlotsPerHistoricalRoot")}
}

func (_c *ChainSpec_SlotsPerHistoricalRoot_Call) Run(run func()) *ChainSpec_SlotsPerHistoricalRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ChainSpec_SlotsPerHistoricalRoot_Call) Return(_a0 uint64) *ChainSpec_SlotsPerHistoricalRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChainSpec_SlotsPerHistoricalRoot_Call) RunAndReturn(run func() uint64) *ChainSpec_SlotsPerHistoricalRoot_Call {
	_c.Call.Return(run)
	return _c
}

// WithinDAPeriod provides a mock function with given fields: block, current
func (_m *ChainSpec) WithinDAPeriod(block math.U64, current math.U64) bool {
	ret := _m.Called(block, current)

	if len(ret) == 0 {
		panic("no return value specified for WithinDAPeriod")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(math.U64, math.U64) bool); ok {
		r0 = rf(block, current)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ChainSpec_WithinDAPeriod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithinDAPeriod'
type ChainSpec_WithinDAPeriod_Call struct {
	*mock.Call
}

// WithinDAPeriod is a helper method to define mock.On call
//   - block math.U64
//   - current math.U64
func (_e *ChainSpec_Expecter) WithinDAPeriod(block interface{}, current interface{}) *ChainSpec_WithinDAPeriod_Call {
	return &ChainSpec_WithinDAPeriod_Call{Call: _e.mock.On("WithinDAPeriod", block, current)}
}

func (_c *ChainSpec_WithinDAPeriod_Call) Run(run func(block math.U64, current math.U64)) *ChainSpec_WithinDAPeriod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64), args[1].(math.U64))
	})
	return _c
}

func (_c *ChainSpec_WithinDAPeriod_Call) Return(_a0 bool) *ChainSpec_WithinDAPeriod_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChainSpec_WithinDAPeriod_Call) RunAndReturn(run func(math.U64, math.U64) bool) *ChainSpec_WithinDAPeriod_Call {
	_c.Call.Return(run)
	return _c
}

// NewChainSpec creates a new instance of ChainSpec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChainSpec(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChainSpec {
	mock := &ChainSpec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
