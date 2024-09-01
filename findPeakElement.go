func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[mid+1] {
			// 如果中间元素比右边元素大，则峰值在左侧或就是中间元素
			right = mid
		} else {
			// 否则峰值在右侧
			left = mid + 1
		}
	}

	// 最后 left 和 right 会收敛到峰值元素的位置
	return left
}
