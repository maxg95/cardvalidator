package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/joeljunstrom/go-luhn"
)

type validationResponse struct {
	Valid bool   `json:"valid"`
	Error *error `json:"error,omitempty"`
}

type error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (app *application) validatorHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters.
	cardNumber := r.URL.Query().Get("card_number")
	expirationMonth := r.URL.Query().Get("expiration_month")
	expirationYear := r.URL.Query().Get("expiration_year")

	// Validate card number using Luhn algorithm.
	if !luhn.Valid(cardNumber) {
		sendJSONResponse(w, http.StatusOK, validationResponse{
			Valid: false,
			Error: &error{
				Code:    "001",
				Message: "Invalid card number",
			},
		})
		return
	}

	// Validate expiration date.
	if !isValidExpiration(expirationMonth, expirationYear) {
		sendJSONResponse(w, http.StatusOK, validationResponse{
			Valid: false,
			Error: &error{
				Code:    "002",
				Message: "Invalid expiration date",
			},
		})
		return
	}

	sendJSONResponse(w, http.StatusOK, validationResponse{
		Valid: true,
	})
}

func isValidExpiration(expirationMonth, expirationYear string) bool {
	// Validate expiration month and year.
	expiryMonth, err := strconv.Atoi(expirationMonth)
	if err != nil || expiryMonth < 1 || expiryMonth > 12 {
		return false
	}

	currentYear, currentMonth, _ := time.Now().Date()
	expiryYear, err := strconv.Atoi(expirationYear)
	if err != nil || expiryYear < currentYear || (expiryYear == currentYear && time.Month(expiryMonth) < currentMonth) {
		return false
	}

	return true
}

func sendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
