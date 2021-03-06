// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/aarnphm/dha-ps/ingress/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// AttributesRepository is an autogenerated mock type for the AttributesRepository type
type AttributesRepository struct {
	mock.Mock
}

// FetchAll provides a mock function with given fields: ctx
func (_m *AttributesRepository) FetchAll(ctx context.Context) ([]models.Attributes, error) {
	ret := _m.Called(ctx)

	var r0 []models.Attributes
	if rf, ok := ret.Get(0).(func(context.Context) []models.Attributes); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Attributes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
