package main

import "fmt"

func main() {
	var cities = map[string]int{"Goa": 32, "Bangalore": 28, "Kochi": 35}

	fmt.Println(cities)
	for key := range cities {
		delete(cities, key)
	}

	fmt.Println("After truncating all map elements")
	fmt.Println(cities)

	fmt.Println()
	var employees = make(map[string]int)
	employees["vijay"] = 80
	employees["rahul"] = 60
	employees["john"] = 75
	fmt.Println("Printing employees map")
	fmt.Println(employees)
	fmt.Println()

	employees = make(map[string]int)
	fmt.Println("Truncating employees map")
	fmt.Println(employees)

}
