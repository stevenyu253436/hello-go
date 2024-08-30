package main

import "fmt"

// merge function to merge two sorted arrays
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Pointers for nums1, nums2, and the merged array
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	// Traverse from the end of nums1 and nums2, merging in reverse order
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// If there are any remaining elements in nums2, copy them over
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println("Merged array:", nums1) // Output: [1 2 2 3 5 6]

	// Example 2
	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge(nums1, m, nums2, n)
	fmt.Println("Merged array:", nums1) // Output: [1]

	// Example 3
	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	merge(nums1, m, nums2, n)
	fmt.Println("Merged array:", nums1) // Output: [1]
}
