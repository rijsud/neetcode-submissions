func checkValidString(s string) bool {
    leftMin, leftMax := 0, 0
	for _, c := range s {
		if c == '(' {
			leftMax++
			leftMin++
		} else if c == ')' {
			leftMax--
			leftMin--
		} else {
			leftMax++
			leftMin--
		}
		if leftMin < 0 {
			leftMin = 0
		}
		if leftMax < 0 {
			return false
		}
	}

	return leftMin == 0
}
