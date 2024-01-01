package repository

import "github.com/stretchr/testify/mock"

// MockRentalRepository is a mock type for model.RentalRepository
type MockRentalRepository struct {
	mock.Mock
}

// GetRentalByCarID provides a mock function with given fields: ctx, carID
func (_m *MockRentalRepository) GetRentalByCarID(ctx interface{}, carID int) (interface{}, error) {
	ret := _m.Called(ctx, carID)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}, int) interface{}); ok {
		r0 = rf(ctx, carID)
	} else {
		r0 = ret.Get(0)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, int) error); ok {
		r1 = rf(ctx, carID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRental provides a mock function with given fields: ctx, rentalDto
func (_m *MockRentalRepository) CreateRental(ctx interface{}, rentalDto interface{}) (int, error) {
	ret := _m.Called(ctx, rentalDto)

	var r0 int
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) int); ok {
		r0 = rf(ctx, rentalDto)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}) error); ok {
		r1 = rf(ctx, rentalDto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRentalKilometersDriven provides a mock function with given fields: ctx, carID, kilometers_driven
func (_m *MockRentalRepository) UpdateRentalKilometersDriven(ctx interface{}, carID int, kilometers_driven int) error {
	ret := _m.Called(ctx, carID, kilometers_driven)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int, int) error); ok {
		r0 = rf(ctx, carID, kilometers_driven)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRentalEndDate provides a mock function with given fields: ctx, carID, end_date
func (_m *MockRentalRepository) UpdateRentalEndDate(ctx interface{}, carID int, end_date interface{}) error {
	ret := _m.Called(ctx, carID, end_date)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int, interface{}) error); ok {
		r0 = rf(ctx, carID, end_date)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

