package model

// Car is the model for a car.
type Car struct {
	// ID is the unique identifier of the car.
	ID           int    `json:"id" db:"id"`
	// Model is the model of the car.
	Model        string `json:"model" db:"model" validate:"required"`
	// Registration is the registration number of the car.
	Registration string `json:"registration" db:"registration" validate:"required,unique"`
	// Mileage is the mileage of the car.
	Mileage      int    `json:"mileage" db:"mileage"`
	// Available is the status of the car.
	Available    string `json:"available" db:"available" validate:"required,oneof=available rented"`
}
