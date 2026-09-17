func longestPalindrome(s string) string {
    res := ""
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l > -1 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > len(res) {res = s[l:r+1]}
			l--
			r++
		}
		
		l, r = i, i + 1
		for l > -1 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > len(res) {res = s[l:r+1]}
			l--
			r++
		} 
	}
	return res
}
