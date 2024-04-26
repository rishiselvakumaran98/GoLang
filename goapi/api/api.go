package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params represents the parameters our API endpoint will take in
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response outlines the successful response from the Server
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error code
	Code int

	// Error Message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// Construct wrappers that help us guide the errors for writeError
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// An internal error handler to handle error fro internal server errors
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An unexpected error happened from our side, Sorry!", http.StatusInternalServerError)
	}
)
