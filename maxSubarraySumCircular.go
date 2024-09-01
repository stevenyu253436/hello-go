func maxSubarraySumCircular(nums []int) int {
	totalSum := 0
	maxSum := nums[0]
	curMax := 0
	minSum := nums[0]
	curMin := 0

	for _, num := range nums {
		totalSum += num

		// Kadane's algorithm to find max subarray sum
		curMax = max(curMax+num, num)
		maxSum = max(maxSum, curMax)

		// Kadane's algorithm to find min subarray sum
		curMin = min(curMin+num, num)
		minSum = min(minSum, curMin)
	}

	// If all numbers are negative, maxSum will be the maximum number
	if maxSum > 0 {
		return max(maxSum, totalSum-minSum)
	} else {
		return maxSum
	}
}

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
