package main

import (
	"fmt"
	"reflect"
)

// Slice is a segment of a dynamic array that can grow and shrink
// Like Arrays, Slices are indexable and have a length
// Slices have a capacity and length property

func main() {
	var intSlice []int
	var strSlice []string

	fmt.Println("Slices basics")
	fmt.Println("-------------")
	fmt.Println("Integer Slice :", intSlice)
	fmt.Println("String Slice :", strSlice)
	fmt.Println(reflect.TypeOf(intSlice))
	fmt.Println(reflect.TypeOf(strSlice))
	fmt.Println(reflect.TypeOf(intSlice).Kind())
	fmt.Println(reflect.TypeOf(strSlice).Kind())
	fmt.Println()
}
