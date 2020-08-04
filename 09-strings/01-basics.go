package main

import (
	"fmt"
	"reflect"
)

func printBytes(s string) {
	fmt.Println("Individual Bytes")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Println("Individual Chars")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printCharsAndPosition(s string) {
	for index, elem := range s {
		fmt.Printf("%c starts at byte %d\n", elem, index)
	}
}

func loopThruChars(s string) {
	for i := 0; i < len(s); i++ {
		whatsthis := string(s[i])
		fmt.Println(whatsthis, reflect.TypeOf(whatsthis))
	}
}

func main() {
	var custName string = "Shivan LP"
	var custNum = "123123"
	city := "Bangalore"

	fmt.Println()
	fmt.Println("Customer Name : ", custName)
	fmt.Println("Customer Number :", custNum)
	fmt.Println("Customer City :", city)

	fmt.Printf("%s : %s : %s", custName, custNum, city)
	fmt.Println()

	fmt.Println("Length of city string :", len(city))
	fmt.Println()

	printBytes(city)
	fmt.Println()

	printChars(city)
	fmt.Println()

	fmt.Println()
	printCharsAndPosition(city)

	fmt.Println()
	loopThruChars(city)
}
