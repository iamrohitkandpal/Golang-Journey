package main

import (
	"fmt"
	"maps"
)

func main() {
	// Creating map

	m := make(map[string]string)

	// Setting element in map
	m["name"] = "Golang"
	m["area"] = "Backend"

	// Getting element from map
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["nothing"]) // ==> Return Empty value

	fmt.Println(len(m))

	// delete(m, "area") 			// ==> Deletes the element
	// clear(m)						// ==> Clear the map

	mp := map[string]int{"prices": 40, "phones": 3}

	k, check := mp["prices"]
	fmt.Println(k)
	if check {
		fmt.Println("OK")
	} else {
		fmt.Println("Not OK")
	}

	m1 := map[string]int{"price": 50, "phones": 9}
	m2 := map[string]int{"price": 50, "phones": 9}

	fmt.Println(maps.Equal(m1, m2))
}
