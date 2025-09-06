package main

import (
	"fmt"
	"time"
)

// Customer Struct
type customer struct {
	name  string
	phone string
}

// Order Struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // In Nanosecond
	customer
}

func newOrder(id string, amount float32, status string) *order {
	// Constructor Hack in Go
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// Receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// Constructor Usage
	myOrder := newOrder(" 0", 90.80, "Delivered")
	fmt.Println(myOrder)

	// Manual Struct Usage
	newCustomer := customer {
		name: "John",
		phone: "999999090",
	}
	var firstOrder order = order{
		id:     "01",
		amount: 50.00,
		status: "Received",
		customer: newCustomer,
	}

	firstOrder.changeStatus("Returned")

	var myOrder2 order = order{
		id:     "02",
		amount: 150.00,
		status: "Dispatched",
		customer: customer {
			name: "Raju",
			phone: "0909090909",
		},
	}

	fmt.Println(myOrder2.getAmount())

	fmt.Println("Order Struct", firstOrder)
	fmt.Println("Order Struct", myOrder2)

	// Fast Usage of Struct inline
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
