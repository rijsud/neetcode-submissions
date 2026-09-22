func numDecodings(s string) int {
    dp, dp1, dp2 := 0, 1, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '0' {
            dp = 0
        } else {
            dp = dp1
			if i+1 < len(s) && 
			(s[i] == '1' || s[i] == '2' && s[i+1] < '7') {
				dp += dp2
			}
        }
        dp2, dp1, dp = dp1, dp, 0
    }
    return dp1
}