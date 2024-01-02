package dto

// ReturnRequest is the data transfer object for a car.
type ReturnRequest struct {
	// ID is the unique identifier of the car.
	KilometersDriven *int `json:"kilometers_driven"`
}
