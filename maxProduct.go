func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize the result, maxProd, and minProd with the first element.
	maxProd, minProd, result := nums[0], nums[0], nums[0]

	// Traverse the array starting from the second element
	for i := 1; i < len(nums); i++ {
		// If the current number is negative, swap maxProd and minProd
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}

		// Update maxProd and minProd
		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])

		// Update the result with the maximum product found so far
		result = max(result, maxProd)
	}

	return result
}

// Utility functions to get max and min of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
