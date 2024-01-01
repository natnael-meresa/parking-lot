package dto

type Car struct {
	ID           int    `json:"id"`
	Model        string `json:"model"`
	Registration string `json:"registration"`
	Mileage      int    `json:"mileage"`
	Available    string `json:"available"` // Available or Rented
}
