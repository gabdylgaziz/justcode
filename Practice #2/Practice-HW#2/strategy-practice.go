package main

import (
	"fmt"
)

type PaymentStrategy interface {
	Pay(amount int)
}

type CreditCardPayment struct {}

func (c CreditCardPayment) Pay(amount int) {
	fmt.Println("Payment is success to", amount, "tg, with Credit Card")
}

type KaspiQrPayment struct {}

func (k KaspiQrPayment) Pay(amount int) {
	fmt.Println("Payment is success to", amount, "tg, with Credit Card")
}

type ShoppingCart struct {
	PaymentStrategy PaymentStrategy
}

func (s *ShoppingCart) SetStrategy(strategy PaymentStrategy){
	s.PaymentStrategy = strategy
}

func (s *ShoppingCart) Checkout(amount int){
	s.PaymentStrategy.Pay(amount)
}


func main()  {
	strategy := &ShoppingCart{PaymentStrategy: CreditCardPayment{}}
	strategy.Checkout(10000)
	strategy.SetStrategy(KaspiQrPayment{})
	strategy.Checkout(15000)
}

