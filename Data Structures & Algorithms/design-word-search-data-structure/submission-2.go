type TrieNode struct {
	children [26]*TrieNode
	endOfWord bool
}

type WordDictionary struct {
   root *TrieNode 
}

func Constructor() WordDictionary {
    return WordDictionary{root : &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string)  {
	cur := this.root
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

func (this *WordDictionary) dfs(word string, root *TrieNode) bool {
	cur := root
	for idx, c := range word {
		if c == '.' {
			for _, child := range cur.children {
				if child != nil && this.dfs(word[idx+1:], child) {return true}
			}
			return false
		} else {
			i := c - 'a'
			if cur.children[i] == nil {return false}
			cur = cur.children[i]
		}
	}
	return cur.endOfWord
}


func (this *WordDictionary) Search(word string) bool {
    return this.dfs(word, this.root)
}
