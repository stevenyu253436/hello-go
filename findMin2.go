func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// The minimum is in the right half
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// The minimum is in the left half including mid
			right = mid
		} else {
			// When nums[mid] == nums[right], we cannot determine which half the minimum is in, so we reduce the search space by shrinking the right boundary
			right--
		}
	}

	return nums[left]
}