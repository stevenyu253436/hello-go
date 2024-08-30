package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// Update the minimum price if the current price is lower
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// Calculate the profit if selling at the current price
			profit := prices[i] - minPrice
			// Update maxProfit if the current profit is greater
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max Profit:", maxProfit(prices)) // Output: 5
}
