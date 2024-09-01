package main

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	currentNode := this.root
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			currentNode.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchFromNode(word, this.root)
}

func (this *WordDictionary) searchFromNode(word string, node *TrieNode) bool {
	for i, char := range word {
		if char == '.' {
			// 如果是 '.', 遍历当前节点的所有孩子节点
			for _, child := range node.children {
				if this.searchFromNode(word[i+1:], child) {
					return true
				}
			}
			return false
		} else {
			// 如果不是 '.', 正常查找
			if _, exists := node.children[char]; !exists {
				return false
			}
			node = node.children[char]
		}
	}
	return node.isEnd
}

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	println(wordDictionary.Search("pad")) // 输出 false
	println(wordDictionary.Search("bad")) // 输出 true
	println(wordDictionary.Search(".ad")) // 输出 true
	println(wordDictionary.Search("b..")) // 输出 true
}
