package main

import (
	"errors"
	"fmt"
	"os"
)

// div floor-divides numerator by denominator
func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

// MyFunc simulates named and optional params
func MyFunc(opts MyFuncOpts) {
	// do stuff
	fmt.Println(opts.FirstName, opts.LastName, opts.Age)
}

// addTo adds a base number to a variable number of params and returns result
// as a slice of int
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// divAndRemainder returns the quotient and remainder of numerator divided by
// denominator. Example of named return values and multiple return values.
func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func main() {
	result := div(5, 2)
	fmt.Println(result)

	// simulated named and optional param function call
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{FirstName: "Joe", LastName: "Smith"})

	// variadic function calls
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

}
