package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize the slow pointer
	slow := 0

	// Fast pointer iterates through the array
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	// Return the length of the unique elements
	return slow + 1
}

func main() {
	// Example 1
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Printf("Output: %d, nums = %v\n", k1, nums1[:k1])

	// Example 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Printf("Output: %d, nums = %v\n", k2, nums2[:k2])
}
