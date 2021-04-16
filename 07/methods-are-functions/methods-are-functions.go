package main

import "fmt"

type Adder struct {
	start int
}

// AddTo adds a value to an Adder.
func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))
	f1 := myAdder.AddTo
	fmt.Println(f1(10))
	// method expression
	f2 := Adder.AddTo
	val := f2(myAdder, 15)
	fmt.Println(val)
}
