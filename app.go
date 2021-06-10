package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mainMux := mux.NewRouter()
	
	//mainMux.Use(Middleware) // this will contain email or auth check

	rentalMux := mainMux.PathPrefix("/rental-vehicle").Subrouter()
	rentalMux.HandleFunc("/rates", GetRentalVehicleRates()).Methods("GET")
	rentalMux.HandleFunc("/available",GetAvailableVehicles()).Methods("GET")
	rentalMux.HandleFunc("/book",BookVehicle()).Methods("POST")


	log.Println("Server Run at: localhost:5000")
	http.ListenAndServe("localhost:5000", (mainMux))
}
