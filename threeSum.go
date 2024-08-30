package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 初始化結果陣列
	result := [][]int{}

	// 先將陣列排序
	sort.Ints(nums)

	// 遍歷每個元素作為三元組的第一個數
	for i := 0; i < len(nums)-2; i++ {
		// 如果當前元素與前一個元素相同，則跳過以避免重複解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 初始化左右指針
		left, right := i+1, len(nums)-1

		// 使用雙指針法尋找剩餘的兩個數字
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// 找到一組符合條件的三元組，加入結果
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 跳過重複的左指針元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 跳過重複的右指針元素
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 移動指針
				left++
				right--
			} else if sum < 0 {
				// 如果和小於0，移動左指針使和變大
				left++
			} else {
				// 如果和大於0，移動右指針使和變小
				right--
			}
		}
	}

	return result
}

func main() {
	// 測試數據
	nums := []int{-1, 0, 1, 2, -1, -4}

	// 調用函數並打印結果
	result := threeSum(nums)
	fmt.Println(result)
}
