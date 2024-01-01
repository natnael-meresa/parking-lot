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

type CarUseCase interface {
	CreateCar(ctx context.Context, carDto dto.Car) (int, error)
	ListCars(ctx context.Context) ([]dto.Car, error)
	RentCar(ctx context.Context, registration string) error
	ReturnCar(ctx context.Context, registration string, kilometers_driven float64) error
}
type carUseCase struct {
	carRepo    repository.CarRepository
	rentalRepo repository.RentalRepository
	log        *logger.Logger
}

func NewCar(r repository.CarRepository, rentalRepo repository.RentalRepository, log *logger.Logger) CarUseCase {
	return &carUseCase{
		carRepo:    r,
		rentalRepo: rentalRepo,
		log:        log,
	}
}

func (c *carUseCase) CreateCar(ctx context.Context, carDto dto.Car) (int, error) {

	_, err := c.carRepo.GetCarByRegistration(ctx, carDto.Registration)
	if err == nil {
		return 0, constant.ErrDuplicateRegistration
	}

	id, err := c.carRepo.CreateCar(ctx, model.Car{
		Registration: carDto.Registration,
		Model:        carDto.Model,
		Mileage:      carDto.Mileage,
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *carUseCase) ListCars(ctx context.Context) ([]dto.Car, error) {
	cars, err := c.carRepo.ListCars(ctx)
	if err != nil {
		return nil, err
	}

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

	return carDtos, nil
}

func (c *carUseCase) RentCar(ctx context.Context, registration string) error {
	car, err := c.carRepo.GetCarByRegistration(ctx, registration)
	if err != nil {
		if err == sql.ErrNoRows {
			return constant.ErrNotFound
		}
		return err
	}

	if car.Available != constant.Available {
		return constant.ErrAlreadyRented
	}

	_, err = c.rentalRepo.CreateRental(ctx, model.Rental{
		StartDate: time.Now(),
		CarID:     car.ID,
	})

	if err != nil {
		return err
	}

	err = c.carRepo.UpdateCarStatus(ctx, registration, constant.Unavailable)
	if err != nil {
		return err
	}

	return nil
}

func (c *carUseCase) ReturnCar(ctx context.Context, registration string, kilometers_driven float64) error {
	car, err := c.carRepo.GetCarByRegistration(ctx, registration)
	if err != nil {
		if err == sql.ErrNoRows {
			return constant.ErrNotFound
		}
		return err
	}

	if car.Available == constant.Available {
		return constant.ErrNotRented
	}

	_, err = c.rentalRepo.GetRentalByCarID(ctx, car.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return constant.ErrNotFound
		}

		return err
	}

	err = c.rentalRepo.UpdateRentalEndDate(ctx, car.ID, time.Now())
	if err != nil {
		return err
	}

	err = c.rentalRepo.UpdateRentalKilometersDriven(ctx, car.ID, int(kilometers_driven))
	if err != nil {
		return err
	}

	err = c.carRepo.UpdateCarStatus(ctx, registration, constant.Available)
	if err != nil {
		return err
	}

	err = c.carRepo.UpdateCarMileage(ctx, registration, car.Mileage+int(kilometers_driven))
	if err != nil {
		return err
	}

	return nil
}
