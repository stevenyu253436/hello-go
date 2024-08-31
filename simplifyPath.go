package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	// 分割路徑
	parts := strings.Split(path, "/")
	stack := []string{}

	for _, part := range parts {
		if part == "" || part == "." {
			// 空字符串或"."表示當前目錄，忽略
			continue
		} else if part == ".." {
			// ".."表示返回上一級目錄
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			// 將有效的目錄壓入棧
			stack = append(stack, part)
		}
	}

	// 組合簡化路徑
	return "/" + strings.Join(stack, "/")
}

func main() {
	path := "/home//foo/"
	fmt.Println("Simplified Path:", simplifyPath(path))
}
