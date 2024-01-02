package model

import "time"

// Rental is the model for a rental.
type Rental struct {
	// ID is the unique identifier of the rental.
	ID               int        `json:"id" db:"id"`
	// UserID is the unique identifier of the user.
	CarID            int        `json:"car_id" db:"car_id" validate:"required"`
	// StartDate is the start date of the rental.
	StartDate        time.Time  `json:"start_date" db:"start_date" validate:"required"`
	// EndDate is the end date of the rental.
	EndDate          *time.Time `json:"end_date" db:"end_date"`
	// KilometersDriven is the kilometers driven of the car.
	KilometersDriven *int       `json:"kilometers_driven" db:"kilometers_driven"`
}
