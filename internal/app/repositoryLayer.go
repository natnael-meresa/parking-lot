package app

import (
	"database/sql"
	"parking-lot/internal/repository"
	"parking-lot/pkg/logger"
)

// RepoLayer is the repository layer.
type RepoLayer struct {
	CarRepo repository.CarRepository
	RentalRepo repository.RentalRepository
}

// NewRepoLayer creates a new repository layer.
func NewRepoLayer(db *sql.DB, log *logger.Logger) RepoLayer {
	return RepoLayer{
		CarRepo: repository.NewCar(db, log),
		RentalRepo: repository.NewRental(db, log),
	}
}
