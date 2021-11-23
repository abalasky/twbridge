// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/dstdfx/twbridge/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// WhatsappClient is an autogenerated mock type for the WhatsappClient type
type WhatsappClient struct {
	mock.Mock
}

// GetContacts provides a mock function with given fields:
func (_m *WhatsappClient) GetContacts() map[string]domain.WhatsappContact {
	ret := _m.Called()

	var r0 map[string]domain.WhatsappContact
	if rf, ok := ret.Get(0).(func() map[string]domain.WhatsappContact); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]domain.WhatsappContact)
		}
	}

	return r0
}

// Restore provides a mock function with given fields:
func (_m *WhatsappClient) Restore() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: msg
func (_m *WhatsappClient) Send(msg domain.WhatsappMessage) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.WhatsappMessage) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
