func lengthOfLongestSubstring(s string) int {
	stringMap := make(map[rune]int)
	startIndex := 0
	maxLength := 0
	for index, letter := range s {
		if val, ok := stringMap[letter]; ok && val >= startIndex {
			startIndex = val + 1
		}
		stringMap[letter] = index
		maxLength = max(index - startIndex + 1, maxLength)
	}
	return maxLength
}
