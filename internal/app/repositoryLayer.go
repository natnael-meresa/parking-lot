package app

import (
	"database/sql"
	"parking-lot/internal/repository"
	"parking-lot/pkg/logger"
)

type RepoLayer struct {
	CarRepo repository.CarRepository
	RentalRepo repository.RentalRepository
}

func NewRepoLayer(db *sql.DB, log *logger.Logger) RepoLayer {
	return RepoLayer{
		CarRepo: repository.NewCar(db, log),
		RentalRepo: repository.NewRental(db, log),
	}
}
