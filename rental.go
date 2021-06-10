package main

import (
	"net/http"
	"Practice/models/vehicle"
	"Practice/utils"
	"Practice/models/common"
	"Practice/models/request"
	"time"
)


// get all the vehicels rate
func GetRentalVehicleRates() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// dummy data add
		data := []vehicle.RatesToRentVehicles{
			vehicle.RatesToRentVehicles{
			Vehicle: "Maruti",
			Type: "Car",
			Amount: 100,
			},
		}
		
		utils.RespondWithJSON(w, http.StatusOK,common.CommonResponse{
			true,
			"Success",
			http.StatusOK,
			data,
		})
	}
}

// get all available vehicles
func GetAvailableVehicles() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		
		data := []vehicle.Vehicle{
			vehicle.Vehicle {
			Id: 1,	
			Vehicle: "Maruti",
			Type: "Car",
			IsAvailable: true,
			TimeFrame: []time.Time{time.Now()}, // set time now availble for sample
			},
		}

		utils.RespondWithJSON(w, http.StatusOK,common.CommonResponse{
			true,
			"Success",
			http.StatusOK,
			data,
		})
	}
}

// book a vehicle
func BookVehicle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := &request.BookVehicleRequest{}
		utils.ParseBody(r,body)
		var data []vehicle.Vehicle {
			vehicle.Vehicle {
				Id: 1,	
				Vehicle: "Maruti",
				Type: "Car",
				IsAvailable: true,
				TimeFrame: []time.Time{time.Now()}, // set time now availble for sample
				},	
		}

		// check time frame
		// do calculation of amount to pay
		

	}
}