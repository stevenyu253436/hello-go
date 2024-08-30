package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	// 測試範例
	fmt.Println(strStr("sadbutsad", "sad"))  // Output: 0
	fmt.Println(strStr("leetcode", "leeto")) // Output: -1
}
