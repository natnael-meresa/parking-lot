package repository

import (
	"context"
	"database/sql"
	"parking-lot/internal/model"
	"parking-lot/pkg/logger"
)

type CarRepository interface {
	CreateCar(ctx context.Context, carDto model.Car) (*model.Car, error)
}

type carRepository struct {
	db  *sql.DB
	log *logger.Logger
}

func NewCar(db *sql.DB, log *logger.Logger) CarRepository {
	return &carRepository{
		db:  db,
		log: log,
	}
}

func (c *carRepository) CreateCar(ctx context.Context, carDto model.Car) (*model.Car, error) {
	return nil, nil
}
