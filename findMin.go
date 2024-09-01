func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	// 如果数组没有旋转，直接返回第一个元素
	if nums[left] < nums[right] {
		return nums[left]
	}

	// 二分查找
	for left < right {
		mid := left + (right-left)/2

		// 判断中间元素和右边界
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
