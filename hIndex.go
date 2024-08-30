package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations) // Sort the citations in non-decreasing order

	n := len(citations)
	h := 0

	// Iterate through the sorted list from the end to find the H-index
	for i := 0; i < n; i++ {
		if citations[i] >= n-i {
			h = n - i
			break
		}
	}

	return h
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations)) // Output: 3
}
