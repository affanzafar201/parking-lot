package model

import (
	"errors"
)

type ParkingSlot struct {
	occupied bool
	slotNo   int
	vehicle  *Vehicle
}

func NewParkingSlot(slotNo int) *ParkingSlot {
	p := new(ParkingSlot)
	p.occupied = false
	p.slotNo = slotNo
	p.vehicle = nil
	return p
}

func (p *ParkingSlot) GetVehicle() *Vehicle {
	return p.vehicle
}

func (p *ParkingSlot) GetSlotNo() int {
	return p.slotNo
}

func (p *ParkingSlot) IsAvailable() bool {
	return !p.occupied
}

func (p *ParkingSlot) FreeParkingSlot() {
	p.occupied = false
	p.vehicle = nil
}

func (p *ParkingSlot) AllotVehicle(v *Vehicle) error {
	if p.IsAvailable() {
		p.occupied = true
		p.vehicle = v
		return nil
	}
	return errors.New("Slot is not empty")
}
