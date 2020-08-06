package main

import "fmt"

// Go does not have exceptions
// Array out of bounds, nil pointer dereference - runtime errors
// When runtime error occurs, go panics and stops normal execution
// then all deferred func calls in that gorouting are executed
// and then program crashes with a log message

// panic either occurs automatically at runtime or can be called
// manually also thru panic(), takes any parameter type
// Panic should be used when impossible situations arise in your
// program

// During normal execution, a call to recover will return nil and have no other effect.
// If the current goroutine is panicking, a call to recover will capture the value
// given to panic and resume normal execution

func main() {
	var action int
	fmt.Println("Enter 1 for Customers, 2 for Products")
	fmt.Scanln(&action)

	defer func() {
		panickedOn := recover()
		fmt.Println("Recovering :", panickedOn)
	}()

	switch action {
	case 1:
		fmt.Println("Listing all Customers")
	case 2:
		fmt.Println("List of all Products")
	default:

		panic(fmt.Sprintf("Improper action type :%d", action))
	}

}
