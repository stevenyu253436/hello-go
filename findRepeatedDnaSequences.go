package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	sequenceMap := make(map[string]int)
	var result []string

	// 只有當字符串長度大於 10 時，才有可能存在重複的 10 字符串
	if len(s) <= 10 {
		return result
	}

	// 遍歷字符串，提取每個長度為 10 的子字符串
	for i := 0; i <= len(s)-10; i++ {
		subStr := s[i : i+10]
		sequenceMap[subStr]++

		// 當某個子字符串的出現次數等於 2 時，加入結果
		// 只在第一次發現時加入，避免重複
		if sequenceMap[subStr] == 2 {
			result = append(result, subStr)
		}
	}

	return result
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	result := findRepeatedDnaSequences(s)
	fmt.Println(result)
}
