package main

import "fmt"

func div60(i int) {
	// must call recover in defer because upon panic, only deferred funcs are run
	defer func() { // register function to handle panic
		if v := recover(); v != nil { // call recover() to see if value found
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
