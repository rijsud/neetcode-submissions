func isValid(s string) bool {
    stack := list.New()
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, letter := range s {
		if open, exists := closeToOpen[letter]; exists {
			if stack.Len() != 0 {
				top := stack.Remove(stack.Front())
				if top.(rune) != open {
					return false
				}
			} else { return false }
		} else {
			stack.PushFront(letter)
		}
	}

	return stack.Len() == 0
}
