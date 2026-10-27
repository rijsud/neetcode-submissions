func checkValidString(s string) bool {
	leftStack, starStack := []int{}, []int{}
	for i, c := range s {
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

	for len(leftStack) > 0 && len(starStack) > 0 {
		if leftStack[len(leftStack)-1] > starStack[len(starStack)-1] {return false}
		leftStack = leftStack[:len(leftStack)-1]
		starStack = starStack[:len(starStack)-1]
	}

	return len(leftStack) == 0
}
