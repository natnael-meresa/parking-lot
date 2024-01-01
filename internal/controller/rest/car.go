package rest

import (
	"encoding/json"
	"net/http"
	"parking-lot/internal/controller/dto"
	"parking-lot/internal/usecase"

	"parking-lot/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type carHandler struct {
	carUseCase usecase.CarUseCase
	log        *logger.Logger
}

func newCarRoutes(handler *mux.Router, u usecase.CarUseCase, log *logger.Logger) {
	r := &carHandler{
		carUseCase: u,
		log:        log,
	}

	{
		// handle user endpoint
		handler.HandleFunc("/cars", r.listCarsHandler).Methods(http.MethodGet)
		handler.HandleFunc("/cars", r.addCarHandler).Methods(http.MethodPost)
		handler.HandleFunc("/cars/{registration}/rentals", r.rentCarHandler).Methods(http.MethodPost)
		handler.HandleFunc("/cars/{registration}/returns", r.returnCarHandler).Methods(http.MethodPost)
	}
}

func (ch *carHandler) listCarsHandler(w http.ResponseWriter, r *http.Request) {
	cars, err := ch.carUseCase.ListCars(r.Context())
	if err != nil {
		ch.handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(cars)
}

func (ch *carHandler) addCarHandler(w http.ResponseWriter, r *http.Request) {
	var newCar dto.Car
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	if err := validator.New().Struct(newCar); err != nil {
		ch.handleValidationError(w, err)
		return
	}

	id, err := ch.carUseCase.CreateCar(r.Context(), newCar)
	if err != nil {
		ch.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (ch *carHandler) rentCarHandler(w http.ResponseWriter, r *http.Request) {
	registration := mux.Vars(r)["registration"] // Assuming registration is a path parameter
	err := ch.carUseCase.RentCar(r.Context(), registration)
	if err != nil {
		ch.handleError(w, err)
		return
	}

	w.Write([]byte("Car rented successfully"))
}

func (ch *carHandler) returnCarHandler(w http.ResponseWriter, r *http.Request) {
	var returnRequest dto.ReturnRequest
	if err := json.NewDecoder(r.Body).Decode(&returnRequest); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	registration := mux.Vars(r)["registration"] // Assuming registration is a path parameter
	err := ch.carUseCase.ReturnCar(r.Context(), registration, float64(*returnRequest.KilometersDriven))
	if err != nil {
		ch.handleError(w, err)
		return
	}

	w.Write([]byte("Car returned successfully"))
}

// func (ch *carHandler) handleValidationError(w http.ResponseWriter, err error) {
//     // Extract validation errors and format response
//     // ...
// }
