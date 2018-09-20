// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import api "github.com/cdempsie/payments-example/api"
import mock "github.com/stretchr/testify/mock"

// PaymentStore is an autogenerated mock type for the PaymentStore type
type PaymentStore struct {
	mock.Mock
}

// Create provides a mock function with given fields: payment
func (_m *PaymentStore) Create(payment *api.Payment) error {
	ret := _m.Called(payment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*api.Payment) error); ok {
		r0 = rf(payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: paymentUID
func (_m *PaymentStore) Delete(paymentUID string) error {
	ret := _m.Called(paymentUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(paymentUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *PaymentStore) List() (*api.ListHolder, error) {
	ret := _m.Called()

	var r0 *api.ListHolder
	if rf, ok := ret.Get(0).(func() *api.ListHolder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ListHolder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Load provides a mock function with given fields: paymentUID
func (_m *PaymentStore) Load(paymentUID string) (*api.Payment, error) {
	ret := _m.Called(paymentUID)

	var r0 *api.Payment
	if rf, ok := ret.Get(0).(func(string) *api.Payment); ok {
		r0 = rf(paymentUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(paymentUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: payment
func (_m *PaymentStore) Update(payment *api.Payment) error {
	ret := _m.Called(payment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*api.Payment) error); ok {
		r0 = rf(payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}