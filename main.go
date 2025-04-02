// main.go
package main

import (
	"fmt"

	"github.com/Nuttapong14/go-oop/abstraction"
	"github.com/Nuttapong14/go-oop/encapsulation"
	"github.com/Nuttapong14/go-oop/polymorphism"
    "github.com/Nuttapong14/go-oop/inheritance"
)

func main() {
	// Abstraction
	Truck := abstraction.Bike{ Type: "Road Bike" }

	Car := abstraction.Car{ Seats: 4, Year:  2020, }

	Truck.Drive()
	Car.Drive()

    print("\n")

	// Polymorphism
	credit := polymorphism.CreditCard{CardNumber: "1234-5678-9012"}
	paypal := polymorphism.PayPal{Email: "user@example.com"}
	crypto := polymorphism.Crypto{WalletAddress: "xyz123"}

	fmt.Println("🛒 Use Polymorphism")

	credit.Pay(100.50)
	paypal.Pay(200.75)
	crypto.Pay(300.00)

    print("\n")

	// Encapsulation
    fmt.Println("🛒 Use Encapsulation")
	encapsulation.PaymentPaid(credit, 100.50)
	encapsulation.PaymentPaid(paypal, 200.75)
	encapsulation.PaymentPaid(crypto, 300.00)

    // Inheritance
    	// Creating a Car object
	myCar := inheritance.Car{
		Vehicle: inheritance.Vehicle{
			Brand: "Toyota",
			Speed: 120,
		},
		Seats: 5,
	}

	// Creating a Bike object
	myBike := inheritance.Bike{
		Vehicle: inheritance.Vehicle{
			Brand: "Yamaha",
			Speed: 80,
		},
		Type: "Sport",
	}

    print("\n")
	// Calling Drive() method
    fmt.Println("Inheritance")
	myCar.Drive()
	myBike.Drive()

}
