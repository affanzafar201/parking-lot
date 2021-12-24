package model

import (
	"errors"
)

type ParkingLot struct {
	capacity int
	slots    []*ParkingSlot
}

func NewParkingLot(capacity int) *ParkingLot {
	parkingLot := new(ParkingLot)
	parkingLot.capacity = capacity
	parkingLot.slots = make([]*ParkingSlot, capacity)
	return parkingLot
}

func (p *ParkingLot) GetCapacity() int {
	return p.capacity
}

func (p *ParkingLot) getNearestParkingSlot() *ParkingSlot {
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] == nil || p.slots[i].IsAvailable() {
			if p.slots[i] == nil {
				p.slots[i] = NewParkingSlot(i + 1)
			}
			return p.slots[i]
		}
	}
	return nil
}

func (p *ParkingLot) ParkVehicle(vehicle *Vehicle) (*ParkingSlot, error) {
	pAvailableSlot := p.getNearestParkingSlot()
	if pAvailableSlot == nil {
		err := errors.New("No Empty Parking Slot Available")
		return nil, err
	}
	err := pAvailableSlot.AllotVehicle(vehicle)
	if err != nil {
		return nil, err
	}
	return pAvailableSlot, nil
}

func (p *ParkingLot) UnparkVehicle(slotNo int) error {
	if slotNo > p.capacity {
		err := errors.New("Wrong slot no. provided")
		return err
	} else if p.slots[slotNo-1] == nil || p.slots[slotNo-1].IsAvailable() {
		err := errors.New("Slot already empty")
		return err
	} else {
		p.slots[slotNo-1].FreeParkingSlot()
		return nil
	}
}

func (p *ParkingLot) GetFilledSlots() []*ParkingSlot {
	list := make([]*ParkingSlot, 0)
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && p.slots[i].IsAvailable() == false {
			list = append(list, p.slots[i])
		}
	}
	return list
}

func (p *ParkingLot) GetSlotsByVehicleColor(color string) []*ParkingSlot {
	ans := make([]*ParkingSlot, 0, p.capacity)
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && p.slots[i].IsAvailable() == false {
			if p.slots[i].GetVehicle().GetColor() == color {
				ans = append(ans, p.slots[i])
			}
		}
	}
	return ans
}

func (p *ParkingLot) GetSlotByVehicleRegistrationNo(regNo string) (*ParkingSlot, error) {
	var ans *ParkingSlot
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && p.slots[i].IsAvailable() == false {
			if p.slots[i].GetVehicle().GetRegistrationNo() == regNo {
				ans := p.slots[i]
				return ans, nil
			}
		}
	}
	if ans == nil {
		return ans, errors.New("Not Found")
	}
	return ans, nil
}
