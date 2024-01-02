package usecase

import (
	"context"
	"database/sql"
	"parking-lot/internal/constant"
	"parking-lot/internal/controller/dto"
	"parking-lot/internal/model"
	"parking-lot/internal/repository"
	"parking-lot/pkg/logger"
	"time"
)

// CarUseCase is the car usecase interface.
type CarUseCase interface {
	// CreateCar creates a new car.
	CreateCar(ctx context.Context, carDto dto.Car) (int, error)
	// ListCars lists all cars.
	ListCars(ctx context.Context) ([]dto.Car, error)
	// RentCar rents a car.
	RentCar(ctx context.Context, registration string) error
	// ReturnCar returns a car.
	ReturnCar(ctx context.Context, registration string, kilometers_driven float64) error
}

// carUseCase is the usecase for cars.
type carUseCase struct {
	carRepo    repository.CarRepository
	rentalRepo repository.RentalRepository
	log        *logger.Logger
}

// NewCar creates a new car usecase.
func NewCar(r repository.CarRepository, rentalRepo repository.RentalRepository, log *logger.Logger) CarUseCase {
	return &carUseCase{
		carRepo:    r,
		rentalRepo: rentalRepo,
		log:        log,
	}
}

// CreateCar creates a new car.
func (c *carUseCase) CreateCar(ctx context.Context, carDto dto.Car) (int, error) {

	// Check if car with registration already exists
	_, err := c.carRepo.GetCarByRegistration(ctx, carDto.Registration)
	if err == nil {
		return 0, constant.ErrDuplicateRegistration
	}

	// Create car
	id, err := c.carRepo.CreateCar(ctx, model.Car{
		Registration: carDto.Registration,
		Model:        carDto.Model,
		Mileage:      carDto.Mileage,
	})

	if err != nil {
		return 0, err
	}

	// Return id of created car
	return id, nil
}

// ListCars lists all cars.
func (c *carUseCase) ListCars(ctx context.Context) ([]dto.Car, error) {
	// Get all cars
	cars, err := c.carRepo.ListCars(ctx)
	if err != nil {
		return nil, err
	}

	// Convert cars to carDtos
	var carDtos []dto.Car
	for _, car := range cars {
		carDtos = append(carDtos, dto.Car{
			ID:           car.ID,
			Registration: car.Registration,
			Model:        car.Model,
			Mileage:      car.Mileage,
			Available:    car.Available,
		})
	}

	// Return carDtos
	return carDtos, nil
}

// RentCar rents a car.
func (c *carUseCase) RentCar(ctx context.Context, registration string) error {
	// Get car by registration
	car, err := c.carRepo.GetCarByRegistration(ctx, registration)
	if err != nil {
		if err == sql.ErrNoRows {
			return constant.ErrNotFound
		}
		return err
	}

	// Check if car is available
	if car.Available != constant.Available {
		return constant.ErrAlreadyRented
	}

	// Create rental
	_, err = c.rentalRepo.CreateRental(ctx, model.Rental{
		StartDate: time.Now(),
		CarID:     car.ID,
	})

	if err != nil {
		return err
	}

	// Update car status
	err = c.carRepo.UpdateCarStatus(ctx, registration, constant.Unavailable)
	if err != nil {
		return err
	}

	return nil
}

// ReturnCar returns a car.
func (c *carUseCase) ReturnCar(ctx context.Context, registration string, kilometers_driven float64) error {
	// Get car by registration
	car, err := c.carRepo.GetCarByRegistration(ctx, registration)
	if err != nil {
		if err == sql.ErrNoRows {
			return constant.ErrNotFound
		}
		return err
	}

	// Check if car is available, if it is available it means it is not rented so we return an error
	if car.Available == constant.Available {
		return constant.ErrNotRented
	}

	// Get rental by car id
	_, err = c.rentalRepo.GetRentalByCarID(ctx, car.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return constant.ErrNotFound
		}

		return err
	}

	// Update rental end date
	err = c.rentalRepo.UpdateRentalEndDate(ctx, car.ID, time.Now())
	if err != nil {
		return err
	}

	// Update rental kilometers driven
	err = c.rentalRepo.UpdateRentalKilometersDriven(ctx, car.ID, int(kilometers_driven))
	if err != nil {
		return err
	}

	// Update car status
	err = c.carRepo.UpdateCarStatus(ctx, registration, constant.Available)
	if err != nil {
		return err
	}

	// Update car mileage
	err = c.carRepo.UpdateCarMileage(ctx, registration, car.Mileage+int(kilometers_driven))
	if err != nil {
		return err
	}

	return nil
}
