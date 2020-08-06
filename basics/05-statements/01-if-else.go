package main

import "fmt"

func main() {

	var city = "Bangalore"
	isCapital := true

	if isCapital {
		fmt.Println(city, "is a capital")
	} else {
		fmt.Println("Not a capital")
	}

	temperature := 40

	if temperature <= 25 {
		fmt.Println("Very cool but manageable")
	} else if temperature < 35 {
		fmt.Println("Bangalore is a cool city")
	} else {
		fmt.Println("Its nowadays becoming hotter")
	}

	if val := 100; val == 100 {
		fmt.Println("Sure victory it is")
	}
}
