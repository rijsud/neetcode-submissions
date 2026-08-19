func partition(s string) [][]string {
	res := [][]string{}
	n := len(s)
	subset := []string{}

	var backtrack func(int)
	backtrack = func(i int) {
		if i >= n {
			res = append(res, append([]string{}, subset...))
			return
		}

		for j := i; j < n; j++ {
			if palindrome(s[i:j+1]) {
				subset = append(subset, s[i:j+1])
				backtrack(j+1)
				subset = subset[:len(subset)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func palindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {return false}
		l++
		r--
	}
	return true
}