package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	fmt.Println("Switch fallthrough")
	fmt.Println("===================")
	switch today.Day() {
	case 1, 2, 3, 4, 5:
		fmt.Println("Salary credited: Party time")
		fallthrough
	case 10, 20:
		fmt.Println("Wash ur bikes")
	case 25, 26, 27:
		fmt.Println("Prepare for end of month")
	case 30, 31:
		fmt.Println("tough times")
	default:
		fmt.Println("Not the day you are looking for")
	}
	fmt.Println()

	fmt.Println("Switch, conditional case statements")
	fmt.Println("================================")
	actualDay := today.Day()
	switch {
	case actualDay < 10:
		fmt.Println("Salary week, Awesome party times")
	case actualDay < 20:
		fmt.Println("Ok, Manageable financially")
	case actualDay < 28:
		fmt.Println("Oh my god, tough week man")
	case actualDay < 31:
		fmt.Println("Literally empty wallete, Having left over peanuts")
	}
	fmt.Println()

	fmt.Println("Switch initializer")
	fmt.Println("==================")
	switch thisDay := time.Now().Day(); {
	case thisDay < 5:
		fmt.Println("Time to work")
	case thisDay > 10:
		fmt.Println("Training over, time to relax")
	default:
		fmt.Println("Not your day")
	}
}
