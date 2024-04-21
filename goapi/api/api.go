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

