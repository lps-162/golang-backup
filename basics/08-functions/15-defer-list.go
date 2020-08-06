package main

import "fmt"

func printMessage(val int) {
	fmt.Println("Hello world :", val)
}

func main() {
	for i := 0; i < 5; i++ {
		defer printMessage(i)
	}
	defer printMessage(145)
}
