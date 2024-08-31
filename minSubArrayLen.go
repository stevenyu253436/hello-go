package main

import (
	"fmt"
)

// minSubArrayLen 函數，實現解題邏輯
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	sum := 0
	minLength := n + 1 // 初始化為比陣列長度大的值

	for right := 0; right < n; right++ {
		sum += nums[right]

		// 當目前的sum大於或等於target時，嘗試收縮左邊界
		for sum >= target {
			minLength = min(minLength, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	// 如果minLength未更新，表示沒有符合條件的子陣列
	if minLength == n+1 {
		return 0
	}

	return minLength
}

// 輔助函數，用來返回兩個整數中的較小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 測試用例
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	result := minSubArrayLen(target, nums)
	fmt.Printf("The minimal length of a subarray is: %d\n", result)
}
