func foreignDictionary(words []string) string {
	res := []byte{}
    adj := map[byte]map[byte]struct{}{}
	cycle, seen := map[byte]bool{}, map[byte]bool{}

	for _, word := range words {
        for i := 0; i < len(word); i++ {
            char := word[i]
            if _, exists := adj[char]; !exists {
                adj[char] = make(map[byte]struct{})
            }
        }
    }

    for i := 0; i < len(words)-1; i++ {
        w1, w2 := words[i], words[i+1]
        minLen := len(w1)
        if len(w2) < minLen {
            minLen = len(w2)
        }
        if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
            return ""
        }
        for j := 0; j < minLen; j++ {
            if w1[j] != w2[j] {
                adj[w1[j]][w2[j]] = struct{}{}
                break
            }
        }
    }

	var dfs func(char byte) bool
	dfs = func(char byte) bool {
		if seen[char] {return cycle[char]}
		seen[char], cycle[char] = true, true

		for nextChar := range adj[char] {
			if dfs(nextChar) {return true}
		}

		cycle[char] = false
		res = append(res, char)
		return false
	}

	for char := range adj {
		if dfs(char) {return ""}
	}

	for i, j := 0, len(res) - 1; i < j; i, j = i + 1, j - 1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
