package main

import "fmt"

// 判斷字符串 s 是否為 t 的子序列
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0

	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == m
}

func main() {
	// 示例測試用例
	fmt.Println(isSubsequence("abc", "ahbgdc")) // 輸出: true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // 輸出: false
	fmt.Println(isSubsequence("ace", "abcde"))  // 輸出: true
}
