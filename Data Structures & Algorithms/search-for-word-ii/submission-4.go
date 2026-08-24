type TrieNode struct {
	children [26]*TrieNode
	word string
}

func (this *TrieNode) addWord(word string) {
	cur := this
	for _, c := range word {
		i := c - 'a'
		if cur.children[i] == nil {
			cur.children[i] = &TrieNode{}
		}
		cur = cur.children[i]
	}
	cur.word = word
	return
}


func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, word := range words { root.addWord(word) }

    rows, cols := len(board), len(board[0])
	res := []string{}

	var backtrack func(int, int, *TrieNode)
	backtrack = func(r, c int, node *TrieNode) {
		if r == rows || c == cols || r < 0 || c < 0 || board[r][c] == '#' { return }

		node = node.children[board[r][c]-'a']
		if node == nil { return }
		if node.word  != "" { 
			res = append(res, node.word)
			node.word = ""
		}

		temp := board[r][c]
		board[r][c] = '#'

		backtrack(r + 1, c    , node)
		backtrack(r - 1, c    , node)
		backtrack(r    , c + 1, node)
		backtrack(r    , c - 1, node)
		
		board[r][c] = temp
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			backtrack(r, c, root)
		}
	}

	return res
}
