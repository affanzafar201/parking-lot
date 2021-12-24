package model

type ParkingSlot struct {
	occupied bool
	vehicle  *Vehicle
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
