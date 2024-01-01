package usecase

import (
	"context"
	"parking-lot/internal/controller/dto"
	"parking-lot/internal/repository"
	"parking-lot/pkg/logger"
)

type CarUseCase interface {
	CreateCar(ctx context.Context, carDto dto.Car) (int, error)
	ListCars(ctx context.Context) ([]dto.Car, error)
	RentCar(ctx context.Context, registration string) error
	ReturnCar(ctx context.Context, registration string, kilometers_driven float64) error
}
type carUseCase struct {
	carRepo repository.CarRepository
	log     *logger.Logger
}

func NewCar(r repository.CarRepository, log *logger.Logger) CarUseCase {
	return &carUseCase{
		carRepo: r,
		log:     log,
	}
}

func (c *carUseCase) CreateCar(ctx context.Context, carDto dto.Car) (*dto.Car, error) {
	return nil, nil
}
