package main

import (
	"fmt"
	"reflect"
)

func main() {
	var customerName = "Swetha"
	var age = 25
	var email = "swetha.190@gmail.com"

	// printing the values of variables
	fmt.Println("Variables values")
	fmt.Println("================")
	fmt.Println("Customer Name :", customerName)
	fmt.Println("Customer Age :", age)
	fmt.Println("Customer Email :", email)
	fmt.Println()

	// fmt.Printf can be used to print the type of a variable
	fmt.Println("Variable types using Printf")
	fmt.Println("===========================")
	fmt.Printf("Type of customerName : %T\n", customerName)
	fmt.Printf("Type of age : %T\n", age)
	fmt.Printf("Type of email : %T\n", email)
	fmt.Println()

	// printing type using reflect package
	fmt.Println("Variable types using reflect")
	fmt.Println("=============================")
	creditLimit := 1234
	varType := reflect.TypeOf(creditLimit).String()
	fmt.Printf("%T\n", varType)
	fmt.Println()

	var customerNumber int
	var buyPrice float32
	var stockAvailable bool

	// printing default values
	fmt.Println("Variable default values")
	fmt.Println("=============================")
	fmt.Println("Integer : ", customerNumber)
	fmt.Println("Float:", buyPrice)
	fmt.Println("Boolean:", stockAvailable)
	fmt.Println()

	// Integers and Floats
	fmt.Println("Integers and Floats")
	fmt.Println("===================")
	var x int = 123
	var y uint = 234234
	a, b := 20.34, 34.56
	fmt.Println("Integers x and y :", x, y)
	fmt.Println("Floats a and b", a, b)
	c := a - b
	fmt.Printf("Result is %f\n", c)
	fmt.Printf("Type of c is %T\n", c)
	fmt.Println()
}
