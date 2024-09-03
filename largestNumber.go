package main

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	// Convert the integers to strings
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = fmt.Sprintf("%d", num)
	}

	// Custom sort based on concatenated values
	sort.Slice(numStrs, func(i, j int) bool {
		return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
	})

	// Join the sorted strings
	result := strings.Join(numStrs, "")

	// Handle the case where the largest number is "0"
	if result[0] == '0' {
		return "0"
	}

	return result
}

func main() {
	// Example test cases
	fmt.Println(largestNumber([]int{10, 2}))           // Output: "210"
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9})) // Output: "9534330"
	fmt.Println(largestNumber([]int{0, 0}))            // Output: "0"
}
