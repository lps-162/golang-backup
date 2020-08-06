package main

import (
	"fmt"
	"math"
)

var life int = 145

// functions passing 2 ints
// functions can return one or more values
func calcRectArea(length int, width int) int {
	area := length * width
	return area
}

func calcRectDiagonal(length, width int) float64 {
	sum_of_squares := (length * length) + (width * width)
	length_of_diagonal := math.Sqrt(float64(sum_of_squares))
	return length_of_diagonal
}

func main() {

	result := calcRectArea(10, 12)
	fmt.Println("Area of Rectangle :", result)

	diagonal := calcRectDiagonal(10, 12)
	fmt.Printf("Length of Diagonal : %.2f", diagonal)

	empname := "Shivan LP"
	fmt.Println(empname)
}
