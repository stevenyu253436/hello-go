package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			// 取得堆疊中的最後兩個操作數
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var res int
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}

			// 將結果推回堆疊
			stack = append(stack, res)
		default:
			// 將操作數轉為整數並推入堆疊
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	// 最終堆疊中應該只有一個元素，即為結果
	return stack[0]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	result := evalRPN(tokens)
	fmt.Println(result) // Output: 9
}
