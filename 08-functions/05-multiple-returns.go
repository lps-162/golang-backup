package main

import "fmt"

func calcValues(l int, b int) (area int, perimeter int) {
	area = l * b
	perimeter = 2 * (l + b)
	return
}

func rectangle(l int, b int) (int, int) {
	area := l * b
	perm := 2 * (l + b)
	return area, perm
}

func main() {
	length, breadth := 9, 10
	area, perimeter := calcValues(length, breadth)

	fmt.Println("Named Return demo")
	fmt.Println("=================")
	fmt.Println("Area and Perimeter of ", area, perimeter)
	fmt.Println()

	length, breadth = 5, 8
	area, perimeter = rectangle(length, breadth)
	fmt.Println("Unnamed Return demo")
	fmt.Println("=================")
	fmt.Println("Area and Perimeter of ", area, perimeter)
	fmt.Println()
}
