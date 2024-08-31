package main

import "fmt"

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		// Update the farthest index we can reach
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		// If we have reached the end of the range for the current jump
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums)) // Output: 2
}
