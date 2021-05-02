package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v    // write v into ch1, blocks until ch1 is read
		v2 := <-ch2 // read from ch2, blocked until ch2 is ready
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v: // write v into ch2
	case v2 = <-ch1: // ch1 is read here
	}
	fmt.Println(v, v2)
}
