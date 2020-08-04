package main

import "fmt"

func main() {
	var x, y = 15, 25

	fmt.Printf("x = %d and y = %d", x, y)
	fmt.Println()
	fmt.Println("==", x == y)
	fmt.Println("!=", x != y)
	fmt.Println("<", x < y)
	fmt.Println("<=", x <= y)
	fmt.Println(">", x > y)
	fmt.Println(">=", x >= y)
}
