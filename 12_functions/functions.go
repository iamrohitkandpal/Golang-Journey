package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, int) {
// 	return "golang", "js", 2
// }

// func processIt(fn func(a int) int)  {
// 	fn(10)
// }

func processIt() func(a int) int {
	return func (a int) int {
		return 4
	}
}

func main() {
	// res := add(3, 5)
	// res1, res2, _ := getLanguages()

	// fn := func(a int) int  {
	// 	return 2
	// }

	fn := processIt()
	fn(6)

	fmt.Println(fn(10))

	// fmt.Println(res1, res2)
}