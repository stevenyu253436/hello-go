package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string

	var backtrack func(current string, open int, close int)
	backtrack = func(current string, open int, close int) {
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		if open < n {
			backtrack(current+"(", open+1, close)
		}
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n)) // 輸出: ["((()))","(()())","(())()","()(())","()()()"]
}
