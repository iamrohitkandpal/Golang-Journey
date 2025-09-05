package main

import "fmt"

// for -> Only construct in Go for looping
func main()  {
	// While Loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}