package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Simple arrays")
	fmt.Println("-------------")
	cities := [5]string{"Goa", "Bangalore", "Trivandrum"}
	fmt.Println("Length :", len(cities))
	fmt.Println(cities)
	fmt.Println()

	fmt.Println("Array of Ints")
	fmt.Println("---------------")
	var primes [5]int = [5]int{2, 3, 5}
	fmt.Println("Primes :", primes)
	fmt.Println()

	fmt.Println("Assignment statements for Arrays")
	fmt.Println("--------------------------------")
	beerTypes := [4]string{"Carlsberg", "Kingfisher", "Bira"}
	fmt.Println(beerTypes)
	fmt.Println(reflect.TypeOf(beerTypes).Kind())
	beerTypes[3] = "Budweisser"
	fmt.Println(beerTypes)
	fmt.Println()

	fmt.Println("Using ellipsis operator")
	fmt.Println("-----------------------")
	bikeBrands := [...]string{"Royal Enfield", "Jawa"}
	fmt.Println(len(bikeBrands))
	fmt.Println()

	fmt.Println("Initializing specific elements")
	fmt.Println("------------------------------")
	distArray := [5]int{1: 500, 3: 250}
	fmt.Println(distArray)
	fmt.Println()

}
