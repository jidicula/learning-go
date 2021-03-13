package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4} // original slice
	y := make([]int, 4)    // allocate space for slice
	fmt.Println("y:", y)

	num := copy(y, x) // copies x into y
	fmt.Println(y, num)
	z := make([]int, 2)
	copy(z, x[2:]) // copy and discard result

	num = copy(x[:3], x[1:])
	fmt.Println(x, num)
	fmt.Println()
	arrayCopy()
}

// arrayCopy copies a slice of an array
func arrayCopy() {
	x := []int{1, 2, 3, 4, 5}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])
	fmt.Println(y)
	copy(d[:], x) // copies up to capacity of array
	fmt.Println(d)
}
