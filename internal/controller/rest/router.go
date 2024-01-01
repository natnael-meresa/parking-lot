package rest

import (
	"parking-lot/internal/usecase"
	"parking-lot/pkg/logger"

	"github.com/gorilla/mux"
)

func NewRouter(app *mux.Router,
	log *logger.Logger,
	carUseCase usecase.CarUseCase) {

	carRouter := app.PathPrefix("/cars").Subrouter()

	{
		newCarRoutes(carRouter, carUseCase, log)
	}
}
