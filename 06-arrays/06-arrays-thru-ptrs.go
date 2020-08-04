package main

import (
	"fmt"
	"strings"
)

func modifyList(arr *[3]string, val string) {

	for i, elem := range arr {
		if elem == val {
			arr[i] = strings.Repeat(elem, 3)
		}
	}
}

func main() {

	cities := [3]string{"Bangalore", "Goa", "Trivandrum"}

	fmt.Println()
	fmt.Println("Demonstrating pointers to arrays")
	fmt.Println("--------------------------------")
	fmt.Println("Array before modification")
	fmt.Println(cities)
	modifyList(&cities, "Goa")

	fmt.Println()
	fmt.Println("Array after modification")
	fmt.Println(cities)
}
