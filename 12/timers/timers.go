package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create context
	ctx := context.Background()
	// create parent context with timeout, along with cancel function
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	child, cancel2 := context.WithTimeout(parent, 3*time.Second)
	defer cancel2()
	start := time.Now()
	<-child.Done()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
