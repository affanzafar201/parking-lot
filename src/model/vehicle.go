// Copyright 2019. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package model

type Vehicle struct {
	registrationNo string
	color          string
}

func NewVehicle(registrationNo string, color string) *Vehicle {
	v := new(Vehicle)
	v.registrationNo = registrationNo
	v.color = color
	return v
}

func (v *Vehicle) GetRegistrationNo() string {
	return v.registrationNo
}

func (v *Vehicle) GetColor() string {
	return v.color
}
