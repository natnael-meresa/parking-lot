package mocks

import (
	"context"
	"parking-lot/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockCarRepository is a mock type for model.CarRepository
type MockCarRepository struct {
	mock.Mock
}

// GetCarByRegistration provides a mock function with given fields: ctx, registration
func (_m *MockCarRepository) GetCarByRegistration(ctx context.Context, registration string) (*model.Car, error) {
	ret := _m.Called(ctx, registration)

	var r0 model.Car
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Car); ok {
		r0 = rf(ctx, registration)
	} else {
		r0 = ret.Get(0).(model.Car)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, registration)
	} else {
		r1 = ret.Error(1)
	}

	return &r0, r1
}

// CreateCar provides a mock function with given fields: ctx, carModel
func (_m *MockCarRepository) CreateCar(ctx context.Context, carModel model.Car) (int, error) {
	ret := _m.Called(ctx, carModel)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, model.Car) int); ok {
		r0 = rf(ctx, carModel)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Car) error); ok {
		r1 = rf(ctx, carModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCars provides a mock function with given fields: ctx
func (_m *MockCarRepository) ListCars(ctx context.Context) ([]model.Car, error) {
	ret := _m.Called(ctx)

	var r0 []model.Car
	if rf, ok := ret.Get(0).(func(context.Context) []model.Car); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Car)
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

// UpdateCarMileage provides a mock function with given fields: ctx, registration, mileage
func (_m *MockCarRepository) UpdateCarMileage(ctx context.Context, registration string, mileage int) error {
	ret := _m.Called(ctx, registration, mileage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, registration, mileage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCarStatus provides a mock function with given fields: ctx, registration, status
func (_m *MockCarRepository) UpdateCarStatus(ctx context.Context, registration string, status string) error {
	ret := _m.Called(ctx, registration, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, registration, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}