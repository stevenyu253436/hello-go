package main

import (
	"fmt"
)

// findSubstring function to solve the problem
func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	totalLen := wordLen * len(words)
	wordCount := make(map[string]int)
	result := []int{}

	// Populate the wordCount map
	for _, word := range words {
		wordCount[word]++
	}

	// Sliding window over the string `s`
	for i := 0; i <= len(s)-totalLen; i++ {
		seen := make(map[string]int)
		j := 0

		// Check every word in the current window
		for j < len(words) {
			start := i + j*wordLen
			subStr := s[start : start+wordLen]

			// If the word is not in the words list, break
			if _, exists := wordCount[subStr]; !exists {
				break
			}

			seen[subStr]++

			// If the word is seen more times than it appears in `words`, break
			if seen[subStr] > wordCount[subStr] {
				break
			}

			j++
		}

		// If all words are matched
		if j == len(words) {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	// Test the function with an example
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	result := findSubstring(s, words)
	fmt.Println("Output:", result) // Expected output: [0, 9]
}
