func generateParenthesis(n int) []string {
	res := []string{}

	var backtrack func(int, int, string)
	backtrack = func(openN int, closeN int, current string) {
		if openN == n && closeN == n {
			res = append(res, current)
			return
		}
		if openN < n {
			backtrack(openN + 1, closeN, current + "(")
		}
		if closeN < openN {
			backtrack(openN, closeN + 1, current + ")")
		}
	}

	backtrack(0, 0, "")
	return res
}
