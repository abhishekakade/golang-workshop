// Exercise 2.11 break and continue

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// empty for loop is infinite loop so use with break conditions only
	for {

		number := rand.Intn(8)

		if number%3 == 0 {
			fmt.Println("Skipped: ", number)
			continue
			// continue is used to skip current iteration
		} else if number%2 == 0 {
			fmt.Println("Stopped at: ", number)
			break
			// break is used to break out of the program like when there are some errors
		}
		fmt.Println(number)
	}
}
