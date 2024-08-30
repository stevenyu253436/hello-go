package main

import "fmt"

// productExceptSelf returns an array such that answer[i] is equal to
// the product of all the elements of nums except nums[i].
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Step 1: Calculate the prefix products
	prefixProduct := 1
	for i := 0; i < n; i++ {
		result[i] = prefixProduct
		prefixProduct *= nums[i]
	}

	// Step 2: Calculate the suffix products and multiply with the prefix products
	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffixProduct
		suffixProduct *= nums[i]
	}

	return result
}

func main() {
	// Test case 1
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums1)) // Output: [24, 12, 8, 6]

	// Test case 2
	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums2)) // Output: [0, 0, 9, 0, 0]
}
