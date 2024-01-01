package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"parking-lot/internal/constant"

	"github.com/go-playground/validator/v10"
)

func (ch *carHandler) handleError(w http.ResponseWriter, err error) {
	// Log the error
	log.Println(err)

	// Determine appropriate HTTP status code based on error type
	var statusCode int
	switch {
	case errors.Is(err, constant.ErrNotFound):
		statusCode = http.StatusNotFound
	case errors.Is(err, constant.ErrAlreadyRented):
		statusCode = http.StatusConflict
	case errors.Is(err, constant.ErrInvalidMileage):
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	// Write error response with appropriate message
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func (ch *carHandler) handleValidationError(w http.ResponseWriter, err error) {
	var validationErrors []string
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, fmt.Sprintf("%s: %s", err.Field(), err.Tag()))
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{"errors": validationErrors})
}
