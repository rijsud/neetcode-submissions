func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]struct{})
	startIndex := 0
	maxLength := 0

	for index, _ := range s {
		for _, ok := charSet[s[index]]; ok; _, ok = charSet[s[index]] {
			delete(charSet, s[startIndex])
			startIndex++
		}
		charSet[s[index]] = struct{}{}
		maxLength = max(index - startIndex + 1, maxLength)
	}
	return maxLength
}