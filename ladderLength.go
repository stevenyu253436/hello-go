package main

import (
	"container/list"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 將 wordList 轉換為一個 set 以便快速查找
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	// 如果 endWord 不在 wordList 中，則無法達成轉換
	if !wordSet[endWord] {
		return 0
	}

	// 初始化 BFS 隊列
	queue := list.New()
	queue.PushBack(beginWord)
	level := 1

	// 開始 BFS
	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			// 取出當前層的單詞
			currentWord := queue.Remove(queue.Front()).(string)

			// 嘗試每個字母位置上的變化
			for j := 0; j < len(currentWord); j++ {
				for c := 'a'; c <= 'z'; c++ {
					// 生成新單詞
					newWord := currentWord[:j] + string(c) + currentWord[j+1:]

					// 如果新單詞是 endWord，返回當前層次 + 1
					if newWord == endWord {
						return level + 1
					}

					// 如果新單詞在 wordSet 中，加入到隊列中
					if wordSet[newWord] {
						queue.PushBack(newWord)
						delete(wordSet, newWord) // 避免重複使用
					}
				}
			}
		}
		level++
	}

	return 0
}

func main() {
	// 測試範例
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	result := ladderLength(beginWord, endWord, wordList)
	println(result) // 應該輸出 5
}
