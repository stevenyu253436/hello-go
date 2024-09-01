func maxSubArray(nums []int) int {
	// 初始化当前最大子数组和和全局最大子数组和
	currentMax := nums[0]
	globalMax := nums[0]

	// 从第二个元素开始遍历数组
	for i := 1; i < len(nums); i++ {
		// 更新当前最大子数组和
		currentMax = max(nums[i], currentMax+nums[i])
		// 更新全局最大子数组和
		if currentMax > globalMax {
			globalMax = currentMax
		}
	}

	return globalMax
}

// 辅助函数，用来返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
