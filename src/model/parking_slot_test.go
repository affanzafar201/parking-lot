package model

import (
	"errors"
	"reflect"
	"testing"
)

/*func TestCreateParkingLot(t *testing.T) {
	testParkingLot := CreateParkingLot(6)
	if testParkingLot == nil {
		t.Errorf("Failed to create parking lot")
	} else if testParkingLot.capacity != 6 {
		t.Errorf("Failed to create parking lot with want capacity %v, got %v", testParkingLot.capacity, 6)
	}
}*/

func TestAllotVehicle(t *testing.T) {
	type parkslot struct {
		occupied bool
		slotNo   int
		vehicle  *Vehicle
	}
	type testVehicle struct {
		car Vehicle
	}
	tests := []struct {
		name     string
		parkslot parkslot
		args     testVehicle
		want     error
		wantErr  bool
	}{
		{
			"TestCase 1: Parking slot is empty where vehicle is to park",
			parkslot{
				occupied: false,
				slotNo:   1,
				vehicle:  nil,
			},
			testVehicle{car: Vehicle{registrationNo: "KA01-1498", color: "Red"}},
			nil,
			false,
		},
		{
			"TestCase 2: Parking slot is full where vehicle is to park",
			parkslot{
				occupied: true,
				slotNo:   1,
				vehicle:  &Vehicle{registrationNo: "KA01-1491", color: "Red"},
			},
			testVehicle{car: Vehicle{registrationNo: "KA01-1498", color: "Red"}},
			errors.New("Slot is not empty"),
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testParkingSlot := &ParkingSlot{
				occupied: test.parkslot.occupied,
				slotNo:   test.parkslot.slotNo,
				vehicle:  test.parkslot.vehicle,
			}
			err := testParkingSlot.AllotVehicle(&test.args.car)
			if (err != nil) != test.wantErr {
				t.Errorf("\x1b[31;1mParkingSlot.AllotVehicle() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("\x1b[31;1mParkingSlot.AllotVehicle() = %v, want %v\x1b[0m", err, test.want)
			}
		})
	}
}
