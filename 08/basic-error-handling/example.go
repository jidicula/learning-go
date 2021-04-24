package main

import (
	"errors"
	"fmt"
	"os"
)

// calcRemainderAndMod returns the quotient and remainder.
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		// set other return values to zero values
		return 0, 0, errors.New("denominator is 0") // don't capitalize err msg
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	numerator := 20
	denominator := 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
}
