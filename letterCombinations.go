package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// 定義數字到字母的映射
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string
	var current []byte

	// 定義回溯函數
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits) {
			result = append(result, string(current))
			return
		}

		letters := digitToLetters[digits[index]]
		for i := 0; i < len(letters); i++ {
			current = append(current, letters[i])
			backtrack(index + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}

func main() {
	digits := "23"
	combinations := letterCombinations(digits)
	fmt.Println(combinations) // 输出: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
}
