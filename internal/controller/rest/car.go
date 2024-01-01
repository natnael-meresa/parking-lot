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
		handler.HandleFunc("", r.listCarsHandler).Methods(http.MethodGet)
		handler.HandleFunc("", r.addCarHandler).Methods(http.MethodPost)
		handler.HandleFunc("/{registration}/rentals", r.rentCarHandler).Methods(http.MethodPost)
		handler.HandleFunc("/{registration}/returns", r.returnCarHandler).Methods(http.MethodPost)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"message": "Car rented successfully"})
}

func (ch *carHandler) returnCarHandler(w http.ResponseWriter, r *http.Request) {
	var returnRequest dto.ReturnRequest
	if err := json.NewDecoder(r.Body).Decode(&returnRequest); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	registration := mux.Vars(r)["registration"] // Assuming registration is a path parameter
	err := ch.carUseCase.ReturnCar(r.Context(), registration, float64(*returnRequest.KilometersDriven))
	if err != nil {
		ch.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"message": "Car returned successfully"})
}
