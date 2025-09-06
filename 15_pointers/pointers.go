package main

import "fmt"

// By value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In ChangeNum", num)
// }

// By reference
func changeRef(num *int) {
	*num = 5
	fmt.Println("In ChangeNum", *num)
}


func main() {
	num := 1

	// changeNum(num)
	changeRef(&num)
	fmt.Println("Memory Address", &num)

	fmt.Println("After ChangeNum in main", num)
}