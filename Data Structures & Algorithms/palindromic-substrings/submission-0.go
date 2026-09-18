func countSubstrings(s string) int {
	res := 0
    for i := 0; i < len(s); i++ {
		l, r := i, i
		for l > -1 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
		l, r = i, i+1
		for l > -1 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}
