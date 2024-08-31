package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	// 定義一個 map 對應開括號和閉括號的匹配關係
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		// 如果是閉括號
		if match, ok := bracketMap[char]; ok {
			// 檢查堆疊是否為空或頂部元素是否與當前閉括號匹配
			if len(stack) == 0 || stack[len(stack)-1] != match {
				return false
			}
			// 匹配成功則彈出堆疊頂部元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是開括號，推入堆疊
			stack = append(stack, char)
		}
	}

	// 最後堆疊應該為空，如果不為空，說明有未匹配的開括號
	return len(stack) == 0
}

func main() {
	// 測試用例
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
