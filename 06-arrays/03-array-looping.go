package main

import "fmt"

func main() {

	cities := [5]string{"Goa", "Bangalore", "Trivandrum"}

	fmt.Println("Looping method 1")
	fmt.Println("----------------")
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	fmt.Println()

	fmt.Println("Looping Method 2")
	fmt.Println("----------------")
	for index, elem := range cities {
		fmt.Println(index, ":", elem)
	}
	fmt.Println()

	fmt.Println("Looping Method 3")
	fmt.Println("----------------")
	techs := [3]string{"Golang", "React", "GCloud"}
	for _, value := range techs {
		fmt.Println(value)
	}
	fmt.Println()

	j := 0
	fmt.Println("Looping Method 4")
	fmt.Println("----------------")
	for range techs {
		fmt.Println(techs[j])
		j++
	}
}
