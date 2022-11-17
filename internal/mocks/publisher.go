// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	git "github.com/go-git/go-git/v5"
	mock "github.com/stretchr/testify/mock"
)

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

// GetRepo provides a mock function with given fields:
func (_m *Publisher) GetRepo() (*git.Repository, error) {
	ret := _m.Called()

	var r0 *git.Repository
	if rf, ok := ret.Get(0).(func() *git.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*git.Repository)
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

// Publish provides a mock function with given fields: r, subnetName, vmName, subnetYAML, vmYAML
func (_m *Publisher) Publish(r *git.Repository, subnetName string, vmName string, subnetYAML []byte, vmYAML []byte) error {
	ret := _m.Called(r, subnetName, vmName, subnetYAML, vmYAML)

	var r0 error
	if rf, ok := ret.Get(0).(func(*git.Repository, string, string, []byte, []byte) error); ok {
		r0 = rf(r, subnetName, vmName, subnetYAML, vmYAML)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewPublisher creates a new instance of Publisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPublisher(t mockConstructorTestingTNewPublisher) *Publisher {
	mock := &Publisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
