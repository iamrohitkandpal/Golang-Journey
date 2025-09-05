package main

import "fmt"

func main() {
	age := 17

	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not Adult")
	}

	if age >= 18 {
		fmt.Println("Person is Adult")
	} else if  age >= 12 {
		fmt.Println("Person is Teen")
	} else {
		fmt.Println("Person is Child")
	}

	// Logical Operators
	var role = "admin"
	var hasPermissions = true

	if role == "admin" && hasPermissions {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Not Allowed")
	}

	// ifElse + Logical Operators
	if age:= 15; age >= 18 {
		fmt.Println("Adult Vibes of age: ", age)
	} else if age >= 12 {
		fmt.Println("Teen Vibes of age:", age)
	}
}
