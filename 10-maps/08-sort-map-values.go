package main

import (
	"fmt"
	"sort"
)

func main() {
	var itCities = map[string]int{"Hyderabad": 40, "Bangalore": 28, "Kochi": 35}

	values := make([]int, 0, len(itCities))

	for _, v := range itCities {
		values = append(values, v)
	}

	sort.Ints(values)

	for _, v := range values {
		fmt.Println(v)
	}
}
