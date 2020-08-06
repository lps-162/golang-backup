package main

import "fmt"

func main() {

	city := "Bangalore"

	fmt.Println("Length of String :", len(city))
	fmt.Println()

	firstName := "Rock"
	lastName := "Dwayne"
	fullName := firstName + " " + lastName

	fmt.Println("Concatenated string :", fullName)
	fmt.Println()

	city1 := "Goa"
	city2 := "Bangalore"

	compareStrings(city2, city1)
	fmt.Println()

	s3 := "Whatever"
	s4 := "Whatever"
	compareStrings(s3, s4)
	compareStrings(s3, s3)

	s5 := s3

	compareStrings(s5, s4)

	bothCities := fmt.Sprintf("%s %s", city1, city2)
	fmt.Println(bothCities)
	fmt.Println()

	var mife string
	fmt.Println("Default value of String is : ", mife)
}

func compareStrings(s1 string, s2 string) {
	if s1 == s2 {
		fmt.Println("The strings", s1, "and", s2, "are equal")
	} else if s1 > s2 {
		fmt.Println(s1, "is greater than", s2)
	} else {
		fmt.Println(s2, "is greater than", s1)
	}

}
