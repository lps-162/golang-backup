package main

import "fmt"

func main() {
	func(a int, b int) {
		fmt.Println("Product of a and b is:", a*b)
	}(10, 5)

	area := func(l int, b int) int {
		return l * b
	}

	fmt.Println("Area of rectangle:", area(34, 45))
}
