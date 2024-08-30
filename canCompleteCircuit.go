package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	tank := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		tank += gas[i] - cost[i]

		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start
}

func main() {
	// Test example 1
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println("Output:", canCompleteCircuit(gas, cost)) // Expected Output: 3

	// Test example 2
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	fmt.Println("Output:", canCompleteCircuit(gas, cost)) // Expected Output: -1
}
