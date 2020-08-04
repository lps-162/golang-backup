package main

import "fmt"

func findIndex(arr [3]string, x string) int {
	for i, val := range arr {
		if val == x {
			return i
		}
	}
	return -1

}

func contains(arr [3]string, x string) bool {
	for _, val := range arr {
		if val == x {
			return true
		}
	}

	return false
}

func main() {

	cities := [3]string{"Bangalore", "Goa", "Trivandrum"}

	isAvailable := contains(cities, "Goa")
	if isAvailable {
		fmt.Println("Goa found in the list of cities")
	}
	fmt.Println(cities)

	index := findIndex(cities, "Trivandrum")
	if index >= 0 {
		fmt.Println("Trivandrum found at index :", index)
	} else {
		fmt.Println("Trivandrum not found in the list")
	}
}
