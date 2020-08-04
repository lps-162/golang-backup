package main

import "fmt"

func main() {

	a := 34
	b := 45

	result := a + b
	fmt.Printf("Result of a + b = %d\n", result)
	fmt.Println()

	result2 := a == b
	fmt.Println("a == b:", result2)

	result3 := a != b
	fmt.Println("a != b :", result3)

	if a < 100 && b < 100 {
		fmt.Println("Both numbers are less than 100")
	} else if a > 100 && b > 100 {
		fmt.Println("Both numbers are larger than 100")
	}

}
