package main

import (
	"fmt"
	"strings"
)

// 檢查字符是否為字母或數字
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

// 檢查字符串是否為回文
func isPalindrome(s string) bool {
	// 將字符串中的大寫字母轉換為小寫字母
	s = strings.ToLower(s)

	// 使用雙指針法來檢查是否為回文
	left, right := 0, len(s)-1
	for left < right {
		// 跳過非字母或數字的字符
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// 比較字符
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	// 測試範例
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // 輸出 true
	fmt.Println(isPalindrome("race a car"))                     // 輸出 false
	fmt.Println(isPalindrome(" "))                              // 輸出 true
}
