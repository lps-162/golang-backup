package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f32 float32 = 10.6556
	fmt.Println(f32, reflect.TypeOf(f32))

	i32 := int32(f32)
	fmt.Println(i32, reflect.TypeOf(i32))

	f64 := float64(i32)
	fmt.Println(f64, reflect.TypeOf(f64))
}
