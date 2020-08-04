package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println()
	fmt.Println("int to string Conversion")
	fmt.Println("============================")
	var i int64 = -654
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)
	fmt.Println()

	var s string = strconv.FormatInt(i, 10)
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(len(s))
	fmt.Println()

	fmt.Println("float to string Conversion")
	fmt.Println("============================")
	var f float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f)

	var fToS string = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(fToS))
	fmt.Println(fToS, len(fToS))
	fmt.Println()

	fmt.Println("bool to string Conversion")
	fmt.Println("============================")
	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var bToS string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(bToS))
	fmt.Println(bToS, len(bToS))
}
