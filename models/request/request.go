package request

type BookVehicleRequest struct {
	Id int `json:"id"` // vehicle id 
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
}
