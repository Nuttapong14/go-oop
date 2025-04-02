package encapsulation

import "fmt"

type (
	// Payment interface
	Payment interface {
		Pay(amount float64)
	}

	CreditCard struct {
		CardNumber string
	}

	PayPal struct {
		Email string
	}

	Crypto struct {
		WalletAddress string
	}
)

func PaymentPaid(p Payment, amount float64) {
	p.Pay(amount)
}

// PaymentMethod interface (Abstraction)
type PaymentMethod interface {
	Pay(amount float64)
}

// Implement Pay() for CreditCard
func (c CreditCard) Pay(amount float64) {
	fmt.Println("💳 Paid", amount, "using Credit Card:", c.CardNumber)
}

// Implement Pay() for PayPal
func (p PayPal) Pay(amount float64) {
	fmt.Println("🅿️ Paid", amount, "using PayPal:", p.Email)
}

// Implement Pay() for Crypto
func (c Crypto) Pay(amount float64) {
	fmt.Println("₿ Paid", amount, "using Crypto Wallet:", c.WalletAddress)
}