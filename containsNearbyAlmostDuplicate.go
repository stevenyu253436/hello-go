package main

import (
	"fmt"
	"math"
)

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if indexDiff < 1 || valueDiff < 0 {
		return false
	}

	getBucket := func(num, width int) int {
		if num >= 0 {
			return num / width
		}
		return (num+1)/width - 1
	}

	buckets := make(map[int]int)
	width := valueDiff + 1 // bucket size

	for i := 0; i < len(nums); i++ {
		bucket := getBucket(nums[i], width)

		// Check current bucket
		if _, exists := buckets[bucket]; exists {
			return true
		}

		// Check the neighboring buckets
		if num, exists := buckets[bucket-1]; exists && math.Abs(float64(nums[i]-num)) <= float64(valueDiff) {
			return true
		}
		if num, exists := buckets[bucket+1]; exists && math.Abs(float64(nums[i]-num)) <= float64(valueDiff) {
			return true
		}

		// Add current number to the appropriate bucket
		buckets[bucket] = nums[i]

		// Remove the oldest element outside the sliding window
		if i >= indexDiff {
			delete(buckets, getBucket(nums[i-indexDiff], width))
		}
	}

	return false
}

func main() {
	nums := []int{1, 5, 9, 1, 5, 9}
	indexDiff := 2
	valueDiff := 3
	fmt.Println(containsNearbyAlmostDuplicate(nums, indexDiff, valueDiff)) // Output: false
}
