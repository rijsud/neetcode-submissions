func minWindow(s string, t string) string {
	if t == "" {return ""}
	
    countT := make(map[rune]int)
	for _, letter := range t {
		countT[letter]++
	}

	have, need := 0, len(countT)
	countS := make(map[rune]int)
	res, l, startIndex, endIndex := len(s) + 1, 0, 0, 0
	for r := 0; r < len(s); r++ {
		curLetter := rune(s[r])
		countS[curLetter]++
		if _, ok := countT[curLetter]; ok && countT[curLetter] == countS[curLetter] {
			have++
		}
		for have == need {
			if (r - l + 1) < res {
				res = r - l + 1
				startIndex = l
				endIndex = r
			}
			leftLetter := rune(s[l])
			countS[leftLetter]--
			if _, ok := countT[leftLetter]; ok && countS[leftLetter] < countT[leftLetter] {
				have--
			}
			l++
		}
	}
	if res > len(s) {return ""}
	return s[startIndex:endIndex + 1]
}
