package vehicle

import (
	"time"
)


type RatesToRentVehicles struct {
	Vehicle string `json:"vehicle"`
	Type string `json:"type"`
	Amount int `json:"amount"` // as per hr bases 
}


type Vehicle struct {
	Id int `json:"id"`
	Vehicle string   `json:"vehicle"`
	Type string `json:"type"`
	IsAvailable bool `json:"isAvailable"`
	TimeFrame []time.Time   `json:"timeFrame"` // multiple slots will contain
}