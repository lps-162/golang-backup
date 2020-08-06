package main

import "fmt"

// for is used to loop thru strings, arrays, slices and maps too
// Golang has no while or do-while, which can be acheived with for itself

func main() {

	fmt.Println("Simple For Loop")
	fmt.Println("===============")
	for i := 0; i < 5; i++ {
		fmt.Println("Bangalore")
	}
	fmt.Println()

	fmt.Println("Skipping initilaization")
	fmt.Println("=======================")
	i := 0
	for ; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Println("Infinite loop, like a while")
	fmt.Println("=======================")
	for {
		fmt.Print(i, " ")
		if i == 20 {
			break
		}
		i++
	}
	fmt.Println()

	fmt.Println("Just with condition")
	fmt.Println("===================")
	i = 97
	for i < 123 {
		fmt.Print(string(i) + " ")
		i++

	}
	fmt.Println()

	fmt.Println("Looping thru strings")
	fmt.Println("====================")
	cloud := "GCloud"
	for range cloud {
		fmt.Println(cloud)
	}
	fmt.Println()
}
