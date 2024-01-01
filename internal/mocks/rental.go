package mocks

import (
	"context"
	"parking-lot/internal/model"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockRentalRepository is a mock type for model.RentalRepository
type MockRentalRepository struct {
	mock.Mock
}

// GetRentalByCarID provides a mock function with given fields: ctx, carID
func (_m *MockRentalRepository) GetRentalByCarID(ctx context.Context, carID int) (*model.Rental, error) {
	ret := _m.Called(ctx, carID)

	var r0 *model.Rental
	if rf, ok := ret.Get(0).(func(context.Context, int) *model.Rental); ok {
		r0 = rf(ctx, carID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Rental)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, carID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRental provides a mock function with given fields: ctx, rentalDto
func (_m *MockRentalRepository) CreateRental(ctx context.Context, rentalDto model.Rental) (int, error) {
	ret := _m.Called(ctx, rentalDto)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, model.Rental) int); ok {
		r0 = rf(ctx, rentalDto)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Rental) error); ok {
		r1 = rf(ctx, rentalDto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRentalKilometersDriven provides a mock function with given fields: ctx, carID, kilometers_driven
func (_m *MockRentalRepository) UpdateRentalKilometersDriven(ctx context.Context, carID int, kilometers_driven int) error {
	ret := _m.Called(ctx, carID, kilometers_driven)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, carID, kilometers_driven)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRentalEndDate provides a mock function with given fields: ctx, carID, end_date
func (_m *MockRentalRepository) UpdateRentalEndDate(ctx context.Context, carID int, end_date time.Time) error {
	ret := _m.Called(ctx, carID, end_date)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, time.Time) error); ok {
		r0 = rf(ctx, carID, end_date)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListRentals provides a mock function with given fields: ctx
func (_m *MockRentalRepository) ListRentals(ctx context.Context) ([]model.Rental, error) {
	ret := _m.Called(ctx)

	var r0 []model.Rental
	if rf, ok := ret.Get(0).(func(context.Context) []model.Rental); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Rental)
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
