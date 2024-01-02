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
		// GET /cars
		handler.HandleFunc("", r.listCarsHandler).Methods(http.MethodGet)
		// POST /cars
		handler.HandleFunc("", r.addCarHandler).Methods(http.MethodPost)
		// POST /cars/{registration}/rentals
		handler.HandleFunc("/{registration}/rentals", r.rentCarHandler).Methods(http.MethodPost)
		// POST /cars/{registration}/returns
		handler.HandleFunc("/{registration}/returns", r.returnCarHandler).Methods(http.MethodPost)
	}
}

// listCarsHandler handles GET /cars endpoint
func (ch *carHandler) listCarsHandler(w http.ResponseWriter, r *http.Request) {
	// get all cars
	cars, err := ch.carUseCase.ListCars(r.Context())
	if err != nil {
		ch.handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(cars)
}

// addCarHandler handles POST /cars endpoint
func (ch *carHandler) addCarHandler(w http.ResponseWriter, r *http.Request) {
	var newCar dto.Car
	// decode request body into newCar
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	// validate newCar
	if err := validator.New().Struct(newCar); err != nil {
		ch.handleValidationError(w, err)
		return
	}

	// create new car
	id, err := ch.carUseCase.CreateCar(r.Context(), newCar)
	if err != nil {
		ch.handleError(w, err)
		return
	}

	// return id of newly created car
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// rentCarHandler handles POST /cars/{registration}/rentals endpoint
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

// returnCarHandler handles POST /cars/{registration}/returns endpoint
func (ch *carHandler) returnCarHandler(w http.ResponseWriter, r *http.Request) {
	var returnRequest dto.ReturnRequest
	// decode request body into returnRequest
	if err := json.NewDecoder(r.Body).Decode(&returnRequest); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	// validate returnRequest
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
