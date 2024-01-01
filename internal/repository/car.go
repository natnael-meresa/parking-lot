package repository

import (
	"context"
	"database/sql"
	"parking-lot/internal/model"
	"parking-lot/pkg/logger"
)

type CarRepository interface {
	CreateCar(ctx context.Context, carModel model.Car) (int, error)
	ListCars(ctx context.Context) ([]model.Car, error)
	GetCarByRegistration(ctx context.Context, registration string) (*model.Car, error)
	UpdateCarStatus(ctx context.Context, registration string, status string) error
	UpdateCarMileage(ctx context.Context, registration string, mileage int) error	
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

func (c *carRepository) CreateCar(ctx context.Context, carModel model.Car) (int, error) {
	var id int

	err := c.db.QueryRowContext(ctx, "INSERT INTO cars(registration, model, mileage) VALUES($1, $2, $3) RETURNING id", carModel.Registration, carModel.Model, carModel.Mileage).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}


func (c *carRepository) ListCars(ctx context.Context) ([]model.Car, error) {
	var cars []model.Car
	var car model.Car

	rows, err := c.db.QueryContext(ctx, "SELECT * FROM cars")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&car.ID, &car.Registration, &car.Model, &car.Mileage, &car.Available)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (c *carRepository) GetCarByRegistration(ctx context.Context, registration string) (*model.Car, error) {
	var car model.Car

	err := c.db.QueryRowContext(ctx, "SELECT * FROM cars WHERE registration = $1", registration).Scan(&car.ID, &car.Registration, &car.Model, &car.Mileage, &car.Available)
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (c *carRepository) UpdateCarStatus(ctx context.Context, registration string, status string) error {
	_, err := c.db.ExecContext(ctx, "UPDATE cars SET available = $1 WHERE registration = $2", status, registration)
	if err != nil {
		return err
	}

	return nil
}

func (c *carRepository) UpdateCarMileage(ctx context.Context, registration string, mileage int) error {
	_, err := c.db.ExecContext(ctx, "UPDATE cars SET mileage = $1 WHERE registration = $2", mileage, registration)
	if err != nil {
		return err
	}

	return nil
}
