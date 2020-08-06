package main

import (
	"fmt"
	"strconv"
)

func main() {
	strVar := "1sdf00"
	intVar, err := strconv.Atoi(strVar)
	if err != nil {
		fmt.Println("Couldnt convert string to int")
	} else {
		fmt.Printf("intVar %v and %T", intVar, intVar)
	}
	fmt.Println()

	strVar1 := "-52541"
	intVar1, _ := strconv.ParseInt(strVar1, 10, 32)
	fmt.Printf("intVar1 %v and %T", intVar1, intVar1)
	fmt.Println()

	strVar2 := "101010101010101010"
	intVar2, _ := strconv.ParseInt(strVar2, 10, 64)
	fmt.Printf("intVar2 %v and %T", intVar2, intVar2)
	fmt.Println()
}
