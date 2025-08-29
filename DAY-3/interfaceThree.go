package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64)
}

type CreditCard struct {
	holder string
}

type PayPal struct {
	email string
}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Paid %.2f using CreditCard (%s)\n", amount, c.holder)
}

func (p PayPal) Pay(amount float64) {
	fmt.Printf("Paid %.2f using CreditCard (%s)\n", amount, p.email)
}

func main() {
	// create objects
	cc := CreditCard{holder: "Alice"}
	pp := PayPal{email: "bob@example.com"}

	// interface variables
	var processor PaymentProcessor

	processor = cc
	processor.Pay(100.0)

	processor = pp
	processor.Pay(100.0)
}
