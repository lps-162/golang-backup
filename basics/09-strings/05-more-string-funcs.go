package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	fmt.Println("More String Functions")
	fmt.Println("=====================")

	words := []string{"Bangalore", "is", "a", "party", "city", "too"}
	fmt.Println(words)

	for _, val := range words {
		fmt.Print(val + " ")
	}
	fmt.Println()

	var sentence string
	sentence = strings.Join(words, " ")
	fmt.Println("Full Sentence : ", sentence)
	sentence = strings.Join(words, "-")
	fmt.Println("Full Sentence : ", sentence)
	fmt.Println()

	itCity := "Bangalore"
	fmt.Println("Repeat :", strings.Repeat(itCity, 5))
	fmt.Println("Replace :", strings.Replace(itCity, "a", "z", -1))
	fmt.Println()

	fmt.Println("Split Example")
	splitSlice := strings.Split(sentence, "-")
	fmt.Println(splitSlice, reflect.TypeOf(splitSlice))
	for _, val := range splitSlice {
		fmt.Printf(val + ", ")
	}
	fmt.Println()

	fmt.Println("Other functions")
	fmt.Println("Title : ", strings.Title("bangalore and goa"))
	fmt.Println("ToTitle : ", strings.ToTitle("bangalore and goa"))
	fmt.Println("ToLower : ", strings.ToLower("BANGALORE and goa"))
	fmt.Println("ToUpper : ", strings.ToUpper("bangalore and GOA"))

	fmt.Println()

	fmt.Println("Trim Functions")
	fmt.Println("Trim : ", strings.Trim("BangaloreBang", "Bang"))
	fmt.Println("TrimLeft :", strings.TrimLeft("BangaloreBang", "Bang"))
	fmt.Println("TrimRight :", strings.TrimRight("BangaloreBang", "Bang"))
	fmt.Println("Trim : ", strings.TrimLeft("BangBangBangaloreBang", "Bang"))
	fmt.Println("TrimSpace : ", strings.TrimSpace("    Bangalore locked down   "))
}
