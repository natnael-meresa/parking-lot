package model

import "time"

type Rental struct {
	ID               int        `json:"id" db:"id"`
	CarID            int        `json:"car_id" db:"car_id" validate:"required"`
	StartDate        time.Time  `json:"start_date" db:"start_date" validate:"required"`
	EndDate          *time.Time `json:"end_date" db:"end_date"`
	KilometersDriven *int       `json:"kilometers_driven" db:"kilometers_driven"`
}
