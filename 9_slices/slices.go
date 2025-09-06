package main

import (
	"fmt"
	"slices"
)

// Slice -> dynamic
// Most used construct in GO
// + Useful Methods

func main()  {
	// Uninitialized slice is nil
	var nums []int
	fmt.Println(nums)

	// Initialized Sized array - Creagting a slice 
	var numms = make([]int, 2, 3)
	// fmt.Println(numms)

	//Capacity of array
	// fmt.Println(cap(numms))

	numms = append(numms, 1)
	numms = append(numms, 2)
	// fmt.Println(numms)
	// fmt.Println(cap(numms))
	
	// Another way to make slice
	nuums := []int{}
	nuums = append(nuums, 1)
	nuums = append(nuums, 2)
	nuums = append(nuums, 3)
	// fmt.Println(nuums)
	// fmt.Println(cap(nuums))


	// Copy funtion
	var numss = make([]int, 0, 5)
	numss = append(numss, 2)
	var numss2 = make([]int, len(numss))
	
	copy(numss2, numss)

	// fmt.Println(numss, numss2)

	// Slice operator
	var nnums = []int{1, 2, 3}
	fmt.Println(nnums[0:])

	// Slice Package
	var nnums1 = []int{1, 2, 4, 3}
	var nnums2 = []int{1, 2, 4, 5}

	fmt.Println(slices.Equal(nnums1, nnums2))

	// 2d Slices
	var nums2D = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums2D)
}