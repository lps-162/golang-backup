package main

import "fmt"

func main() {
	cities := []string{"Kochi", "Coimbatore", "Mangalore", "Trivandrum", "Goa"}

	fmt.Printf("cities: %v\n", cities)

	fmt.Printf(":2 %v\n", cities[:2])

	fmt.Printf("1:3 %v\n", cities[1:3])

	fmt.Printf("2: %v\n", cities[2:])

	fmt.Printf("2:5 %v\n", cities[2:5])

	fmt.Printf("0:3 %v\n", cities[0:3])

	fmt.Printf("Last element: %v\n", cities[4])
	fmt.Printf("Last element: %v\n", cities[len(cities)-1])
	fmt.Printf("Last element: %v\n", cities[4:])

	fmt.Printf("All elements: %v\n", cities[0:len(cities)])

	fmt.Printf("Last two elements: %v\n", cities[3:len(cities)])
	fmt.Printf("Last two elements: %v\n", cities[len(cities)-2:len(cities)])

	fmt.Println(cities[:])
	fmt.Println(cities[0:])
	fmt.Println(cities[0:len(cities)])
}
