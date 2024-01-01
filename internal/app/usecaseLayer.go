package app

import (
	"parking-lot/internal/usecase"
	"parking-lot/pkg/logger"
)

type UseCaseLayer struct {
	UserUseCase usecase.CarUseCase
}

func NewUsecaseLayer(repoLayer RepoLayer, log *logger.Logger) UseCaseLayer {
	return UseCaseLayer{
		UserUseCase: usecase.NewCar(repoLayer.CarRepo, repoLayer.RentalRepo, log),
	}
}
