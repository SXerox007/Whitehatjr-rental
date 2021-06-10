debug:
	go run app.go rental.go


	curl -X GET localhost:5000/rental-vehicle/rates -v
	curl -X GET localhost:5000/rental-vehicle/available -v