// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import vtwilio "github.com/vendasta/vtwilio-go"

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// AvailablePhoneNumbers provides a mock function with given fields: countryCode, opts
func (_m *Interface) AvailablePhoneNumbers(countryCode string, opts ...vtwilio.AvailableOption) (*vtwilio.AvailablePhoneNumbers, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, countryCode)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *vtwilio.AvailablePhoneNumbers
	if rf, ok := ret.Get(0).(func(string, ...vtwilio.AvailableOption) *vtwilio.AvailablePhoneNumbers); ok {
		r0 = rf(countryCode, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vtwilio.AvailablePhoneNumbers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...vtwilio.AvailableOption) error); ok {
		r1 = rf(countryCode, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessage provides a mock function with given fields: messageSID
func (_m *Interface) GetMessage(messageSID string) (*vtwilio.Message, error) {
	ret := _m.Called(messageSID)

	var r0 *vtwilio.Message
	if rf, ok := ret.Get(0).(func(string) *vtwilio.Message); ok {
		r0 = rf(messageSID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vtwilio.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(messageSID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncomingPhoneNumber provides a mock function with given fields: number, opts
func (_m *Interface) IncomingPhoneNumber(number string, opts ...vtwilio.IncomingPhoneNumberOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, number)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...vtwilio.IncomingPhoneNumberOption) error); ok {
		r0 = rf(number, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListMessages provides a mock function with given fields: opts
func (_m *Interface) ListMessages(opts ...vtwilio.ListOption) (*vtwilio.List, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *vtwilio.List
	if rf, ok := ret.Get(0).(func(...vtwilio.ListOption) *vtwilio.List); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vtwilio.List)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...vtwilio.ListOption) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: message, to, opts
func (_m *Interface) SendMessage(message string, to string, opts ...vtwilio.SendOption) (*vtwilio.Message, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, message, to)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *vtwilio.Message
	if rf, ok := ret.Get(0).(func(string, string, ...vtwilio.SendOption) *vtwilio.Message); ok {
		r0 = rf(message, to, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vtwilio.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, ...vtwilio.SendOption) error); ok {
		r1 = rf(message, to, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPhoneNumber provides a mock function with given fields: n
func (_m *Interface) SetPhoneNumber(n string) *vtwilio.VTwilio {
	ret := _m.Called(n)

	var r0 *vtwilio.VTwilio
	if rf, ok := ret.Get(0).(func(string) *vtwilio.VTwilio); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vtwilio.VTwilio)
		}
	}

	return r0
}
