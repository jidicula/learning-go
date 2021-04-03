package main

import "fmt"

func main() {
	// C-style for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Condition-only for, like `while` in other languages
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	// FizzBuzz with if statements and continue
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			// iterates over runes, not bytes
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	// labelling for loops
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()

	}

	for i, v := range evenVals[1:] {
		fmt.Println(i, v)
	}

}
