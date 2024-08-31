package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 創建兩個長度為26的計數數組
	countS := make([]int, 26)
	countT := make([]int, 26)

	for i := 0; i < len(s); i++ {
		countS[s[i]-'a']++
		countT[t[i]-'a']++
	}

	// 比較兩個計數數組是否相等
	for i := 0; i < 26; i++ {
		if countS[i] != countT[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
}
