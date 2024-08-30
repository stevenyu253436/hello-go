package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			// 因為題目要求返回的索引是1-indexed的，所以加1
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil // 這行理論上永遠不會被執行到，因為題目保證有解
}

func main() {
	// 你可以在這裡測試你的函數
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(numbers, target)
	fmt.Println(result) // 應輸出: [1, 2]
}
