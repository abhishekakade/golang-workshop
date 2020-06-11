package main

import "fmt"

func triggerNumberWraparound() {
	// Declare an int8 variable with an initial value of 125:
	var a int8 = 125
	// Declare an uint8 variable with an initial value of 253:
	var b uint8 = 253
	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}
}

// run with `go run` instead of building an executable
// since there are other `main` packages as well in this folder

func main() {
	triggerNumberWraparound()
}
