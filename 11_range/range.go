package main

import "fmt"

func main() {

	// Range in Slice
	nums := []int{6, 7, 8}
	sum := 0
	for index, i := range nums {
		sum = sum + i
		fmt.Println(sum, index)
	}

	// Range in Map
	mp := map[string]string{"fname": "John", "lname": "Doe"}

	for k, v := range mp {
		fmt.Println(k, v)
	}

	// Range over String (Gives Unicode)
	// Starting byte of Rune
	// 255 -> 1 byte, 2 byte
	for i, c := range "Golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}

}