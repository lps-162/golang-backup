package main

import (
	"fmt"
	"reflect"
)

func main() {

	var intArray [5]int
	var strArray [5]string

	type1 := reflect.TypeOf(intArray)
	type2 := reflect.TypeOf(strArray)
	fmt.Println("Type and Kind of intArray :", type1, type1.Kind())
	fmt.Println("Type and Kind of strArray : ", type2, type2.Kind())
	fmt.Println()

	fmt.Println("Assinging values to array elements")
	fmt.Println("==================================")
	var cities [3]string
	cities[0] = "Bangalore"
	cities[1] = "Goa"
	cities[2] = "Trivandrum"

	fmt.Println("City 0: ", cities[0])
	fmt.Println("City 1: ", cities[1])
	fmt.Println("City 2: ", cities[2])

}
