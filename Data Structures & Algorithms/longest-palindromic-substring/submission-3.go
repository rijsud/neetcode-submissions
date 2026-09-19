func longestPalindrome(s string) string {
    res := ""
	for i := 0; i < len(s); i++ {
		res = checkPali(i, i, s, res)
		res = checkPali(i, i + 1, s, res)
	}
	return res
}

func checkPali(l, r int, s, res string) string {
	for l > -1 && r < len(s) && s[l] == s[r] {
		if (r - l + 1) > len(res) {res = s[l:r+1]}
		l--
		r++
	}
	return res
}