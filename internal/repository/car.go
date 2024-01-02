package repository

import (
	"context"
	"database/sql"
	"parking-lot/internal/model"
	"parking-lot/pkg/logger"
)

// CarRepository is the interface for the car repository.
type CarRepository interface {
	// CreateCar creates a new car.
	CreateCar(ctx context.Context, carModel model.Car) (int, error)
	// ListCars lists all cars.
	ListCars(ctx context.Context) ([]model.Car, error)
	// GetCarByRegistration gets a car by registration.
	GetCarByRegistration(ctx context.Context, registration string) (*model.Car, error)
	// UpdateCarStatus updates a car status.
	UpdateCarStatus(ctx context.Context, registration string, status string) error
	// UpdateCarMileage updates a car mileage.
	UpdateCarMileage(ctx context.Context, registration string, mileage int) error
}

// carRepository is the repository for cars.
type carRepository struct {
	db  *sql.DB
	log *logger.Logger
}

// NewCar creates a new car repository.
func NewCar(db *sql.DB, log *logger.Logger) CarRepository {
	return &carRepository{
		db:  db,
		log: log,
	}
}

// CreateCar creates a new car by executing a prepared statement.
func (c *carRepository) CreateCar(ctx context.Context, carModel model.Car) (int, error) {
	// Prepare statement for inserting data
	crt, err := c.db.Prepare("INSERT INTO cars (registration, model, mileage) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	// Execute statement
	res, err := crt.Exec(carModel.Registration, carModel.Model, carModel.Mileage)
	if err != nil {
		return 0, err
	}

	// Get the last inserted row id
	rowID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Return the inserted row id
	return int(rowID), nil
}


// ListCars lists all cars.
func (c *carRepository) ListCars(ctx context.Context) ([]model.Car, error) {
	
	var cars []model.Car
	var car model.Car

	// Execute the query
	rows, err := c.db.QueryContext(ctx, "SELECT * FROM cars")
	if err != nil {
		return nil, err
	}

	// Iterate over the rows
	for rows.Next() {
		err = rows.Scan(&car.ID, &car.Registration, &car.Model, &car.Mileage, &car.Available)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

// GetCarByRegistration gets a car by registration.
func (c *carRepository) GetCarByRegistration(ctx context.Context, registration string) (*model.Car, error) {
	var car model.Car

	// Execute the query
	err := c.db.QueryRowContext(ctx, "SELECT * FROM cars WHERE registration = ?", registration).Scan(&car.ID, &car.Registration, &car.Model, &car.Mileage, &car.Available)
	if err != nil {
		return nil, err
	}

	return &car, nil
}

// UpdateCarStatus updates a car status.
func (c *carRepository) UpdateCarStatus(ctx context.Context, registration string, status string) error {
	// Execute the query
	_, err := c.db.ExecContext(ctx, "UPDATE cars SET available = ? WHERE registration = ?", status, registration)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCarMileage updates a car mileage.
func (c *carRepository) UpdateCarMileage(ctx context.Context, registration string, mileage int) error {
	// Execute the query
	_, err := c.db.ExecContext(ctx, "UPDATE cars SET mileage = ? WHERE registration = ?", mileage, registration)
	if err != nil {
		return err
	}

	return nil
}
