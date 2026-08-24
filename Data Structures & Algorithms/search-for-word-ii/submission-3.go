type TrieNode struct {
	children [26]*TrieNode
	endOfWord bool
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
	cur.endOfWord = true
	return
}


func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, word := range words { root.addWord(word) }
    rows, cols := len(board), len(board[0])
	wordSet := map[string]bool{}
	res := []string{}

	getIndex := func(c byte) int { return int(c - 'a') }

	var backtrack func(int, int, *TrieNode, string)
	backtrack = func(r, c int, node *TrieNode, word string) {
		if r == rows || c == cols || r < 0 || c < 0 || board[r][c] == '#' ||
		node.children[getIndex(board[r][c])] == nil  {return}

		temp := board[r][c]
		board[r][c] = '#'
		node = node.children[getIndex(temp)]
		word += string(temp)
		if node.endOfWord { wordSet[word] = true }

		backtrack(r + 1, c    , node, word)
		backtrack(r - 1, c    , node, word)
		backtrack(r    , c + 1, node, word)
		backtrack(r    , c - 1, node, word)
		
		board[r][c] = temp
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			backtrack(r, c, root, "")
		}
	}

	for word := range wordSet {
		res = append(res, word)
	}

	return res
}
