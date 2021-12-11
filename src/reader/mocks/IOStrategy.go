// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IOStrategy is an autogenerated mock type for the IOStrategy type
type IOStrategy struct {
	mock.Mock
}

// Read provides a mock function with given fields: input
func (_m *IOStrategy) Read(input string) ([][]string, error) {
	ret := _m.Called(input)

	var r0 [][]string
	if rf, ok := ret.Get(0).(func(string) [][]string); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
