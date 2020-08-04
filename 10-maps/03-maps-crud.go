package main

import "fmt"

func main() {
	var cities = map[string]int{"Goa": 32, "Bangalore": 28, "Kochi": 35}

	fmt.Println("Accessing map element and printing full map")
	fmt.Println("--------------------------------------------")
	fmt.Println(cities["Bangalore"])
	fmt.Println(cities)
	fmt.Println()

	cities["Coimbatore"] = 36
	cities["Madurai"] = 40

	fmt.Println("After adding 2 more values in map")
	fmt.Println("---------------------------------")
	fmt.Println(cities)
	fmt.Println()

	fmt.Println("After Updating a map element")
	fmt.Println("---------------------------------")
	cities["Madurai"] = 38
	fmt.Println(cities)
	fmt.Println()

	fmt.Println("After deleting an element in map")
	fmt.Println("---------------------------------")
	delete(cities, "Goa")
	fmt.Println(cities)
	fmt.Println()
}
