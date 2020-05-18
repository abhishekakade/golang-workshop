package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "hello!", time.Now()
}

func main() {
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)

	// Declare our variables with an initial value
	query, limit, offset := "bat", 10, 0

	// Change each variable's values using a one-line statement
	query, limit, offset = "ball", offset, 20

	fmt.Println(query, limit, offset)
}
