package main

import (
	"fmt"
)

func main() {

	const (
		FIZZ     = 3
		BUZZ     = 5
		FIZZBUZZ = 15
	)

	// With for loop
	for i := 1; i < 101; i++ {
		// below line works fine but directly using 15 will save some calculations and make it faster
		// if i%FIZZ == 0 && i%BUZZ == 0 {
		// so we will use that as an optimization
		if i%FIZZBUZZ == 0 {
			fmt.Println("FizzBuzz")
		} else if i%FIZZ == 0 {
			fmt.Println("Fizz")
		} else if i%BUZZ == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// With switch case
	for i := 1; i < 101; i++ {
		switch {
		// below line works fine but directly using 15 will save some calculations and make it faster
		// case ((i%FIZZ == 0) && (i%BUZZ == 0)):
		// so we will use that as an optimization
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
