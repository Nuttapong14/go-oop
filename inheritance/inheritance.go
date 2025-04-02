package inheritance

import (
	"fmt"
)

type (
	Vehicle struct {
		Brand string
		Speed int
	}

	Car struct {
		Vehicle  // Inheritance
		Seats int
		Year int
	}

	Bike struct {
		Vehicle  // Inheritance
		Type string
	}
)


// Method for Vehicle
func (v Vehicle) Drive() {
	fmt.Println(v.Brand, "is driving at", v.Speed, "km/h")
}

// Overriding Drive() for Car
func (c Car) Drive() {
	fmt.Println(c.Brand, "car is driving at", c.Speed, "km/h with", c.Seats, "seats year", c.Year)
}

// Overriding Drive() for Bike
func (b Bike) Drive() {
	fmt.Println(b.Brand, b.Type, "bike is driving at", b.Speed, "km/h")
}