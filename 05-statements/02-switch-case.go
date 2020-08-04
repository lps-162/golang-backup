package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()

	fmt.Println(today.Day())

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th")
	case 10:
		fmt.Println("Today is 10th")
	case 30:
		fmt.Println("Almost end of month")
	default:
		fmt.Println("One off date")
	}

	switch today.Day() {
	case 10, 20:
		fmt.Println("Wash ur bikes")
	case 25, 26, 27:
		fmt.Println("Prepare for end of month")
	case 30, 31:
		fmt.Println("tough times")
	case 1, 2, 3, 4, 5:
		fmt.Println("Salary credited: Party time")
	default:
		fmt.Println("Not the day you are looking for")
	}

	city := "Trivandrum"
	switch city {
	case "Bangalore":
		fmt.Println("Capital of Karnataka")
	case "Chennai":
		fmt.Println("Capital of Tamilnadu")
	case "Trivandrum":
		fmt.Println("Capital of Kerala")
	default:
		fmt.Println("Not in the list")
	}
}
