package main

import "fmt"

func main() {
	// Slice expressions
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// overlapping storage
	x[1] = 20
	y[0] = 10
	z[1] = 30
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// appending
	appendOverlappingSlice()
	fmt.Println()
	confusingOverlappingSlice()
	fmt.Println()
	fullSliceExpression()
}

// appendOverlappingSlice example of confusing append()
func appendOverlappingSlice() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 30)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

// confusingOverlappingSlice Example 3-7
func confusingOverlappingSlice() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	// fills up 3, 4 onwards in original slice
	y = append(y, 30, 40, 50)
	// appends 60 to original
	x = append(x, 60)
	// replaces 60 with 70
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

// fullSliceExpression modifies confusingOverlappingSlice() to use full slice
// expressions.
func fullSliceExpression() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]  // creates new slice upon append
	z := x[2:4:4] // creates new slice upon append
	fmt.Println(cap(x), cap(y), cap(z))
	// appends 30, 40, 50 to new slice
	y = append(y, 30, 40, 50)
	// appends 60 to original
	x = append(x, 60)
	// appends 70 to new slice
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
