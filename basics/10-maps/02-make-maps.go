package main

import "fmt"

func main() {

	var employees = make(map[string]int)
	employees["vijay"] = 80
	employees["rahul"] = 60
	employees["john"] = 75
	fmt.Println(employees)

	empList := make(map[string]int)
	empList["amritha"] = 80
	empList["swetha"] = 60
	fmt.Println(empList)

	stocks := make(map[string]int)

	fmt.Println("Length of slices")
	fmt.Println("----------------")
	fmt.Println(len(employees))
	fmt.Println(len(empList))
	fmt.Println(len(stocks))
	fmt.Println()

}
