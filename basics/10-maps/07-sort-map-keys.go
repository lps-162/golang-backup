package main

import (
	"fmt"
	"sort"
)

func main() {
	var itCities = map[string]int{"Hyderabad": 40, "Bangalore": 28, "Kochi": 35}

	keys := make([]string, 0, len(itCities))

	for k := range itCities {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)

	fmt.Println("Cities map in sorted key order")
	fmt.Println("------------------------------")
	for _, k := range keys {
		fmt.Println(k, itCities[k])
	}
}
