package main

import "fmt"

func main() {
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	a = append(a, 30, 40, 50, 60, 70, 80, 90, 100)
	a = append(a, 111)
	fmt.Println("Slice A after appending data:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	var cities1 = []string{"Goa", "Bangalore", "Mangalore"}
	var cities2 = []string{"Kochi", "Trivandrum"}

	cities2 = append(cities2, cities1...)
	fmt.Println(cities2)
}
