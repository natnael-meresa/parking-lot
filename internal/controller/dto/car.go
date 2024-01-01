package dto

type Car struct {
	ID           int    `json:"id,omitempty"`
	Model        string `json:"model" validate:"required"`
	Registration string `json:"registration" validate:"required"`
	Mileage      int    `json:"mileage"`
	Available    string `json:"available,omitempty"`
}
