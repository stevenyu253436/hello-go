package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func main() {
	// 测试用例
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))                   // 输出: true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))                   // 输出: false
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // 输出: true
}
