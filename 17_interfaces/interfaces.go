package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, accouny string)
}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorPaymentGw := razorpay{}
	// razorPaymentGw := stripe{}
	// razorPaymentGw.payS(amount)
	p.gateway.pay(amount)
}

type razorpay struct {}


func (r razorpay) pay(amount float32) {
	// Payment Logic
	fmt.Println("Making payment in Razorpay", amount)
}

// type stripe struct{}

// func (s stripe) payS(amount float32) {
// 	// Payment logic
// 	fmt.Println("making payment in Stripe", amount)
// }


type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	// Payment logic
	fmt.Println("Making fake payment for testing", amount)
}

type paypal struct {}

func (p paypal) pay(amount float32) {
	fmt.Println("Making Payment using Paypal", amount)
}
func (p paypal) refund(amount float32, account string) {
	fmt.Println("Making Payment using Paypal", amount)
}

func main() {
	// fakeGw := fakepayment{}
	// razorpayGw = razorpay{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}