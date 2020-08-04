package main

import "fmt"

func removeElement(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	cities := []string{"Kochi", "Coimbatore", "Mangalore", "Trivandrum", "Goa"}

	fmt.Println("Cities slice before removing")
	fmt.Println("----------------------------")
	fmt.Println(cities)

	cities = removeElement(cities, 2)

	fmt.Println("Cities slice after removing")
	fmt.Println(cities)
}
