package main

type ParkingService interface {
	CreateParkingLot(slot int64)
	ParkACar(carNumber string, color string) int64
	LeaveParkingLot(int64)
	GetStatus()
	GetRegistrationNoOfCarsWithColor(color string) []string
	GetSlotNoWithRegistrationNo(regNo string) string
	GetSlotNumbersFromColor(color string) []int64
}
