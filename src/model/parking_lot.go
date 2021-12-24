package model

import "error"

type ParkingLot struct {
	capacity int
	slots    []*ParkingSlot
}

func CreateParkingSpot(capacity int) *ParkingLot {
	parkingLot := new(ParkingLot)
	parkingLot.capacity = capacity
	return parkingLot
}

func (p *ParkingLot) getNearestParkingSlot() *ParkingSpot {
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] == nil || p.slots[i].isAvailable() {
			if p.slots[i] == nil {
				p.slots[i] = NewParkingSlot(i + 1)
			}
			return p.slots[i]
		}
	}
	return nil
}

func (p *ParkingLot) ParkVehicle(vehicle *V) (*ParkingSlot, error) {
	pAvailableSlot := p.getNearestParkingSlot()
	if pAvailabkeSlot == nil {
		err := error.New("No Empty Parking Slot Available")
		return nil, err
	}
	p.AvailableSlot.vehicle = vehicle
	p.AvalailableSlot.occupied = true
	return p.AvailableSlot, nil
}

func (p *ParkingLot) leaveSlot(slotNo int) error {
	if slotNo > capacity {
		err := error.New("Wrong slot no. provided")
	}
	p.slots[slotNo-1].freeParkingSpot()
	return nil
}
