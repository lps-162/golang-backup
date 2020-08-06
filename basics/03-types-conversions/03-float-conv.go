package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3.1415926535"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Printf("%T, %v\n", f, f)

	s1 := "-3.141"
	f1, _ := strconv.ParseFloat(s1, 8)
	fmt.Printf("%T, %v\n", f1, f1)

	s2 := "-3.141"
	f2, _ := strconv.ParseFloat(s2, 32)
	fmt.Printf("%T, %v\n", f2, f2)
}
