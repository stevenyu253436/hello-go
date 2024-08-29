package main

import (
    "fmt"
)

// twoSum returns the indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if idx, found := numMap[complement]; found {
            return []int{idx, i}
        }
        numMap[num] = i
    }

    return nil // According to the problem statement, there will always be exactly one solution
}

func main() {
    // Test cases
    nums1 := []int{2, 7, 11, 15}
    target1 := 9
    fmt.Println("Input:", nums1, "Target:", target1)
    fmt.Println("Output:", twoSum(nums1, target1))

    nums2 := []int{3, 2, 4}
    target2 := 6
    fmt.Println("Input:", nums2, "Target:", target2)
    fmt.Println("Output:", twoSum(nums2, target2))

    nums3 := []int{3, 3}
    target3 := 6
    fmt.Println("Input:", nums3, "Target:", target3)
    fmt.Println("Output:", twoSum(nums3, target3))
}
