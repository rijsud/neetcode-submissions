func longestPalindrome(s string) string {
    res := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i] != s[j] { continue }
			if (j + 1 - i) > len(res) && palindrome(s[i:j+1]) {
				res = s[i:j+1]
			}
		}
	}
	return res
}

func palindrome(s string) bool {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		if s[i] != s[j] {return false}
	}
	return true
}