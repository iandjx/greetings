// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Doer is an autogenerated mock type for the Doer type
type Doer struct {
	mock.Mock
}

// DoSomething provides a mock function with given fields: _a0, _a1
func (_m *Doer) DoSomething(_a0 int, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
