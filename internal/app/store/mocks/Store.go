// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	store "hack/internal/app/store"

	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Compaign provides a mock function with given fields:
func (_m *Store) Compaign() store.CompaignRepository {
	ret := _m.Called()

	var r0 store.CompaignRepository
	if rf, ok := ret.Get(0).(func() store.CompaignRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.CompaignRepository)
		}
	}

	return r0
}

// File provides a mock function with given fields:
func (_m *Store) File() store.FileRepository {
	ret := _m.Called()

	var r0 store.FileRepository
	if rf, ok := ret.Get(0).(func() store.FileRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.FileRepository)
		}
	}

	return r0
}

// Lead provides a mock function with given fields:
func (_m *Store) Lead() store.LeadRepository {
	ret := _m.Called()

	var r0 store.LeadRepository
	if rf, ok := ret.Get(0).(func() store.LeadRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.LeadRepository)
		}
	}

	return r0
}

// Region provides a mock function with given fields:
func (_m *Store) Region() store.RegionRepository {
	ret := _m.Called()

	var r0 store.RegionRepository
	if rf, ok := ret.Get(0).(func() store.RegionRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.RegionRepository)
		}
	}

	return r0
}

// Telegram provides a mock function with given fields:
func (_m *Store) Telegram() store.TelegramRepository {
	ret := _m.Called()

	var r0 store.TelegramRepository
	if rf, ok := ret.Get(0).(func() store.TelegramRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.TelegramRepository)
		}
	}

	return r0
}
