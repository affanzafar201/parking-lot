package model

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

func (p *ParkingSlot) isAvailable() bool {
	return p.occupied
}

func (p *ParkingSlot) freeParkingSpot() {
	p.occupied = false
	p.vehicle = nil
}

func (p *ParkingSlot) allotCarToSlot(v *Vehicle) {
	p.occupied = true
	p.vehicle = v
}
