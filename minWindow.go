package main

import "fmt"

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	targetCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		targetCount[t[i]]++
	}

	left, right := 0, 0
	required := len(targetCount)
	formed := 0
	windowCounts := make(map[byte]int)
	minLen := len(s) + 1
	minLeft := 0

	for right < len(s) {
		char := s[right]
		windowCounts[char]++

		if targetCount[char] > 0 && windowCounts[char] == targetCount[char] {
			formed++
		}

		for left <= right && formed == required {
			char = s[left]

			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
			}

			windowCounts[char]--
			if targetCount[char] > 0 && windowCounts[char] < targetCount[char] {
				formed--
			}
			left++
		}

		right++
	}

	if minLen > len(s) {
		return ""
	}

	return s[minLeft : minLeft+minLen]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	result := minWindow(s, t)
	fmt.Println(result) // 應該輸出 "BANC"
}
