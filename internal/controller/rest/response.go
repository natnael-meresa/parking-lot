package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (ch *carHandler) handleError(w http.ResponseWriter, err error) {
	// Log the error
	log.Println(err)

	// Determine appropriate HTTP status code based on error type
	var statusCode int
	switch {
	case errors.Is(err, car.ErrNotFound):
		statusCode = http.StatusNotFound
	case errors.Is(err, car.ErrAlreadyRented):
		statusCode = http.StatusConflict
	case errors.Is(err, car.ErrInvalidMileage):
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	// Write error response with appropriate message
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
