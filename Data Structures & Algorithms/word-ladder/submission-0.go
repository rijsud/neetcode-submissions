func similar(w1 string, w2 string) bool {
	res := 1
	for i := range w1 {
		if w2[i] != w1[i] {res--}
	}
	return res == 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	res := 1
    hashMap := map[string][]string{}
	q := []string{}
	visit := map[string]bool{}
	
	for _, word := range wordList {
		hashMap[word] = []string{}
		if similar(beginWord, word) {q = append(q, word)}
		for _, word2 := range wordList {
			if similar(word, word2) {
				hashMap[word] = append(hashMap[word], word2)
			}
		}
	}

	for len(q) > 0 {
		qLen := len(q)
		res++
		if res > len(wordList) + 1 {return 0}
		for i := 0; i < qLen; i++ {
			checkWord := q[0]
			if checkWord == endWord {return res}
			q = q[1:]
			visit[checkWord] = true
			for _, word := range hashMap[checkWord] {
				if !visit[word] {q = append(q, word)}
			}
		}
	}
	return 0
}
