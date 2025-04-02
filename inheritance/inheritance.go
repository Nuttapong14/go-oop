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
		Vehicle  
	}

	Bike struct {
		Vehicle  
	}
)


// Method for Vehicle
func (v Vehicle) Drive() {
	fmt.Println(v.Brand, "is driving at", v.Speed, "km/h")
}

// Overriding Drive() for Car
func (c Car) Drive() {
	fmt.Println(c.Brand, "car is driving at", c.Speed, "km/h")
}

// Overriding Drive() for Bike
func (b Bike) Drive() {
	fmt.Println(b.Brand, "bike is driving at", b.Speed, "km/h")
}