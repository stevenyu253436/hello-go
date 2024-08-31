package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	// 如果 pattern 和 words 長度不一致，則無法匹配
	if len(pattern) != len(words) {
		return false
	}

	// 建立兩個 map，分別存儲 pattern 到 word 和 word 到 pattern 的映射
	mapPatternToWord := make(map[byte]string)
	mapWordToPattern := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]

		// 檢查 pattern 到 word 的映射
		if val, exists := mapPatternToWord[p]; exists {
			if val != w {
				return false
			}
		} else {
			mapPatternToWord[p] = w
		}

		// 檢查 word 到 pattern 的映射
		if val, exists := mapWordToPattern[w]; exists {
			if val != p {
				return false
			}
		} else {
			mapWordToPattern[w] = p
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  // 输出: true
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // 输出: false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))  // 输出: false
	fmt.Println(wordPattern("abba", "dog dog dog dog"))  // 输出: false
}
