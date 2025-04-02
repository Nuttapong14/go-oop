// File Name: abstraction.go
package abstraction

import "fmt"

type (
    Vehicle interface {
        Drive()
    }
    
    Car struct {
        Seats int
        Year int
    }

    Bike struct {
        Type string
    }

)

// Vehicle interface (Abstraction)

func (v Bike) Drive() {
    fmt.Println("Bike is drive :", v.Type)
}

func (v Car) Drive() {
    fmt.Println("Car is driving : ", v.Seats, " Can seat : ", v.Year)
}
