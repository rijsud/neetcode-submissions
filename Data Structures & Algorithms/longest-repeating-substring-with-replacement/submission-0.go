func characterReplacement(s string, k int) int {
	letterMap := make(map[byte]int)
	maxLength := 0
	l := 0
	maxFreq := 0

	for r := 0; r < len(s); r++ {
		letterMap[s[r]]++
		maxFreq = max(maxFreq, letterMap[s[r]])
		for ((r - l + 1) - maxFreq) > k {
			letterMap[s[l]]--
			l++
		}
		maxLength = max(maxLength, r - l + 1)
	}

	return maxLength
}
