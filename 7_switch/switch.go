package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple Switch
	i := 5

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Second")
	case 3:
		fmt.Println("Third")
	default:
		fmt.Println("Default")
	}

	// Multiple Condition Switch
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("It's Weekday")
	case time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's WOrkday")
	}

	// Type Switch
	whoAmI := func(i interface{})  {
		switch i.(type) {
		case int:
			fmt.Println("It's a integer")
		case string:
			fmt.Println("It's a string")
		default:
			fmt.Println("Something else")
		}
	}

	whoAmI(true)
	whoAmI("true")
}