func isValid(s string) bool {
	stack := []rune{}
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, letter := range s {
		if open, exists := closeToOpen[letter]; exists {
			if len(stack) != 0 {
				top := stack[len(stack)-1]
				if top != open {
					return false
				}
				stack = stack[:len(stack)-1]
			} else { return false } // edge case ")"
		} else {
			stack = append(stack, letter)
		}
	}

	return len(stack) == 0
}
