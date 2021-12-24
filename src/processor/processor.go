package processor

import (
	"constants"
	"fmt"
	"model"
	"strconv"
	"strings"
)

type Processor struct {
	parkingLot *model.ParkingLot
}

func (p *Processor) setParkingLot(parkingLot *model.ParkingLot) {
	p.parkingLot = parkingLot
}

func (p *Processor) Execute(input string) {
	inputs := strings.Split(input, " ")

	switch command := inputs[0]; command {

	case constants.CreateParkingLot:
		if len(inputs) == 2 {
			capacity, err := strconv.Atoi(inputs[1])
			if err != nil {
				fmt.Println(fmt.Errorf("Not a valid value in capacity"))
			} else {
				p.setParkingLot(p.CreateParkingLot(capacity))
			}
		} else {
			if len(inputs) == 1 {
				fmt.Println("Capacity of parking lot not provided")
			} else {
				fmt.Println("More than one capacity provided")
			}
			fmt.Println("Command : create_parking_lot <capacity>")
		}

	case constants.Park:
		if len(inputs) == 3 {
			newSlot, err := p.parkingLot.ParkVehicle(model.NewVehicle(inputs[1], inputs[2]))
			if err != nil {
				fmt.Println(fmt.Errorf("Sorry, parking lot is full"))
			} else {
				fmt.Println("Allocated slot number:", newSlot.GetSlotNo())
			}
		} else {
			if len(inputs) == 1 {
				fmt.Println("Vehicle registration no not provided")
			} else if len(inputs) == 2 {
				fmt.Println("Vehicle color not provided")
			}
			fmt.Println("Command : park <reg no> <color>")
		}
	case constants.Leave:
		if len(inputs) == 2 {
			slotNo, err := strconv.Atoi(inputs[1])
			if err != nil {
				fmt.Println(fmt.Errorf("Not a valid slot_number with error: %v", err))
			} else {
				err := p.parkingLot.UnparkVehicle(slotNo)
				if err != nil {
					fmt.Println(fmt.Errorf("Leave car from slot failed with err: %v", err))
				} else {
					fmt.Printf("Slot number %d is free\n", slotNo)
				}
			}
		} else {
			if len(inputs) == 1 {
				fmt.Println("Slot No for leaving parking lot not provided")
			} else {
				fmt.Println("More than one slot no. provided")
			}
			fmt.Println("Command : leave <slot no>")
		}
	case constants.Status:
		p.GetStatus()
	case constants.RegistrationNoForCarsWithColor:
		p.GetRegistrationNoForColor(inputs[1])
	case constants.SlotNoForCarsWithColor:
		p.GetSlotNoFromColor(inputs[1])
	case constants.SlotNoForRegistrationNo:
		p.GetSlotNoFromRegistrationNo(inputs[1])
	default:
		fmt.Printf("Command %s not found\n", command)
		fmt.Println("Command Supported")
		fmt.Println("\tcreate_parking_lot <capacity>")
		fmt.Println("\tpark <reg no> <color>")
		fmt.Println("\tleave <slot no>")
		fmt.Println("\tstatus")
		fmt.Println("\tregistration_numbers_for_cars_with_colour <color>")
		fmt.Println("\tslot_numbers_for_cars_with_colour <color>")
		fmt.Println("\tslot_number_for_registration_number <reg no>")
	}
}

func (p *Processor) CreateParkingLot(capacity int) *model.ParkingLot {
	newParkingLot := model.NewParkingLot(capacity)
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
	return newParkingLot
}

func (p *Processor) GetStatus() {
	if p.parkingLot == nil {
		fmt.Println("Parking Lot Not created\nRun: \tcreate_parking_lot <capacity>")
		return
	}
	listOccupiedSlots := p.parkingLot.GetFilledSlots()
	var list = []string{fmt.Sprintf("%-8s    %-15s    %s", "Slot No.", "Registration No", "Colour")}
	for _, occupiedSlot := range listOccupiedSlots {
		list = append(list, fmt.Sprintf("%-8v    %-15v    %v", occupiedSlot.GetSlotNo(), occupiedSlot.GetVehicle().GetRegistrationNo(), occupiedSlot.GetVehicle().GetColor()))
	}
	output := strings.Join(list, "\n")
	fmt.Println(output)
}

func (p *Processor) GetRegistrationNoForColor(color string) {
	if p.parkingLot == nil {
		fmt.Println("Parking Lot Not created\nRun: \tcreate_parking_lot <capacity>")
		return
	}
	listOccupiedSlots := p.parkingLot.GetSlotsByVehicleColor(color)
	list := make([]string, 0, p.parkingLot.GetCapacity())
	for _, occupiedSlot := range listOccupiedSlots {
		list = append(list, fmt.Sprintf(occupiedSlot.GetVehicle().GetRegistrationNo()))
	}
	output := strings.Join(list, ", ")
	fmt.Println(output)
}

func (p *Processor) GetSlotNoFromColor(color string) {
	if p.parkingLot == nil {
		fmt.Println("Parking Lot Not created\nRun: \tcreate_parking_lot <capacity>")
		return
	}
	listOccupiedSlots := p.parkingLot.GetSlotsByVehicleColor(color)
	list := make([]string, 0, p.parkingLot.GetCapacity())
	for _, occupiedSlot := range listOccupiedSlots {
		list = append(list, fmt.Sprintf("%d", occupiedSlot.GetSlotNo()))
	}
	output := strings.Join(list, ", ")
	fmt.Println(output)
}

func (p *Processor) GetSlotNoFromRegistrationNo(regNo string) {
	if p.parkingLot == nil {
		fmt.Println("Parking Lot Not created\nRun: \tcreate_parking_lot <capacity>")
		return
	}
	occupiedSlot, err := p.parkingLot.GetSlotByVehicleRegistrationNo(regNo)
	if err != nil {
		fmt.Println("Not found")
		return
	}
	fmt.Println(occupiedSlot.GetSlotNo())
}
