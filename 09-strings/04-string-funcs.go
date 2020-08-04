package main

import (
	"fmt"
	"strings"
)

func main() {

	beachCity := "Goa"
	itCity := "Bangalore"
	var result int

	fmt.Println("Comparing strings")
	fmt.Println("=================")
	result = strings.Compare(beachCity, itCity)
	fmt.Println(result)

	result = strings.Compare(itCity, beachCity)
	fmt.Println(result)

	result = strings.Compare(itCity, itCity)
	fmt.Println(result)

	fmt.Println("strings Contains")
	fmt.Println("=================")
	fmt.Println(strings.Contains("Bangalore", "banga"))
	fmt.Println(strings.Contains("Bangalore", "Banga"))
	fmt.Println(strings.Contains(itCity, "Banga"))
	fmt.Println(strings.Contains(itCity, "Bangalore"))
	fmt.Println()

	fmt.Println("strings Count")
	fmt.Println("===================")
	fmt.Println(strings.Count("Bangalore", "a"))
	fmt.Println(strings.Count("Bangalore", "bang"))
	fmt.Println(strings.Count("Bangalore", "")) // returns length + 1
	fmt.Println(strings.Count("bang", itCity))

	fmt.Println("strings EqualFold") // case insensitive comparision
	fmt.Println("===================")
	fmt.Println(strings.Count(itCity, "bAnGalore"))

	fmt.Println("Splitting strings")
	fmt.Println("===================")
	sentence := "Bangalore has better climate but horrible traffic"
	fmt.Println(sentence)
	words := strings.Fields(sentence)
	fmt.Println(words)

	for _, v := range words {
		fmt.Println(v)
	}
	fmt.Println()

	fmt.Println("Index, HasPrefix and HasSuffix")
	fmt.Println("=============================")
	fmt.Println("HasPrefix : ", strings.HasPrefix(itCity, "Bang"))
	fmt.Println("HasPrefix : ", strings.HasPrefix(itCity, "goa"))

	fmt.Println("HasSuffix : ", strings.HasSuffix(itCity, "alore"))
	fmt.Println("HasSuffix : ", strings.HasPrefix(itCity, "al"))

	fmt.Println("Index ", strings.Index(itCity, "alore"))
	fmt.Println("Index ", strings.Index(itCity, "alorez"))
	fmt.Println("LastIndex ", strings.LastIndex(itCity, "a"))

	fmt.Println("IndexByte", strings.IndexByte(itCity, 'r'))
	fmt.Println("IndexByte ", strings.IndexByte(itCity, 'z'))

}
