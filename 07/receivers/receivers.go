package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// Increment increases a Counter.
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// String returns a string representation of a Counter.
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	fmt.Printf("-----\n")
	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())
}

// doUpdateWrong does the update wrong by modifying a copy of the passed value.
func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

// doUpdateRight does the update right by modifying the passed pointer.
func doUpdateRight(c *Counter) {
	c.Increment()
	// Note that calling String() on the pointer is valid.
	fmt.Println("in doUpdateRight:", c.String())
}
