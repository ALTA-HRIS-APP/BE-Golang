// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	target "be_golang/klp3/features/target"

	mock "github.com/stretchr/testify/mock"
)

// TargetData is an autogenerated mock type for the TargetDataInterface type
type TargetData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: targetID
func (_m *TargetData) Delete(targetID string) error {
	ret := _m.Called(targetID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(targetID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByIDAPI provides a mock function with given fields: idUser
func (_m *TargetData) GetUserByIDAPI(idUser string) (target.PenggunaEntity, error) {
	ret := _m.Called(idUser)

	var r0 target.PenggunaEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (target.PenggunaEntity, error)); ok {
		return rf(idUser)
	}
	if rf, ok := ret.Get(0).(func(string) target.PenggunaEntity); ok {
		r0 = rf(idUser)
	} else {
		r0 = ret.Get(0).(target.PenggunaEntity)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: input
func (_m *TargetData) Insert(input target.TargetEntity) (string, error) {
	ret := _m.Called(input)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(target.TargetEntity) (string, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(target.TargetEntity) string); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(target.TargetEntity) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Select provides a mock function with given fields: targetID
func (_m *TargetData) Select(targetID string) (target.TargetEntity, error) {
	ret := _m.Called(targetID)

	var r0 target.TargetEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (target.TargetEntity, error)); ok {
		return rf(targetID)
	}
	if rf, ok := ret.Get(0).(func(string) target.TargetEntity); ok {
		r0 = rf(targetID)
	} else {
		r0 = ret.Get(0).(target.TargetEntity)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(targetID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAll provides a mock function with given fields: param
func (_m *TargetData) SelectAll(param target.QueryParam) (int64, []target.TargetEntity, error) {
	ret := _m.Called(param)

	var r0 int64
	var r1 []target.TargetEntity
	var r2 error
	if rf, ok := ret.Get(0).(func(target.QueryParam) (int64, []target.TargetEntity, error)); ok {
		return rf(param)
	}
	if rf, ok := ret.Get(0).(func(target.QueryParam) int64); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(target.QueryParam) []target.TargetEntity); ok {
		r1 = rf(param)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]target.TargetEntity)
		}
	}

	if rf, ok := ret.Get(2).(func(target.QueryParam) error); ok {
		r2 = rf(param)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectAllKaryawan provides a mock function with given fields: idUser, param
func (_m *TargetData) SelectAllKaryawan(idUser string, param target.QueryParam) (int64, []target.TargetEntity, error) {
	ret := _m.Called(idUser, param)

	var r0 int64
	var r1 []target.TargetEntity
	var r2 error
	if rf, ok := ret.Get(0).(func(string, target.QueryParam) (int64, []target.TargetEntity, error)); ok {
		return rf(idUser, param)
	}
	if rf, ok := ret.Get(0).(func(string, target.QueryParam) int64); ok {
		r0 = rf(idUser, param)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, target.QueryParam) []target.TargetEntity); ok {
		r1 = rf(idUser, param)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]target.TargetEntity)
		}
	}

	if rf, ok := ret.Get(2).(func(string, target.QueryParam) error); ok {
		r2 = rf(idUser, param)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: targetID, targetData
func (_m *TargetData) Update(targetID string, targetData target.TargetEntity) error {
	ret := _m.Called(targetID, targetData)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, target.TargetEntity) error); ok {
		r0 = rf(targetID, targetData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTargetData creates a new instance of TargetData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTargetData(t interface {
	mock.TestingT
	Cleanup(func())
}) *TargetData {
	mock := &TargetData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}