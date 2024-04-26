package handlers

import (
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/rishiselvakumaran98/goapi/internal/middleware"
)

// Starting a function in Uppercase means the function is Public
func Handler(r *chi.Mux){
	// Global Middleware
	r.Use(chimiddle.StripSlashes) // A prebuilt function from chi to remove trailing slashes from the URL

	r.Route("/account", func(r chi.Router){ // will be used to define a group of routes
		// Middleware for the /account route
		r.Use(middleware.Authorization) // used to check if the user is authorized to access the route

		r.Get("/coins", GetCoinBalance) // GetCoinBalance is a function that will be called when the /account/coins route is hit
	})

}
