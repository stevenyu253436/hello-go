package main

import "fmt"

// canJump checks if you can reach the last index of the array.
func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return true
}

func main() {
	// Test case 1
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums1)) // Output: true

	// Test case 2
	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums2)) // Output: false
}
