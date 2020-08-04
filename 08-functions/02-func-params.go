package main

import "fmt"

func add(a int, b int) {
	total := 0
	total = a + b
	fmt.Println("Total is :", total)
}

func main() {
	add(10, 23)
}
