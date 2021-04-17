package main

import "fmt"

type Inner struct {
	A int
}

// IntPrinter prints an int.
func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

// Double doubles the Inner.
func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

// IntPrinter also prints an int.
func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func main() {
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	// Method on embedded field is in the method set of containing struct, so
	// this is legal.
	fmt.Println(o.Double())
}
