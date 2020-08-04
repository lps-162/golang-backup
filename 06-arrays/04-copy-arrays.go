package main

import "fmt"

func main() {

	strArray1 := [3]string{"India", "USA", "Canada"}
	strArray2 := strArray1
	strArray3 := &strArray1

	fmt.Println("Before Modification")
	fmt.Println("--------------------")
	fmt.Println("First Array :", strArray1)
	fmt.Println("Second Array :", strArray2)
	fmt.Println("Third Array :", *strArray3)
	fmt.Println()

	strArray1[1] = "States"

	fmt.Println("After Modification")
	fmt.Println("--------------------")
	fmt.Println("First Array :", strArray1)
	fmt.Println("Second Array :", strArray2)
	fmt.Println("Third Array :", *strArray3)
	fmt.Println()

}
