package app

import (
	"parking-lot/internal/usecase"
	"parking-lot/pkg/logger"
)

// UseCaseLayer is the usecase layer.
type UseCaseLayer struct {
	UserUseCase usecase.CarUseCase
}

// NewUsecaseLayer creates a new usecase layer.
func NewUsecaseLayer(repoLayer RepoLayer, log *logger.Logger) UseCaseLayer {
	return UseCaseLayer{
		UserUseCase: usecase.NewCar(repoLayer.CarRepo, repoLayer.RentalRepo, log),
	}
}
