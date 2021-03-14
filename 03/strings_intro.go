package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(b) // prints byte value
	fmt.Printf("%c\n", b)
	fmt.Printf("%c\n", rune(b))

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Printf("%s\n", s2)
	fmt.Printf("%s\n", s3)
	fmt.Printf("%s\n", s4)

}
