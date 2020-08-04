package main

import (
	"fmt"
	"reflect"
)

func rectOuter() func() {
	l := 10
	b := 5

	return func() {
		fmt.Println("Area of Rectangle :", l*b)
	}
}
func main() {
	area := rectOuter()
	fmt.Println("Type of area :", reflect.TypeOf(area))

	area()
}
