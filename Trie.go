package main

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (this *Trie) Insert(word string) {
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

func (this *Trie) Search(word string) bool {
	currentNode := this.root
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this.root
	for _, char := range prefix {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	println(trie.Search("apple"))   // 输出 true
	println(trie.Search("app"))     // 输出 false
	println(trie.StartsWith("app")) // 输出 true
	trie.Insert("app")
	println(trie.Search("app")) // 输出 true
}
