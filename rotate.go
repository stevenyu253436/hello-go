package main

import "fmt"

// Function to rotate the array to the right by k steps
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // In case k is greater than the length of the array

	reverse(nums, 0, n-1) // Reverse the entire array
	reverse(nums, 0, k-1) // Reverse the first k elements
	reverse(nums, k, n-1) // Reverse the remaining elements
}

// Function to reverse the elements in the array between indices start and end
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// Main function to test the rotate function
func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	rotate(nums1, k1)
	fmt.Println(nums1) // Expected output: [5 6 7 1 2 3 4]

	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	rotate(nums2, k2)
	fmt.Println(nums2) // Expected output: [3 99 -1 -100]
}
