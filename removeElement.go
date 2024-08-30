package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement(nums, val)
	fmt.Printf("Output: %d, nums = %v\n", k, nums[:k])

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	k = removeElement(nums, val)
	fmt.Printf("Output: %d, nums = %v\n", k, nums[:k])
}
