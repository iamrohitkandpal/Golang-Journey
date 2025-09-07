package main

import "fmt"

func printSlice[T int | string | bool](items []T) {
	fmt.Println("Using Print Slice with Generics")
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice[T comparable, V string](items []T, name V) {
	fmt.Println("Using Print Slice with Comparable")
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type Stack[T any] struct {
	elements []T
}

func main() {
	myStack := Stack[int]{
		elements: []int{1, 2, 3},
	} 
	fmt.Println(myStack)

	names := []string{"Rohit", "Navin"}
	printSlice([]int{1, 2, 3})
	printSlice(names)


	printStringSlice([]string{"a", "b"}, "V Generic")
}
