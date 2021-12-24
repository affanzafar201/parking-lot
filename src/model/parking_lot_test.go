package model

import (
	"errors"
	"reflect"
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	testParkingLot := CreateParkingLot(6)
	if testParkingLot == nil {
		t.Errorf("Failed to create parking lot")
	} else if testParkingLot.capacity != 6 {
		t.Errorf("Failed to create parking lot with want capacity %v, got %v", testParkingLot.capacity, 6)
	}
}

func TestParkVehicle(t *testing.T) {
	type parklot struct {
		capacity int
		slots    []*ParkingSlot
	}
	type testVehicle struct {
		car Vehicle
	}
	tests := []struct {
		name    string
		parklot parklot
		args    testVehicle
		want    *ParkingSlot
		wantErr bool
	}{
		{
			"TestCase 1: Parking slots is not full",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: false,
						slotNo:   1,
						vehicle:  nil,
					},
					{
						occupied: false,
						slotNo:   2,
						vehicle:  nil,
					},
					{
						occupied: false,
						slotNo:   3,
						vehicle:  nil,
					},
				},
			},
			testVehicle{car: Vehicle{registrationNo: "KA01-1498", color: "Red"}},
			&ParkingSlot{
				occupied: true,
				slotNo:   1,
				vehicle:  &Vehicle{registrationNo: "KA01-1498", color: "Red"},
			},
			false,
		},
		{
			"TestCase 2: Parking slots is full",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: true,
						slotNo:   1,
						vehicle:  &Vehicle{registrationNo: "BR01-1234", color: "White"},
					},
					{
						occupied: true,
						slotNo:   2,
						vehicle:  &Vehicle{registrationNo: "DL12-5608", color: "Blue"},
					},
					{
						occupied: true,
						slotNo:   3,
						vehicle:  &Vehicle{registrationNo: "KA01-1523", color: "Black"},
					},
				},
			},
			testVehicle{car: Vehicle{registrationNo: "KA01-1498", color: "Red"}},
			nil,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingLot := &ParkingLot{
				capacity: test.parklot.capacity,
				slots:    test.parklot.slots,
			}
			got, err := testParkingLot.ParkVehicle(&test.args.car)
			if (err != nil) != test.wantErr {
				t.Errorf("\x1b[31;1mParkingLot.ParkVehicle() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("\x1b[31;1mParking.ParkVehicle() = %v, want %v\x1b[0m", got, test.want)
			}
		})
	}
}

func TestUnparkVehicle(t *testing.T) {
	type parklot struct {
		capacity int
		slots    []*ParkingSlot
	}
	type testVehicle struct {
		car Vehicle
	}
	tests := []struct {
		name    string
		parklot parklot
		args    int
		err     error
		wantErr bool
	}{
		{
			"TestCase 1: Unpark Vehicle from parking slot where car is parked",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: true,
						slotNo:   1,
						vehicle:  &Vehicle{registrationNo: "BR01-1234", color: "White"},
					},
					{
						occupied: false,
						slotNo:   2,
						vehicle:  nil,
					},
					{
						occupied: false,
						slotNo:   3,
						vehicle:  nil,
					},
				},
			},
			1,
			nil,
			false,
		},
		{
			"TestCase 2: Unpark Vehicle from parking slot where no car is parked",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: true,
						slotNo:   1,
						vehicle:  &Vehicle{registrationNo: "BR01-1234", color: "White"},
					},
					{
						occupied: false,
						slotNo:   2,
						vehicle:  nil,
					},
					{
						occupied: true,
						slotNo:   3,
						vehicle:  &Vehicle{registrationNo: "KA01-1523", color: "Black"},
					},
				},
			},
			2,
			errors.New("Slot already empty"),
			true,
		},
		{
			"TestCase 3: Unpark Vehicle from parking slot where slot no exceeds capacity",
			parklot{
				capacity: 3,
				slots: []*ParkingSlot{
					{
						occupied: true,
						slotNo:   1,
						vehicle:  &Vehicle{registrationNo: "BR01-1234", color: "White"},
					},
					{
						occupied: false,
						slotNo:   2,
						vehicle:  nil,
					},
					{
						occupied: true,
						slotNo:   3,
						vehicle:  &Vehicle{registrationNo: "KA01-1523", color: "Black"},
					},
				},
			},
			4,
			errors.New("Wrong slot no. provided"),
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingLot := &ParkingLot{
				capacity: test.parklot.capacity,
				slots:    test.parklot.slots,
			}
			err := testParkingLot.UnparkVehicle(test.args)
			if (err != nil) != test.wantErr {
				t.Errorf("\x1b[31;1mParkingLot.UnparkVehicle() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if (err != nil) == test.wantErr && !reflect.DeepEqual(err, test.err) {
				t.Errorf("\x1b[31;1mParkingLot.UnparkVehicle() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
func TestGetRegistrationNoForColor(t *testing.T) {

}
