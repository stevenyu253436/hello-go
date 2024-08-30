package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// 計算目前容器的面積
		width := right - left
		h := min(height[left], height[right])
		currentArea := width * h

		// 更新最大面積
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// 移動較矮的那根柱子，嘗試找到更大的容器
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// 輔助函數，用來返回兩個數字中的較小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 測試範例
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("最大容積:", maxArea(height)) // Output: 49
}
