package main

type TrieNode struct {
	children map[byte]*TrieNode
	word     string
}

type Trie struct {
	root *TrieNode
}

func Constructor() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[byte]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if node.children[char] == nil {
			node.children[char] = &TrieNode{
				children: make(map[byte]*TrieNode),
			}
		}
		node = node.children[char]
	}
	node.word = word
}

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	var result []string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, trie.root, &result)
		}
	}

	return result
}

func dfs(board [][]byte, i, j int, node *TrieNode, result *[]string) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == '#' {
		return
	}

	char := board[i][j]
	if node.children[char] == nil {
		return
	}

	node = node.children[char]
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = "" // 避免重复添加
	}

	board[i][j] = '#' // 标记当前字符为已访问
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range directions {
		dfs(board, i+dir[0], j+dir[1], node, result)
	}
	board[i][j] = char // 恢复当前字符
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	result := findWords(board, words)
	for _, word := range result {
		println(word)
	}
}
