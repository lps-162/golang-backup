package main

import "fmt"

func main() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)

	randBytes := []byte{97, 98, 99, 100, 101}
	randString := string(randBytes)
	fmt.Println(randString)
}
