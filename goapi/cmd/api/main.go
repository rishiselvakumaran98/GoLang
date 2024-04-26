package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rishiselvakumaran98/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	// This will help to setup our router which will be used as the endpoint
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting Go Service from Cmd ....")

	err := http.ListenAndServe("localhost:3000", r) // pass in the handler that our Mux type satisfies
	if err != nil{
		log.Error(err)
	}
}