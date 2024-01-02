package rest

import (
	"parking-lot/internal/usecase"
	"parking-lot/pkg/logger"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router.
func NewRouter(app *mux.Router,
	log *logger.Logger,
	carUseCase usecase.CarUseCase) {

	// handle car endpoint
	carRouter := app.PathPrefix("/cars").Subrouter()

	{
		newCarRoutes(carRouter, carUseCase, log)
	}
}
