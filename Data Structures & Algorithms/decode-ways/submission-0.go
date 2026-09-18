func numDecodings(s string) int {
    dp := map[int]int{len(s): 1}
	
	var dfs func(int) int
	dfs = func(i int) int {
		if val, ok := dp[i]; ok {return val}
		if s[i] == '0' {return 0}

		res := dfs(i+1)
		if i + 1 < len(s) {
			if s[i] == '1' || (s[i] == '2' && s[i+1] < '7') {
				res += dfs(i+2)
			}
		}
		dp[i] = res
		return res
	}

	return dfs(0)
}
