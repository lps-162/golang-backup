package main

import "fmt"

func main() {
	var cities = map[string]int{"Goa": 32, "Bangalore": 28, "Kochi": 35}

	for key, value := range cities {
		fmt.Println("Key: ", key, "=>", value)
	}
}
