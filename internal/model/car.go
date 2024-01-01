package model

type Car struct {
	ID           int    `json:"id" db:"id"`
	Model        string `json:"model" db:"model" validate:"required"`
	Registration string `json:"registration" db:"registration" validate:"required,unique"`
	Mileage      int    `json:"mileage" db:"mileage"`
	Available    string `json:"available" db:"available" validate:"required,oneof=available rented"`
}
