package main

import "fmt"

func main() {
	var itCities = map[string]int{"Hyderabad": 40, "Bangalore": 28, "Kochi": 35}
	var newCities = map[string]int{"Coimbatore": 40, "Trivandrum": 28}

	for k, v := range newCities {
		itCities[k] = v
	}

	fmt.Println("The IT Cities map after merging with new cities")
	fmt.Println("-----------------------------------------------")
	fmt.Println(itCities)
	for k, v := range itCities {
		fmt.Println(k, "=>", v)
	}
	fmt.Println()
}
