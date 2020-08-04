package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"Bangalore", "is", "a", "party", "city", "too"}

	sort.Strings(words)
	fmt.Println(words)

	hourlyRates := []int{90, 80, 45, 23, 40, 65}
	sort.Ints(hourlyRates)
	fmt.Println(hourlyRates)

	buyPrices := []float64{15.4, 9.99, 4.45, 2.34, 10.09}
	sort.Float64s(buyPrices)
	fmt.Println(buyPrices)
}
