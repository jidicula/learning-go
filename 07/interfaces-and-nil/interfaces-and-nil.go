package main

import "fmt"

func main() {
	// var s *string
	// fmt.Printf("%v\n", s == nil) // prints true
	var i interface{}
	fmt.Printf("%v\n", i == nil) //  prints true
	// i = s
	fmt.Printf("%v\n", i == nil) //  prints false
}
