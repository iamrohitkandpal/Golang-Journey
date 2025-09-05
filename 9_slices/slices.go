package main

import "fmt"

// Slice -> dynamic
// Most used construct in GO
// + Useful Methods

func main()  {
	// Uninitialized slice is nil
	var nums []int
	fmt.Println(nums)

	// Initialized Sized array - Creagting a slice 
	var numms = make([]int, 0, 3)
	fmt.Println(numms)

	//Capacity of array
	fmt.Println(cap(numms))

	numms = append(numms, 1)
	numms = append(numms, 2)
	fmt.Println(numms)
	fmt.Println(cap(numms))
	
	// Another way to make slice
	nuums := []int{}
	nuums = append(nuums, 1)
	nuums = append(nuums, 2)
	nuums = append(nuums, 3)
	fmt.Println(nuums)
	fmt.Println(cap(nuums))

}