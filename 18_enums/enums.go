package main

import "fmt"

//Enumerated Types

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delevired"
)

func changedOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changedOrderStatus(Delivered)
}
