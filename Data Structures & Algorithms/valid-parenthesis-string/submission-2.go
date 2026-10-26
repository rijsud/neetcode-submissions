func checkValidString(s string) bool {
	leftStack, starStack := []int{}, []int{}
	for i := range s {
		c := s[i]
		if c == '(' {leftStack = append(leftStack, i)}
		if c == '*' {starStack = append(starStack, i)}
		if c == ')' {
			if len(leftStack) > 0 {
				leftStack = leftStack[:len(leftStack)-1]
			} else if len(starStack) > 0 {
				starStack = starStack[:len(starStack)-1]
			} else {
				return false
			}
		}
	}

	for len(leftStack) > 0 {
		if len(starStack) == 0 {return false}
		left, star := 0, 0
		left, leftStack = leftStack[len(leftStack)-1], leftStack[:len(leftStack)-1]
		star, starStack = starStack[len(starStack)-1], starStack[:len(starStack)-1]
		if left > star {return false}
	}

	return true
}
