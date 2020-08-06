package main

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func cleanup() {
	fmt.Println("Cleanup")
}

// Defers often used with paired operations like
// open and close, connect and disconnect, and
// lock and unlock

// - It keeps our Close call near our Open call so it's easier to understand.
// - If our function had multiple return statements
//   (perhaps one in an if and one in an else), Close will happen before both of them.
// - Deferred Functions are run even if a runtime panic occurs.
// - Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
func main() {
	defer cleanup()
	first()
	second()
	fmt.Println("Third")
}
