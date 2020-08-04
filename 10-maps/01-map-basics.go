package main

import (
	"fmt"
	"reflect"
)

// map is a key value pair, faster lookups
// hash table implementation
// maps are unordered collections

func main() {

	var cities = map[string]int{"Goa": 32, "Bangalore": 28, "Kochi": 35}

	fmt.Println(cities)

	var customers = map[string]int{}
	fmt.Println(customers)
	fmt.Println(reflect.TypeOf(customers))
}
