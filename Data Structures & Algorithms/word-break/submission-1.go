type TrieNode struct {
	children [26]*TrieNode
	endOfWord bool
}

func (this *TrieNode) addWord(word string) {
	cur := this
	for _, c := range word {
		i := c - 'a'
		if cur.children[i] == nil {	cur.children[i] = &TrieNode{} }
		cur = cur.children[i]
	}
	cur.endOfWord = true
}

func wordBreak(s string, wordDict []string) bool {
    root := &TrieNode{}
	for _, word := range wordDict {
		root.addWord(word)
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := range s {
		if !dp[i] {continue}

		cur := root
		for j := i; j < len(s); j++ {
			idx := s[j] - 'a'
			if cur.children[idx] == nil {break}
			cur = cur.children[idx]
			if cur.endOfWord {dp[j+1] = true}
		}
	}

	return dp[len(s)]
}
