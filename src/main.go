package main

import (
	"fmt"
	"model"
)

func main() {
	fmt.Println("vim-go")
	v := model.NewVehicle("Abcd", "brown")
	fmt.Printf("%+v", v)
}
