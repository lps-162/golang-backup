package main

import "fmt"

func compute(fn func(int, int) int) int {
	length, breadth := 8, 9
	result := fn(length, breadth)
	return result
}

func calcPerimeter(l int, b int) int {
	return 2 * (l + b)
}

func calcArea(l int, b int) int {
	return l * b
}

func main() {

	perimeter := compute(calcPerimeter)
	area := compute(calcArea)

	fmt.Println("Area and Perimeter :", area, perimeter)
}
