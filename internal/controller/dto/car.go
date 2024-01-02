package dto

// Car is the data transfer object for a car.
type Car struct {
	// ID is the unique identifier of the car.
	ID           int    `json:"id,omitempty"`
	// Model is the model of the car.
	Model        string `json:"model" validate:"required"`
	// Registration is the registration number of the car.
	Registration string `json:"registration" validate:"required"`
	// Mileage is the mileage of the car.
	Mileage      int    `json:"mileage"`
	// Available is the status of the car.
	Available    string `json:"available,omitempty"`
	// RentedBy is the user that rented the car.
}
