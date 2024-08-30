package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var currentLine []string
	currentLen := 0

	for _, word := range words {
		if currentLen+len(currentLine)+len(word) > maxWidth {
			if len(currentLine) == 1 {
				// 如果只有一個單詞，不需要插入額外空格
				result = append(result, currentLine[0]+strings.Repeat(" ", maxWidth-len(currentLine[0])))
			} else {
				// 多個單詞，插入額外空格
				for i := 0; i < maxWidth-currentLen; i++ {
					currentLine[i%(len(currentLine)-1)] += " "
				}
				result = append(result, strings.Join(currentLine, ""))
			}
			currentLine = []string{}
			currentLen = 0
		}
		currentLine = append(currentLine, word)
		currentLen += len(word)
	}

	// 處理最後一行，該行需要左對齊
	lastLine := strings.Join(currentLine, " ")
	lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
	result = append(result, lastLine)

	return result
}

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	result := fullJustify(words, maxWidth)
	for _, line := range result {
		fmt.Println(line)
	}
}
