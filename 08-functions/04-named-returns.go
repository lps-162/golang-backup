package main

import "fmt"

func calculateArea(x int, y int) (area int) {
	area = x * y
	return
}
func main() {
	area := calculateArea(9, 10)
	fmt.Println("Area is :", area)
}
