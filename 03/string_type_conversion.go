package main

import "fmt"

func main() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(s)
	fmt.Println(s2)
}
