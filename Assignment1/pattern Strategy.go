package main

import (
	"errors"
	"fmt"
)

type Payment struct {
	amount     float64
	payFunc    func() (string, error)
	methodName string
}

type PaymentOption func(*Payment)

func PayWithCreditCard() PaymentOption {
	return func(p *Payment) {
		p.methodName = "Credit Card"
		p.payFunc = func() (string, error) {
			return fmt.Sprintf("Paid %.2f kzt by Credit Card", p.amount), nil
		}
	}
}

func PayWithCash() PaymentOption {
	return func(p *Payment) {
		p.methodName = "Cash"
		p.payFunc = func() (string, error) {
			return fmt.Sprintf("Paid %.2f kzt in Cash", p.amount), nil
		}
	}
}

func PayWithCrypto(hasFunds bool) PaymentOption {
	return func(p *Payment) {
		p.methodName = "Crypto"
		p.payFunc = func() (string, error) {
			if hasFunds {
				return fmt.Sprintf("Paid %.2f kzt with Crypto", p.amount), nil
			}
			return "", errors.New("insufficient funds in Crypto Wallet")
		}
	}
}

func NewPayment(amount float64, options ...PaymentOption) *Payment {
	p := &Payment{amount: amount}

	for _, option := range options {
		option(p)
	}

	return p
}

func main() {
	shoppingCart := []struct {
		item  string
		price float64
	}{
		{"Phone", 150000.0},
		{"Laptop", 550000.0},
		{"Car", 9000000.0},
	}
	walletAmount := 6000000.0
	payments := []*Payment{
		NewPayment(walletAmount, PayWithCreditCard()),
		NewPayment(walletAmount, PayWithCash()),
		NewPayment(walletAmount, PayWithCrypto(true)),
		NewPayment(walletAmount, PayWithCrypto(false)),
	}

	for _, payment := range payments {
		remainingAmount := walletAmount
		fmt.Printf("Attempting to pay using %s...\n", payment.methodName)
		for _, item := range shoppingCart {
			if remainingAmount >= item.price {
				payment.amount = item.price
				result, err := payment.payFunc()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
				remainingAmount -= item.price
			} else {
				fmt.Printf("Not enough funds for %s using %s.\n", item.item, payment.methodName)
			}
		}
		fmt.Println()
	}
}
