package model

type ParkingLot struct {
	capacity int64
	slots    []*ParkingSlot
}

func CreateParkingSpot(capacity int64) *ParkingLot {
	parkingLot := new(ParkingLot)
	parkingLot.capacity = capacity
	return parkingLot
}

func (p *ParkingLot) parkVehicle(vehicle *V) ParkingSlot {

}

func (p *ParkingLot) unParkVehicle(vehicle *V) {
}
