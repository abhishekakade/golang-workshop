package main

import (
	"fmt"
)

func main() {
	unsortedNums := []int{9, 1, 0, 4, 2, 8, 3, 7, 5, 6}
	fmt.Println("Unsorted: ", unsortedNums)
	bubbleSort(unsortedNums)
}

func bubbleSort(nums []int) {

	// In for loop, in the `initial statement`, define a `Boolean` with the initial value of `true`
	// In the condition, check that Boolean. Leave the post statement empty
	for swapped := true; swapped; {

		// Set the Boolean variable to `false`
		swapped = false

		// Create a nested `for i` loop that steps over the whole slice of int values
		// Start the loop from the second element
		for i := 1; i < len(nums); i++ {

			// if previous number > next number
			if nums[i-1] > nums[i] {

				// next number, previous number = previous number, next number (value swapping)
				nums[i], nums[i-1] = nums[i-1], nums[i]
				swapped = true
			}
		}
	}

	fmt.Println("Sorted:   ", nums)
}
