// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	auth "go-echo-api/auth"
	entity "go-echo-api/entity"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Login provides a mock function with given fields: email
func (_m *Repository) Login(email string) (entity.User, error) {
	ret := _m.Called(email)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: dto
func (_m *Repository) Register(dto auth.RegisterDto) (entity.User, error) {
	ret := _m.Called(dto)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(auth.RegisterDto) entity.User); ok {
		r0 = rf(dto)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(auth.RegisterDto) error); ok {
		r1 = rf(dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}