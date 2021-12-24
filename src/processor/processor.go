package processor

import (
	"constants"
	"fmt"
	"model"
	"strconv"
	"strings"
)

type Processor struct {
	parkingService *model.ParkingLot
}

func (p *Processor) setParkingService(parkingService *model.ParkingLot) {
	p.parkingService = parkingService
}

func (p *Processor) Execute(input string) {
	inputs := strings.Split(input, " ")

	switch command := inputs[0]; command {

	case constants.CreateParkingLot:
		capacity, err := strconv.Atoi(inputs[1])
		if err != nil {
			fmt.Println(fmt.Errorf("Not a valid value in capacity and failed with error: %v", err))
		} else {
			p.setParkingService(model.CreateParkingLot(capacity))
		}
	case constants.Park:
		newSlot, err := p.parkingService.ParkVehicle(model.NewVehicle(inputs[1], inputs[2]))
		if err != nil {
			fmt.Println(fmt.Errorf("Park car failed with error: %v", err))
		}
		fmt.Println("Allocated slot no.", newSlot.GetSlotNo())
	case constants.Leave:
		slotNo, err := strconv.Atoi(inputs[1])
		if err != nil {
			fmt.Println(fmt.Errorf("Not a valid slot_number with error: %v", err))
		} else {
			p.parkingService.UnparkVehicle(slotNo)
			fmt.Printf("SlotNo %d is free\n", slotNo)
		}
	case constants.Status:
		p.parkingService.GetStatus()
	case constants.RegistrationNoForCarsWithColor:
		p.parkingService.GetRegistrationNoForColor(inputs[1])
	case constants.SlotNoForCarsWithColor:
		p.parkingService.GetSlotNoFromColor(inputs[1])
	case constants.SlotNoForRegistrationNo:
		p.parkingService.GetSlotNoFromRegistrationNo(inputs[1])

	}
}
