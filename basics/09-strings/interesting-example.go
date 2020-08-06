package main

import (
	"fmt"
	"time"
)

func main() {
	beachCity := "goa"

	// Length is 13 times 5 which is 65.
	randString := "cat0123456789" +
		"cat0123456789" +
		"cat0123456789" +
		"cat0123456789" +
		"cat0123456789"

	fmt.Println("Length of goa :", len(beachCity))
	fmt.Println("Length of randString :", len(randString))
	fmt.Println()

	t0 := time.Now()

	for i := 0; i < 1000000000; i++ {
		result := len(beachCity)
		if result != 3 {
			return
		}
	}

	t1 := time.Now()

	for i := 0; i < 1000000000; i++ {
		result := len(randString)
		if result != 65 {
			return
		}
	}

	t2 := time.Now()

	fmt.Println(t1.Sub(t0))
	fmt.Println(t2.Sub(t1))
}
