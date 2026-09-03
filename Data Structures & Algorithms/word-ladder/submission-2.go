func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
    hashMap := map[string][]string{}
	
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			hashMap[pattern] = append(hashMap[pattern], word)
		}
	}

	visit := map[string]bool{beginWord : true}
	q := []string{beginWord}
	res := 1

	for len(q) > 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			word := q[0]
			if word == endWord {return res}
			q = q[1:]
			visit[word] = true
			for i := 0; i < len(word); i++ {
				pattern := word[:i] + "*" + word[i+1:]
				for _, neiWord := range hashMap[pattern] {
					if !visit[neiWord] {
						visit[neiWord] = true
						q = append(q, neiWord)
					}
				}
			}
		}
		res++
	}
	return 0
}
