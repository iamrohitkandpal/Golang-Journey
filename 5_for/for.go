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

	// Classical For Loop
	for i:= 0; i <= 20; i++ {
		if i == 16 {
			break
		}

		if i % 4 == 0 {
			continue
		}

		if(i % 2 == 0) {
			fmt.Println(i)
		}
	}

	// Range
	for i:= range 4 {
		fmt.Println(i)
	}
}