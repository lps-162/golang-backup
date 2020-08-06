package main

import "fmt"

func main() {
	var city string = "Bangalore"

	fmt.Println("=============")
	fmt.Println("Value :", city)
	fmt.Printf("Type %T\n", city)
	fmt.Printf("Length :%d\n", len(city))
	fmt.Println()

	fmt.Println("=============")
	str := "Shivan's Boutique"
	fmt.Println("Value :", str)
	fmt.Printf("Type %T\n", str)
	fmt.Printf("Length :%d\n", len(str))
	fmt.Println()

	str1 := "Silk Sarees"
	str2 := "Salwar Materials"
	fmt.Println("String Equality")
	fmt.Println("===============")
	result := str1 == str2
	fmt.Println(result)
	fmt.Printf("Type of result : %T\n", result)
	fmt.Println()
}
