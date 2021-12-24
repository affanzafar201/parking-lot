package model

import (
	"errors"
	"fmt"
	"strings"
)

type ParkingLot struct {
	capacity int
	slots    []*ParkingSlot
}

func CreateParkingLot(capacity int) *ParkingLot {
	parkingLot := new(ParkingLot)
	parkingLot.capacity = capacity
	parkingLot.slots = make([]*ParkingSlot, capacity)
	return parkingLot
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
	} else if p.slots[slotNo-1].IsAvailable() {
		err := errors.New("Slot already empty")
		return err
	} else {
		p.slots[slotNo-1].FreeParkingSlot()
		return nil
	}
}

func (p *ParkingLot) GetStatus() {
	var list = []string{fmt.Sprintf("%-12s%-20s%-10s", "Slot No.", "Registration No", "Colour")}
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && p.slots[i].IsAvailable() == false {
			list = append(list, fmt.Sprintf("%-12v%-20v%-10v", p.slots[i].GetSlotNo(), p.slots[i].GetVehicle().GetRegistrationNo(), p.slots[i].GetVehicle().GetColor()))
		}
	}
	output := strings.Join(list, "\n")
	fmt.Println(output)
}

func (p *ParkingLot) GetRegistrationNoForColor(color string) {
	ans := make([]string, 0, p.capacity)
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && p.slots[i].IsAvailable() == false {
			if p.slots[i].vehicle.color == color {
				ans = append(ans, p.slots[i].vehicle.GetRegistrationNo())
			}
		}
	}
	fmt.Printf("%+v \n", ans)
}

func (p *ParkingLot) GetSlotNoFromColor(color string) {
	ans := make([]string, 0, p.capacity)
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && p.slots[i].IsAvailable() == false {
			if p.slots[i].GetVehicle().GetColor() == color {
				ans = append(ans, string(p.slots[i].GetSlotNo()))
			}
		}
	}
	fmt.Printf("%+v \n", ans)
}

func (p *ParkingLot) GetSlotNoFromRegistrationNo(regNo string) {
	ans := make([]string, 0, p.capacity)
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != nil && p.slots[i].IsAvailable() == false {
			if p.slots[i].GetVehicle().GetRegistrationNo() == regNo {
				ans = append(ans, p.slots[i].GetVehicle().GetRegistrationNo())
			}
		}
	}
	fmt.Printf("%+v \n", ans)
}
