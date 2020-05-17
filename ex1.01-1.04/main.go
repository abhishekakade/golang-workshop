// Exercise 1.02 to 1.04
package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		firstName     string = "Bob"
		lastName      string = "Smith"
		age           int8   = 34
		peanutAllergy bool   = false
		// for exercise 1.03
		Debug       bool      = false
		LogLevel    string    = "info"
		startUpTime time.Time = time.Now()
	)
	// multiple variables can also be declared in a single line as below
	// var firstName, lastName, age, peanutAllergy = "Bob", "Smith", 34, false

	// type declaration of variables can be skipped

	// variables declared without an explicit initial value are given their zero values:-
	// `0` for all integer types
	// `0.0` for floating point numbers
	// `false` for booleans
	// `""` for strings
	// `nil` for interfaces, slices, channels, maps, pointers and functions

	// `Debug bool` without an explicit initial value will be `false` (zero value) by default
	// var (
	// 	Debug bool
	// 	LogLevel = "info"
	// 	startUpTime = time.Now()
	// )

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(peanutAllergy)
	fmt.Println("Exercise 1.03")
	fmt.Println(Debug, LogLevel, startUpTime)
}
