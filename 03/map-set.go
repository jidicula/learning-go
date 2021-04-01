package main

import "fmt"

// Simulating a rudimentary set using a map with bool value type
// Only includes membership, no union, intersection, or subtraction
func main() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	// add items into set
	for _, v := range vals {
		intSet[v] = true
	}
	// len will differ since dupes are already in map
	fmt.Println(len(vals), len(intSet))
	// fetching value for key in dict returns true
	fmt.Println(intSet[5])
	// fetching value for key not in dict returns false
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}
