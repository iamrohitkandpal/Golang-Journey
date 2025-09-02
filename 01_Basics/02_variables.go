package main

import "fmt"

func main() {
	// Variable declaration with explicit type
	var name string = "Gopher"
	
	// Variable declaration with type inference
	age := 5
	
	// Multiple variable declaration
	var (
		isAwesome bool   = true
		weight    float64 = 0.75
	)
	
	// Constants
	const maxAge = 100
	
	// Display all variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Awesome?", isAwesome)
	fmt.Println("Weight:", weight)
	fmt.Println("Max Age:", maxAge)
	
	// Changing variable values
	name = "Super Gopher"
	age = 6
	
	fmt.Println("Updated Name:", name)
	fmt.Println("Updated Age:", age)
	
	// This would cause an error - constants cannot be changed
	// maxAge = 101
}
