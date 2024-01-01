package constant

import "errors"

var (
	// ErrNotFound is returned when a resource is not found
	ErrNotFound = errors.New("resource not found")

	// ErrAlreadyRented is returned when a car is already rented
	ErrAlreadyRented = errors.New("car is already rented")

	// ErrInvalidMileage is returned when the mileage provided is invalid
	ErrInvalidMileage = errors.New("mileage provided is invalid")

	// ErrNotRented is returned when a car is not rented
	ErrNotRented = errors.New("car is not rented")

	// ErrDuplicateRegistration is returned when a car is already registered
	ErrDuplicateRegistration = errors.New("car is already registered")
)
