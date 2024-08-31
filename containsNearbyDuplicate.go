package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	// 使用 map 來存儲元素值及其索引
	numIndexMap := make(map[int]int)

	// 遍歷數組
	for i, num := range nums {
		// 如果數組中的元素已經在 map 中，且索引差值小於等於 k，返回 true
		if index, found := numIndexMap[num]; found && i-index <= k {
			return true
		}
		// 更新當前元素的索引值
		numIndexMap[num] = i
	}

	// 如果遍歷完數組沒有找到符合條件的元素，返回 false
	return false
}

func main() {
	// 測試範例
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)) // false
}
