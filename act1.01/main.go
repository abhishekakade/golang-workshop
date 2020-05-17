// will print random number of stars * (between 1 - 5) to the console

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// with  + 1, the number will be between 1 to 5 instead of 0 to 4
	randomNum := rand.Intn(5) + 1
	stars := strings.Repeat("*", randomNum)
	fmt.Println(stars)
}
