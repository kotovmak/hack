// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	context "context"
	model "hack/internal/app/model"

	mock "github.com/stretchr/testify/mock"
)

// CompaignRepository is an autogenerated mock type for the CompaignRepository type
type CompaignRepository struct {
	mock.Mock
}

// GetList provides a mock function with given fields: _a0
func (_m *CompaignRepository) GetList(_a0 context.Context) ([]model.Compaign, error) {
	ret := _m.Called(_a0)

	var r0 []model.Compaign
	if rf, ok := ret.Get(0).(func(context.Context) []model.Compaign); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Compaign)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
