package main

import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	// Initialize a slice to store the number of candies each child should get
	candies := make([]int, n)

	// Each child must have at least one candy
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// First pass: left to right, make sure right child has more candies if rating is higher
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Second pass: right to left, make sure left child has more candies if rating is higher
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// Sum up the candies to get the minimum number required
	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	ratings1 := []int{1, 0, 2}
	fmt.Println("Output for ratings1:", candy(ratings1)) // Expected output: 5

	// Example 2
	ratings2 := []int{1, 2, 2}
	fmt.Println("Output for ratings2:", candy(ratings2)) // Expected output: 4
}
