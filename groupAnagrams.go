package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		// 將字符串的字符排序
		sortedStr := sortString(str)
		// 將排序後的字符串作為鍵，將原始字符串加入對應的列表中
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	// 將結果從 map 中提取出來
	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

// 輔助函數：對字符串進行排序
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	// 測試範例
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs)) // [["bat"],["nat","tan"],["ate","eat","tea"]]
}
