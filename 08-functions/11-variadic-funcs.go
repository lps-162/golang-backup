package main

import "fmt"

func listCities(cities ...string) {

	fmt.Println("List of Cities :")
	fmt.Println("================")
	fmt.Println(cities)

	for _, val := range cities {
		fmt.Println(val)
	}

	//fmt.Println(cities[5])
}

func main() {

	listCities()
	listCities("Goa")
	listCities("Bangalore", "Goa")
	listCities("Coimbatore", "Trivandrum", "Kochi")
}
