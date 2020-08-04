package main

import "fmt"

func main() {
	cities := []string{"Goa", "Mangalore", "Kochi", "Trivandrum"}

	fmt.Println()
	fmt.Println("Accessing slices ")
	fmt.Println("-----------------")
	fmt.Println(cities[0])
	fmt.Println(cities[2])
	fmt.Println(cities[0:4])

	cities[0] = "Margoa"
	fmt.Println(cities)
	fmt.Println()

}
