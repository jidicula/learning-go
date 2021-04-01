package main

import "fmt"

func main() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(s)
	fmt.Println(s2)
	stringToSlice()
}

// stringToSlice
func stringToSlice() {
	var s string = "Hello, 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
