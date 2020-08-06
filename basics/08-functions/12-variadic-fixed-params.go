package main

import "fmt"

func doOperation(opType string, numbers ...int) int {
	if opType == "sum" {
		opResult := 0
		for _, val := range numbers {
			opResult += val
		}
		return opResult
	} else if opType == "product" {
		opResult := 1
		for _, val := range numbers {
			opResult *= val
		}
		return opResult
	} else {
		panic("Wrong Operation Type")
	}
}

func main() {

	res1 := doOperation("sum", 1, 2, 3, 4, 5)
	res2 := doOperation("product", 1, 2, 3, 4)

	fmt.Println("Sum and Product values are :", res1, res2)
}
