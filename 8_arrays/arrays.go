package main

import "fmt"

// - Fixed size, that is predicatable
// - Memory Optimization
// - Constant TIme Access

func main() {
	// Zeros filled with int
	var nums [5]int
	// Empty strings filled with string
	var strings [5]string
	// False filled with bool
	var bools [5]bool

	nums[0] = 1
	strings[0] = "Hi"
	bools[0] = true

	// Length of an array
	fmt.Println(len(nums))

	fmt.Println(nums)
	fmt.Println(strings)
	fmt.Println(bools)

	// Instant Declaration of array values
	numms := [3]int{1, 2, 3}
	fmt.Println(numms)

	// 2d Arrays in Go
	numms2D := [2][2]int{{3, 4},{5, 6}}
	fmt.Println(numms2D)
}