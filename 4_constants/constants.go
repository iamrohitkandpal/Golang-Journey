package main

import "fmt"

func main() {
	const name string = "Rohit"
	const surname = "Surname"

	const (
		fName  = "Navin"
		fSName = "Kandpal"
	)

	fmt.Println(name, surname, fName, fSName)
}