// Copyright 2019. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package parkinglot

type Vehicle struct {
	registrationNo string
	color          string
}

func (v *Vehicle) Init(registrationNo string, color string) {
	v.registrationNo = registrationNo
	v.color = color
}

func (v *Vehicle) SetRegistrationNo(registrationNo string) {
	v.registrationNo = registrationNo
}

func (v *Vehicle) GetRegistrationNo() string {
	return v.registrationNo
}

func (v *Vehicle) SetColor(color string) {
	v.color = color
}

func (v *Vehicle) GetColor() string {
	return v.color
}
