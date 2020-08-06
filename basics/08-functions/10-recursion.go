package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {

	result := fact(5)
	fmt.Println("Factorial of 5 :", result)
}
