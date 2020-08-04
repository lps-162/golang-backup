package main

import "fmt"

func main() {
	cities := []string{"Kochi", "Coimbatore", "Mangalore", "Trivandrum", "Goa"}

	fmt.Println("\nLooping method 1 \n")
	for index, element := range cities {
		fmt.Println(index, "--", element)
	}

	fmt.Println("\nLooping method 2 \n")
	for _, value := range cities {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\nLooping method 3 \n")
	for range cities {
		fmt.Println(cities[j])
		j++
	}
}
