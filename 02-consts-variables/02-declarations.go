package main

import (
	"fmt"
)

var s = "Goa"
var m string = "Mangalore"

// city := "Coimbatore" // this short form is not allowed outside functions

func main() {
	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}

// Note that short variable declaration is allowed only for declaring local variables,
// variables declared within the function. When you declare variables outside the
// function, you must do so using the var keyword.
