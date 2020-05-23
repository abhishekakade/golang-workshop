package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("Input cannot be negative!")
	} else if input > 100 {
		return errors.New("Input cannot be greater than 100!")
	} else if input%7 == 0 {
		return errors.New("Input cannot be divisible by 7!")
	} else {
		return nil
	}
}

func main() {
	input := 28

	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}
}