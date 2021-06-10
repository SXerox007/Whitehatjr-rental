package common

type CommonResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
}