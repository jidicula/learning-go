package main

import "fmt"

func main() {
	var x = []int{10, 20, 30}
	fmt.Println(x)
	// sparse slice
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)

	var z []int
	fmt.Println(z == nil)
	z = append(z, 10)
	fmt.Println(z == nil)
	fmt.Println(z)

	y = append(y, 4, 5, 6)
	fmt.Println(y)

	x = append(x, y...)
	fmt.Println(x)

	capacity_example()
	make_slice()
}

// capacity_example shows the difference between length and capacity
func capacity_example() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

// make_slice demonstrates the make keyword for creating a slice
func make_slice() {
	x := make([]int, 5)
	fmt.Println(x)

	// make accepts capacity as an arg too
	y := make([]int, 5, 10)
	fmt.Println(y, len(y), cap(y))
	// can create a slice with length zero but nonzero capacity
	z := make([]int, 0, 10)
	fmt.Println(z, len(z), cap(z))
	// we can append to this zero-length slice
	z = append(z, 5, 6, 7, 8)
	fmt.Println(z, len(z), cap(z))
}
