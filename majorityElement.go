package main

import "fmt"

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	// Phase 1: Find the candidate
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	// Since the problem guarantees that the majority element always exists,
	// we can return the candidate directly.
	return candidate
}

func main() {
	nums1 := []int{3, 2, 3}
	fmt.Println(majorityElement(nums1)) // Output: 3

	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums2)) // Output: 2
}
